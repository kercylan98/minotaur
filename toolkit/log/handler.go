package log

import (
	"context"
	"encoding"
	"fmt"
	"github.com/kercylan98/minotaur/internal/utils/str"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"io"
	"log/slog"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

func NewHandler(w io.Writer, opts ...*Options) *MinotaurHandler {
	h := &MinotaurHandler{
		opts: DefaultOptions(),
		w:    w,
	}
	for _, opt := range opts {
		h.opts.Apply(opt)
	}
	return h
}

type MinotaurHandler struct {
	opts        *Options
	groupPrefix string
	groups      []string
	mu          sync.Mutex
	w           io.Writer
}

func (h *MinotaurHandler) GetOptions() *Options {
	return h.opts
}

func (h *MinotaurHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.GetLevel()
}

func (h *MinotaurHandler) Handle(ctx context.Context, record slog.Record) (err error) {
	h.opts.getMany(func(opt *Options) {
		if !h.Enabled(ctx, opt.level) {
			return
		}

		var buffer = new(strings.Builder)
		defer buffer.Reset()

		processTime(buffer, record, opt)
		processLevel(buffer, record, opt)
		processCaller(buffer, record, opt)
		processMessage(buffer, record, opt)
		processAttrs(buffer, h, record, opt, record.Level, h.groupPrefix, h.groups)

		if buffer.Len() == 0 {
			return
		}
		buffer.WriteByte('\n')

		h.mu.Lock()
		defer h.mu.Unlock()
		_, err = h.w.Write([]byte(buffer.String()))
		return
	})

	return
}

func (h *MinotaurHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	handler := h.clone()

	buffer := new(strings.Builder)
	defer buffer.Reset()

	h.opts.getMany(func(opt *Options) {
		for _, attr := range attrs {
			processAttrsAttr(buffer, h, attr, opt, levelNone, h.groupPrefix, h.groups)
		}
	})

	//MinotaurHandler.opts.w = h.opts.FieldPrefix + string(*buf.bytes)
	return handler
}

func (h *MinotaurHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	handler := h.clone()
	handler.groupPrefix += name + "."
	handler.groups = append(handler.groups, name)
	return handler
}

func (h *MinotaurHandler) clone() *MinotaurHandler {
	return &MinotaurHandler{
		groupPrefix: h.groupPrefix,
		opts:        DefaultOptions().Apply(h.opts),
		groups:      h.groups,
		w:           h.w,
	}
}

func processTime(buffer *strings.Builder, record slog.Record, opt *Options) {
	if record.Time.IsZero() {
		return
	}

	processAttrType(buffer, opt, AttrTypeTime, record.Time.Format(opt.timeLayout))
}

func processLevel(buffer *strings.Builder, record slog.Record, opt *Options) {
	var levelColor = opt.levelColor[record.Level]
	var levelText = opt.levelText[record.Level]
	if levelText == charproc.None {
		return
	}

	if opt.disabledColor || levelColor == nil {
		buffer.WriteString(levelText)
	} else {
		buffer.WriteString(levelColor.Sprint(levelText))
	}
	buffer.WriteByte(' ')
}

func processCaller(buffer *strings.Builder, record slog.Record, opt *Options) {
	if opt.disabledCaller {
		return
	}

	pcs := make([]uintptr, 1)
	runtime.CallersFrames(pcs[:runtime.Callers(opt.callerSkip, pcs)])
	fs := runtime.CallersFrames(pcs)
	f, _ := fs.Next()
	if f.File == str.None {
		return
	}

	file, line := opt.callerFormatter(f.File, f.Line)
	processAttrType(buffer, opt, AttrTypeCaller, file+":"+line)
}

func processMessage(buffer *strings.Builder, record slog.Record, opt *Options) {
	processAttrType(buffer, opt, AttrTypeMessage, record.Message)
}

func processAttrs(buffer *strings.Builder, handler *MinotaurHandler, record slog.Record, opt *Options, level Level, groupsPrefix string, groups []string) {
	record.Attrs(func(attr slog.Attr) bool {
		processAttrsAttr(buffer, handler, attr, opt, level, groupsPrefix, groups)
		return true
	})
}

func processAttrsAttr(buffer *strings.Builder, handler *MinotaurHandler, attr slog.Attr, opt *Options, level Level, groupsPrefix string, groups []string) {
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
			processAttrsAttr(buffer, handler, groupAttr, opt, level, groupsPrefix, groups)
		}
	default:
		switch v := attr.Value.Any().(type) {
		case error:
			if opt.stackTrace[level] {
				stackTraceAttr := slog.Attr{Key: attr.Key, Value: formatTraceError(v, opt.stackTraceBeauty[level])}
				processAttrsAttr(buffer, handler, stackTraceAttr, opt, level, handler.groupPrefix, handler.groups)
				return
			}
			color := opt.keyColor[AttrTypeError]
			if opt.disabledColor || color == nil {
				processAttrsKey(buffer, opt, attr.Key, groupsPrefix)
			} else {
				processAttrsKey(buffer, opt, attr.Key, groupsPrefix, color)
			}
			processAttrsValue(buffer, opt, attr.Value, true)
			buffer.WriteByte(' ')
		case *beautyTrace:
			color := opt.valueColor[AttrTypeTrace]
			for _, s := range v.trace {
				buffer.WriteString("\n\t")
				if opt.disabledColor || color == nil {
					buffer.WriteString(s)
				} else {
					buffer.WriteString(color.Sprint(s))
				}
			}
		default:
			processAttrsKey(buffer, opt, attr.Key, groupsPrefix)
			processAttrsValue(buffer, opt, attr.Value, true)
			buffer.WriteByte(' ')
		}
	}
}

