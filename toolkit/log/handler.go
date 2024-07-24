package log

import (
	"context"
	"encoding"
	"fmt"
	"github.com/fatih/color"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

func NewHandler(w io.Writer, opts ...*HandlerOptions) *Handler {
	return &Handler{
		opts: new(HandlerOptions).applyDefault().apply(opts...),
		w:    w,
	}
}

type Handler struct {
	opts  *HandlerOptions
	w     io.Writer
	attrs []slog.Attr
	group string
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.leveler.Level()
}

func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	if !h.Enabled(ctx, h.opts.leveler.Level()) {
		return nil
	}

	var builder = charproc.NewBuilder()
	defer builder.Reset()

	h.formatTime(ctx, record, builder)
	h.formatLevel(ctx, record, builder)
	h.formatCaller(ctx, record, builder)
	h.formatMessage(ctx, record, builder)

	// fixed attrs
	num := record.NumAttrs()
	fixedNum := len(h.attrs)
	for i, attr := range h.attrs {
		h.formatAttr(ctx, h.group, record.Level, attr, builder, num+fixedNum == i+1)
	}

	idx := 0
	record.Attrs(func(attr slog.Attr) bool {
		idx++
		h.formatAttr(ctx, h.group, record.Level, attr, builder, num == idx)
		return true
	})

	recordBytes, err := builder.Write('\n').Bytes()
	if err != nil {
		return err
	}

	_, err = h.w.Write(recordBytes)
	return err
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	n := h.clone()
	n.attrs = append(n.attrs, attrs...)
	return n
}

func (h *Handler) WithGroup(name string) slog.Handler {
	n := h.clone()
	n.group = name
	return n
}

func (h *Handler) clone() *Handler {
	return &Handler{
		opts:  h.opts,
		w:     h.w,
		attrs: h.attrs,
		group: h.group,
	}
}

func (h *Handler) formatTime(ctx context.Context, record slog.Record, builder *charproc.Builder) {
	h.loadAttrKey(builder, AttrKeyTime)
	h.loadColor(builder, ColorTypeTime).
		WriteString(record.Time.Format(h.opts.TimeLayout)).
		DisableColor().
		Write(' ')
}

func (h *Handler) formatLevel(ctx context.Context, record slog.Record, builder *charproc.Builder) {
	var colorType ColorType
	if h.opts.EnableColor {
		switch record.Level {
		case slog.LevelDebug:
			colorType = ColorTypeDebugLevel
		case slog.LevelInfo:
			colorType = ColorTypeInfoLevel
		case slog.LevelWarn:
			colorType = ColorTypeWarnLevel
		case slog.LevelError:
			colorType = ColorTypeErrorLevel
		}
	}
	h.loadAttrKey(builder, AttrKeyLevel)
	h.loadColor(builder, colorType).
		WriteString(h.opts.LevelStr[record.Level]).
		DisableColor().
		Write(' ')
}

func (h *Handler) formatCaller(ctx context.Context, record slog.Record, builder *charproc.Builder) {
	if !h.opts.Caller {
		return
	}
	pcs := make([]uintptr, 1)
	runtime.CallersFrames(pcs[:runtime.Callers(h.opts.CallerSkip, pcs)])
	fs := runtime.CallersFrames(pcs)
	f, _ := fs.Next()
	if f.File == charproc.None {
		return
	}

	var file, line string
	if h.opts.CallerFormatter != nil {
		file, line = h.opts.CallerFormatter(f.File, f.Line)
	} else {
		file = filepath.Base(f.File)
		line = convert.IntToString(f.Line)
	}

	h.loadAttrKey(builder, AttrKeyCaller)
	h.loadColor(builder, ColorTypeCaller).
		WriteString(file).
		SetColor(h.opts.ColorTypes[ColorTypeAttrDelimiter]).
		WriteString(":").
		SetColor(h.opts.ColorTypes[ColorTypeAttrValue]).
		WriteString(line).
		DisableColor().
		Write(' ')
}

func (h *Handler) formatMessage(ctx context.Context, record slog.Record, builder *charproc.Builder) {
	if record.Message == "" {
		return
	}
	var msg = record.Message
	if h.opts.MessageFormatter != nil {
		msg = h.opts.MessageFormatter(msg)
	}

	h.loadAttrKey(builder, AttrKeyMessage)
	h.loadColor(builder, ColorTypeMessage).
		WriteString(msg).
		DisableColor().
		Write(' ')
}

