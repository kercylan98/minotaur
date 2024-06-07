package log

import (
	"github.com/fatih/color"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"log/slog"
	"reflect"
	"time"
)

// NewDevHandlerOptions 创建一个适用于开发环境的 HandlerOptions
//   - 该可选项默认提供了一个具有色彩且日志级别为 slog.LevelDebug，并且具有对 error 类型的堆栈美化后的追踪的 HandlerOptions
func NewDevHandlerOptions() *HandlerOptions {
	return new(HandlerOptions).
		WithLevel(slog.LevelDebug).
		WithEnableColor(true).
		WithErrTrackLevel(slog.LevelError).
		WithTrackBeautify(true)
}

// NewDevGolandHandlerOptions 创建一个适用于 Goland 开发环境的 HandlerOptions
//   - 与 NewDevHandlerOptions 不同的是，该可选项的日志级别会被输出为适用于 Goland 控制台色彩的字符串
func NewDevGolandHandlerOptions() *HandlerOptions {
	return NewDevHandlerOptions().
		WithLevelStr(slog.LevelDebug, "DEBUG").
		WithLevelStr(slog.LevelInfo, "INFO").
		WithLevelStr(slog.LevelWarn, "WARN").
		WithLevelStr(slog.LevelError, "ERROR")
}

// NewTestHandlerOptions 创建一个适用于测试环境的 HandlerOptions
//   - 该可选项适用于在服务器上运行，该可选项与 NewDevHandlerOptions 相似，但是不包含任何色彩
func NewTestHandlerOptions() *HandlerOptions {
	return NewDevHandlerOptions().
		WithEnableColor(false)
}

// NewProdHandlerOptions 创建一个适用于生产环境的 HandlerOptions
//   - 该可选项适用于在服务器上运行，该可选项与 NewDevHandlerOptions 相似，但是不包含任何色彩，且不包含对 error 的追踪，默认日志级别为 slog.LevelInfo
func NewProdHandlerOptions() *HandlerOptions {
	return new(HandlerOptions).
		WithLevel(slog.LevelInfo).
		WithEnableColor(false)
}

type (
	CallerFormatter  func(file string, line int) (repFile, repLine string)
	MessageFormatter func(message string) string
)

type HandlerOption func(opts *HandlerOptions)

type HandlerOptions struct {
	options          []HandlerOption
	Level            slog.Level                 // 日志级别
	TimeLayout       string                     // 时间格式
	ColorTypes       map[ColorType]*color.Color // 颜色类型
	EnableColor      bool                       // 是否启用颜色
	AttrKeys         map[AttrKey]string         // 属性键
	Delimiter        string                     // 分隔符
	LevelStr         map[slog.Level]string      // 日志级别字符串
	Caller           bool                       // 是否显示调用者
	CallerSkip       int                        // 调用者跳过层数
	CallerFormatter  CallerFormatter            // 调用者格式化
	MessageFormatter MessageFormatter           // 消息格式化
	ErrTrackLevel    map[slog.Level]struct{}    // 错误追踪级别
	TrackBeautify    bool                       // 错误追踪美化
}

func (o *HandlerOptions) applyDefault() *HandlerOptions {
	return o.
		WithLevel(slog.LevelInfo).
		WithTimeLayout(time.DateTime).
		WithDelimiter("=").
		WithLevelStr(slog.LevelDebug, "DBG").
		WithLevelStr(slog.LevelInfo, "INF").
		WithLevelStr(slog.LevelWarn, "WAR").
		WithLevelStr(slog.LevelError, "ERR").
		WithColor(ColorTypeDebugLevel, color.FgHiCyan).
		WithColor(ColorTypeInfoLevel, color.FgHiGreen).
		WithColor(ColorTypeWarnLevel, color.FgHiYellow).
		WithColor(ColorTypeErrorLevel, color.FgHiRed).
		WithColor(ColorTypeMessage, color.FgHiBlack, color.Bold).
		WithColor(ColorTypeAttrDelimiter, color.FgHiBlack).
		WithColor(ColorTypeAttrKey, color.FgWhite).
		WithColor(ColorTypeAttrErrorKey, color.FgHiRed).
		WithColor(ColorTypeAttrErrorValue, color.FgHiRed).
		WithColor(ColorTypeErrorTrack, color.FgWhite).
		WithColor(ColorTypeErrorTrackHeader, color.FgYellow).
		WithCaller(true).
		WithCallerSkip(5).
		WithMessageFormatter(func(message string) string {
			return charproc.BigCamel(message)
		})
}

