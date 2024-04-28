package log

import (
	"context"
	"encoding"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strconv"
	"sync"
	"time"
	"unicode"
)

const (
	ErrKey = "err"
)

// NewHandler 创建一个更偏向于人类可读的处理程序，该处理程序也是默认的处理程序
func NewHandler(w io.Writer, opts *Options) slog.Handler {
	if opts == nil {
		opts = NewOptions()
	}
	return &handler{
		opts: opts.apply(),
		w:    w,
	}
}

type handler struct {
	opts        *Options
	groupPrefix string
	groups      []string

	mu sync.Mutex
	w  io.Writer
}

func (h *handler) clone() *handler {
	return &handler{
		groupPrefix: h.groupPrefix,
		opts:        NewOptions().With(h.opts).apply(),
		groups:      h.groups,
		w:           h.w,
	}
}

func (h *handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Load().(slog.Leveler).Level()
}

func (h *handler) Handle(_ context.Context, r slog.Record) error {
	lv := h.opts.Level.Load().(slog.Leveler).Level()
	if r.Level < lv {
		return nil
	}

	buf := newBuffer(h)
	defer buf.Free()

	if !r.Time.IsZero() {
		h.appendTime(buf, r.Time)
		buf.WriteBytes(' ')
	}

	h.appendLevel(buf, r.Level)
	buf.WriteBytes(' ')

	if h.opts.Caller {
		pcs := make([]uintptr, 1)
		runtime.CallersFrames(pcs[:runtime.Callers(h.opts.CallerSkip, pcs)])
		fs := runtime.CallersFrames(pcs)
		f, _ := fs.Next()
		if f.File != "" {
			src := &slog.Source{
				Function: f.Function,
				File:     f.File,
				Line:     f.Line,
			}

			h.appendSource(buf, src)
			buf.WriteBytes(' ')
		}
	}

	if r.Message != "" {
		buf.WriteColorString(h.opts.MessageColor, r.Message, ColorDefault).WriteBytes(' ')
	}

	if len(h.opts.FieldPrefix) > 0 {
		buf.WriteString(h.opts.FieldPrefix)
	}

	r.Attrs(func(attr slog.Attr) bool {
		h.appendAttr(buf, attr, h.groupPrefix, h.groups, h.opts.ErrTrace)
		return true
	})

	if len(*buf.bytes) == 0 {
		return nil
	}
	buf.WriteBytes(' ', '\n')

	h.mu.Lock()
	defer h.mu.Unlock()

	_, err := h.w.Write(*buf.bytes)
	if lv == DPanicLevel && h.opts.DevMode {
		panic(r.Message)
	}
	return err
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	h2 := h.clone()

	buf := newBuffer(h)
	defer buf.Free()

	for _, attr := range attrs {
		h.appendAttr(buf, attr, h.groupPrefix, h.groups)
	}
	h2.opts.FieldPrefix = h.opts.FieldPrefix + string(*buf.bytes)
	return h2
}

func (h *handler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	h2 := h.clone()
	h2.groupPrefix += name + "."
	h2.groups = append(h2.groups, name)
	return h2
}

func (h *handler) appendTime(buf *buffer, t time.Time) {
	if h.opts.TimePrefix != "" {
		buf.WriteColorString(
			h.opts.TimePrefixColor, h.opts.TimePrefix, ColorDefault, // 时间前缀
			h.opts.TimePrefixDelimiterColor, h.opts.TimePrefixDelimiter, ColorDefault, // 时间前缀分隔符
		)
	}
	buf.WriteColorString(h.opts.TimeColor)
	*buf.bytes = t.AppendFormat(*buf.bytes, h.opts.TimeLayout)
	buf.WriteColorString(ColorDefault)
}

func (h *handler) appendLevel(buf *buffer, level slog.Level) {
	switch {
	case level < InfoLevel:
		buf.WriteColorString(h.opts.LevelColor[DebugLevel], h.opts.LevelText[DebugLevel])
		appendLevelDelta(buf, level-DebugLevel)
		buf.WriteColorString(ColorDefault)
	case level < WarnLevel:
		buf.WriteColorString(h.opts.LevelColor[InfoLevel], h.opts.LevelText[InfoLevel])
		appendLevelDelta(buf, level-InfoLevel)
		buf.WriteColorString(ColorDefault)
	case level < ErrorLevel:
		buf.WriteColorString(h.opts.LevelColor[WarnLevel], h.opts.LevelText[WarnLevel])
		appendLevelDelta(buf, level-WarnLevel)
		buf.WriteColorString(ColorDefault)
	case level < PanicLevel:
		buf.WriteColorString(h.opts.LevelColor[ErrorLevel], h.opts.LevelText[ErrorLevel])
		appendLevelDelta(buf, level-ErrorLevel)
		buf.WriteColorString(ColorDefault)
	case level < FatalLevel:
		var tag = h.opts.LevelText[PanicLevel]
		if level == DPanicLevel {
			tag = h.opts.LevelText[DPanicLevel]
		}
		buf.WriteColorString(h.opts.LevelColor[PanicLevel], tag)
		appendLevelDelta(buf, level-PanicLevel)
		buf.WriteColorString(ColorDefault)
	default:
		buf.WriteColorString(h.opts.LevelColor[FatalLevel], h.opts.LevelText[FatalLevel])
		appendLevelDelta(buf, level-FatalLevel)
		buf.WriteColorString(ColorDefault)
	}
}