func (h *Handler) formatAttr(ctx context.Context, group string, level slog.Level, attr slog.Attr, builder *charproc.Builder, last bool) {
	var key = attr.Key
	if group != "" {
		key = group + "." + key
	}

	switch attr.Value.Kind() {
	case slog.KindGroup:
		groupAttr := attr.Value.Group()
		for _, a := range groupAttr {
			h.formatAttr(ctx, key, level, a, builder, last)
		}
		return
	default:
		h.loadColor(builder, ColorTypeAttrKey)
		switch v := attr.Value.Any().(type) {
		case stackError, stackErrorTracks:
			h.loadColor(builder, ColorTypeAttrErrorKey)
		case error:
			if _, ok := h.opts.ErrTrackLevel[level]; ok && !h.opts.TrackBeautify {
				pc := make([]uintptr, 10)
				n := runtime.Callers(h.opts.CallerSkip+3, pc)
				frames := runtime.CallersFrames(pc[:n])
				var stacks = make(stackErrorTracks, 0, 10)
				for {
					frame, more := frames.Next()
					stacks = append(stacks, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
					if !more {
						break
					}
				}
				attr = slog.Group(attr.Key, slog.Any("info", stackError{v}), slog.Any("stack", stacks))
				h.formatAttr(ctx, group, level, attr, builder, false)
				return
			}
			h.loadColor(builder, ColorTypeAttrErrorKey)
		}
	}

	builder.
		WriteString(key).
		SetColor(h.opts.ColorTypes[ColorTypeAttrDelimiter]).
		WriteString(h.opts.Delimiter)
	h.formatAttrValue(ctx, level, key, attr, builder, last)
}

func (h *Handler) formatAttrValue(ctx context.Context, level slog.Level, fullKey string, attr slog.Attr, builder *charproc.Builder, last bool) {
	h.loadColor(builder, ColorTypeAttrValue)
	defer builder.DisableColor()

	switch attr.Value.Kind() {
	case slog.KindString:
		builder.WriteString(strconv.Quote(attr.Value.String()))
	case slog.KindInt64:
		builder.WriteInt64(attr.Value.Int64())
	case slog.KindUint64:
		builder.WriteUint64(attr.Value.Uint64())
	case slog.KindFloat64:
		builder.WriteFloat64(attr.Value.Float64())
	case slog.KindBool:
		builder.WriteBool(attr.Value.Bool())
	case slog.KindDuration:
		builder.WriteString(strconv.Quote(attr.Value.Duration().String()))
	case slog.KindTime:
		builder.WriteString(strconv.Quote(attr.Value.Time().String()))
	default:
		switch v := attr.Value.Any().(type) {
		case stackError:
			h.loadColor(builder, ColorTypeAttrErrorKey)
			builder.WriteString(strconv.Quote(v.err.Error()))
		case stackErrorTracks:
			h.loadColor(builder, ColorTypeAttrErrorKey)
			builder.WriteString(strconv.Quote(fmt.Sprintf("%+v", attr.Value.Any())))
		case error:
			h.loadColor(builder, ColorTypeAttrErrorValue)
			builder.WriteString(strconv.Quote(v.Error()))

			if _, ok := h.opts.ErrTrackLevel[level]; ok && h.opts.TrackBeautify {
				pc := make([]uintptr, 10)
				n := runtime.Callers(h.opts.CallerSkip+3, pc)
				frames := runtime.CallersFrames(pc[:n])
				if h.opts.TrackBeautify {
					h.loadColor(builder, ColorTypeErrorTrackHeader).
						WriteSprintfToEnd("\tError Track: [%s] >> %s", fullKey, v.Error())
					h.loadColor(builder, ColorTypeErrorTrack)
					for {
						builder.WriteToEnd('\n')
						frame, more := frames.Next()
						builder.WriteToEnd('\t')
						builder.WriteStringToEnd(frame.File)
						builder.WriteToEnd(':')
						builder.WriteIntToEnd(frame.Line)
						builder.WriteToEnd(' ')
						builder.WriteStringToEnd(frame.Function)
						if !more {
							break
						}
					}
					builder.WriteToEnd('\n')
				}
			}
		case nil:
			builder.WriteString("<nil>")
		case encoding.TextMarshaler:
			data, err := v.MarshalText()
			if err != nil {
				break
			}
			builder.WriteString(strconv.Quote(string(data)))
		case []byte:
			builder.WriteString(strconv.Quote(*(*string)(unsafe.Pointer(&v))))
		case stack:
			if len(v) == 0 {
				builder.WriteString("<none>")
			} else {
				lines := strings.Split(string(v), "\n")
				builder.WriteString(fmt.Sprintf("lines(%d)", len(lines)))
				if h.opts.TrackBeautify {
					for _, line := range lines {
						builder.WriteToEnd('\n')
						builder.WriteStringToEnd(line)
					}
					builder.WriteToEnd('\n')
				}
			}

		default:
			//builder.WriteString(strconv.Quote(fmt.Sprintf("%+v", attr.Value.Any())))
			builder.WriteString(string(toolkit.MarshalJSON(attr.Value.Any())))
		}
	}

	if !last {
		builder.Write(' ')
	}
}

func (h *Handler) loadColor(builder *charproc.Builder, t ColorType) *charproc.Builder {
	var c *color.Color
	if h.opts.EnableColor {
		c = h.opts.ColorTypes[t]
	}
	return builder.SetColor(c)
}

func (h *Handler) loadAttrKey(builder *charproc.Builder, key AttrKey) *charproc.Builder {
	v, exist := h.opts.AttrKeys[key]
	if !exist {
		return builder
	}
	return builder.
		SetColor(h.opts.ColorTypes[ColorTypeAttrKey]).
		WriteString(v).
		SetColor(h.opts.ColorTypes[ColorTypeAttrDelimiter]).
		WriteString(h.opts.Delimiter)
}