// WithTrackBeautify 设置错误追踪美化
func (o *HandlerOptions) WithTrackBeautify(beautify bool) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		opts.TrackBeautify = beautify
	})
	return o
}

// WithErrTrackLevel 设置错误追踪级别
func (o *HandlerOptions) WithErrTrackLevel(levels ...slog.Level) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		if opts.ErrTrackLevel == nil {
			opts.ErrTrackLevel = make(map[slog.Level]struct{})
		}
		for _, level := range levels {
			opts.ErrTrackLevel[level] = struct{}{}
		}
	})
	return o
}

// WithMessageFormatter 设置消息格式化
func (o *HandlerOptions) WithMessageFormatter(formatter MessageFormatter) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		opts.MessageFormatter = formatter
	})
	return o
}

// WithCaller 设置是否显示调用者
func (o *HandlerOptions) WithCaller(caller bool) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		opts.Caller = caller
	})
	return o
}

// WithCallerSkip 设置调用者跳过层数
func (o *HandlerOptions) WithCallerSkip(skip int) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		opts.CallerSkip = skip
	})
	return o
}

// WithCallerFormatter 设置调用者格式化
func (o *HandlerOptions) WithCallerFormatter(formatter CallerFormatter) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		opts.CallerFormatter = formatter
	})
	return o
}

// WithLevelStr 设置日志级别字符串
func (o *HandlerOptions) WithLevelStr(level slog.Level, str string) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		if opts.LevelStr == nil {
			opts.LevelStr = make(map[slog.Level]string)
		}
		opts.LevelStr[level] = str
	})
	return o
}

// WithDelimiter 设置分隔符
func (o *HandlerOptions) WithDelimiter(delimiter string) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		opts.Delimiter = delimiter
	})
	return o
}

// WithAttrKey 设置属性键
func (o *HandlerOptions) WithAttrKey(key AttrKey, value string) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		if opts.AttrKeys == nil {
			opts.AttrKeys = make(map[AttrKey]string)
		}
		opts.AttrKeys[key] = value
	})
	return o
}

// WithEnableColor 设置是否启用颜色
func (o *HandlerOptions) WithEnableColor(enable bool) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		opts.EnableColor = enable
	})
	return o
}

// WithColor 设置日志颜色
func (o *HandlerOptions) WithColor(colorType ColorType, attrs ...color.Attribute) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		if opts.ColorTypes == nil {
			opts.ColorTypes = make(map[ColorType]*color.Color)
		}
		c := color.New(attrs...)
		c.EnableColor()
		opts.ColorTypes[colorType] = c
	})
	return o
}

// WithTimeLayout 设置日志时间格式，如 "2006-01-02 15:04:05"
func (o *HandlerOptions) WithTimeLayout(layout string) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		opts.TimeLayout = layout
	})
	return o
}

// WithLevel 设置日志级别
func (o *HandlerOptions) WithLevel(level slog.Level) *HandlerOptions {
	o.options = append(o.options, func(opts *HandlerOptions) {
		opts.Level = level
	})
	return o
}

func (o *HandlerOptions) apply(opts ...*HandlerOptions) *HandlerOptions {
	vof := reflect.ValueOf(o).Elem()
	for _, opt := range append([]*HandlerOptions{o}, opts...) {
		// Fields
		reflectValue := reflect.ValueOf(opt).Elem()
		for i := 0; i < reflectValue.NumField(); i++ {
			field := reflectValue.Field(i)
			// 仅处理公开字段
			if !field.IsZero() && field.CanSet() {
				field.Set(vof.Field(i))
			}
		}

		// Options
		for _, option := range opt.options {
			option(o)
		}
	}

	return o
}