func processAttrsString(s string, quote bool) string {
	quoting := len(s) == 0
	for _, r := range s {
		if unicode.IsSpace(r) || r == '"' || r == '=' || !unicode.IsPrint(r) {
			quoting = true
			break
		}
	}
	if quote && quoting {
		return strconv.Quote(s)
	} else {
		return s
	}
}

func processAttrsKey(buffer *strings.Builder, opt *Options, key, groups string, replaceColor ...*Color) {
	if key == str.None {
		return
	}
	color := opt.keyColor[AttrTypeField]
	if len(replaceColor) > 0 {
		color = replaceColor[0]
	}
	if opt.disabledColor || color == nil {
		buffer.WriteString(processAttrsString(groups+key, true))
	} else {
		buffer.WriteString(color.Sprint(processAttrsString(groups+key, true)))
	}

	delimiterText := opt.delimiterText[AttrTypeField]
	if delimiterText != str.None {
		delimiterColor := opt.delimiterColor[AttrTypeField]
		if len(replaceColor) > 1 {
			delimiterColor = replaceColor[1]
		}
		if opt.disabledColor || delimiterColor == nil {
			buffer.WriteString(delimiterText)
		} else {
			buffer.WriteString(delimiterColor.Sprint(delimiterText))
		}
	}
}

func processAttrsValue(buffer *strings.Builder, opt *Options, v slog.Value, quote bool) {
	var text string
	var color = opt.valueColor[AttrTypeField]
	switch v.Kind() {
	case slog.KindString:
		text = processAttrsString(v.String(), quote)
	case slog.KindInt64:
		text = strconv.FormatInt(v.Int64(), 10)
	case slog.KindUint64:
		text = strconv.FormatUint(v.Uint64(), 10)
	case slog.KindFloat64:
		text = strconv.FormatFloat(v.Float64(), 'g', -1, 64)
	case slog.KindBool:
		text = strconv.FormatBool(v.Bool())
	case slog.KindDuration:
		text = processAttrsString(v.Duration().String(), quote)
	case slog.KindTime:
		text = processAttrsString(v.Time().String(), quote)
	case slog.KindAny:
		switch cv := v.Any().(type) {
		case slog.Level:
			processLevel(buffer, slog.Record{Level: cv}, opt)
		case encoding.TextMarshaler:
			data, err := cv.MarshalText()
			if err != nil {
				break
			}
			text = processAttrsString(string(data), quote)
		case *slog.Source:
			file, line := opt.callerFormatter(cv.File, cv.Line)
			callerColor := opt.valueColor[AttrTypeCaller]
			if opt.disabledColor || callerColor == nil {
				buffer.WriteString(file + ":" + line)
			} else {
				buffer.WriteString(callerColor.Sprintf("%s:%s", file, line))
			}
			buffer.WriteByte(' ')
		default:
			text = processAttrsString(fmt.Sprint(v.Any()), quote)
		}
	default:
	}
	if text == str.None {
		return
	}

	if opt.disabledColor || color == nil {
		buffer.WriteString(text)
	} else {
		buffer.WriteString(color.Sprint(text))
	}
}

func processAttrType(buffer *strings.Builder, opt *Options, attrType AttrType, value string) {
	prefixText := opt.keyText[attrType]
	prefixColor := opt.keyColor[attrType]
	delimiterText := opt.delimiterText[attrType]
	delimiterColor := opt.delimiterColor[attrType]
	valueColor := opt.valueColor[attrType]

	if prefixText != str.None {
		// 前缀
		if opt.disabledColor || prefixColor == nil {
			buffer.WriteString(prefixText)
		} else {
			buffer.WriteString(prefixColor.Sprint(prefixText))
		}

		// 分隔符
		if delimiterText != str.None {
			if opt.disabledColor || delimiterColor == nil {
				buffer.WriteString(delimiterText)
			} else {
				buffer.WriteString(delimiterColor.Sprint(delimiterText))
			}
		}
	}

	// 时间信息
	if opt.disabledColor || valueColor == nil {
		buffer.WriteString(value)
	} else {
		buffer.WriteString(valueColor.Sprint(value))
	}
	buffer.WriteByte(' ')
}