func appendLevelDelta(buf *buffer, delta slog.Level) {
	if delta == 0 {
		return
	} else if delta > 0 {
		buf.WriteBytes('+')
	}
	*buf.bytes = strconv.AppendInt(*buf.bytes, int64(delta), 10)
}

func (h *handler) appendSource(buf *buffer, src *slog.Source) {
	repFile, repLine := h.opts.CallerFormat(src.File, src.Line)
	buf.WriteColorString(h.opts.CallerColor, repFile).
		WriteBytes(':').
		WriteString(repLine).
		WriteColorString(ColorDefault)
}

func (h *handler) appendAttr(buf *buffer, attr slog.Attr, groupsPrefix string, groups []string, errTrace ...bool) {
	attr.Value = attr.Value.Resolve()

	if attr.Equal(slog.Attr{}) {
		return
	}

	switch attr.Value.Kind() {
	case slog.KindGroup:
		if attr.Key != "" {
			groupsPrefix += attr.Key + "."
			groups = append(groups, attr.Key)
		}
		for _, groupAttr := range attr.Value.Group() {
			h.appendAttr(buf, groupAttr, groupsPrefix, groups)
		}
	default:
		switch v := attr.Value.Any().(type) {
		case error:
			if len(errTrace) > 0 && errTrace[0] {
				h.appendAttr(buf, slog.Attr{Key: attr.Key, Value: formatTraceError(v, h.opts.ErrTraceBeauty)}, h.groupPrefix, h.groups)
				return
			}
			h.appendError(buf, v, groupsPrefix)
			buf.WriteBytes(' ')
		case *beautyTrace:
			for _, s := range v.trace {
				buf.WriteBytes('\n', '\t').WriteColorString(h.opts.ErrTraceColor, s, ColorDefault)
			}
		default:
			h.appendKey(buf, attr.Key, groupsPrefix)
			h.appendValue(buf, attr.Value, true)
			buf.WriteBytes(' ')
		}
	}
}

func (h *handler) appendKey(buf *buffer, key, groups string) {
	if key == "" {
		return
	}
	buf.WriteColorString(h.opts.KeyColor)
	appendString(buf, groups+key, true)
	buf.WriteColorString(ColorDefault, h.opts.KVDelimiter)
}

func (h *handler) appendValue(buf *buffer, v slog.Value, quote bool) {
	switch v.Kind() {
	case slog.KindString:
		appendString(buf, v.String(), quote)
	case slog.KindInt64:
		*buf.bytes = strconv.AppendInt(*buf.bytes, v.Int64(), 10)
	case slog.KindUint64:
		*buf.bytes = strconv.AppendUint(*buf.bytes, v.Uint64(), 10)
	case slog.KindFloat64:
		*buf.bytes = strconv.AppendFloat(*buf.bytes, v.Float64(), 'g', -1, 64)
	case slog.KindBool:
		*buf.bytes = strconv.AppendBool(*buf.bytes, v.Bool())
	case slog.KindDuration:
		appendString(buf, v.Duration().String(), quote)
	case slog.KindTime:
		appendString(buf, v.Time().String(), quote)
	case slog.KindAny:
		switch cv := v.Any().(type) {
		case slog.Level:
			h.appendLevel(buf, cv)
		case encoding.TextMarshaler:
			data, err := cv.MarshalText()
			if err != nil {
				break
			}
			appendString(buf, string(data), quote)
		case *slog.Source:
			h.appendSource(buf, cv)
		default:
			appendString(buf, fmt.Sprint(v.Any()), quote)
		}
	}
}

func (h *handler) appendError(buf *buffer, err error, groupsPrefix string) {
	buf.WriteColorString(h.opts.KeyColor)
	appendString(buf, groupsPrefix+ErrKey, true)
	buf.WriteColorString(ColorDefault, h.opts.KVDelimiter, h.opts.ErrorColor)
	appendString(buf, err.Error(), true)
	buf.WriteColorString(ColorDefault)
}

func appendString(buf *buffer, s string, quote bool) {
	quoting := len(s) == 0
	for _, r := range s {
		if unicode.IsSpace(r) || r == '"' || r == '=' || !unicode.IsPrint(r) {
			quoting = true
			break
		}
	}
	if quote && quoting {
		*buf.bytes = strconv.AppendQuote(*buf.bytes, s)
	} else {
		buf.WriteString(s)
	}
}
