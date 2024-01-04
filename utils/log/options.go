package log

import (
	"log/slog"
	"time"
)

const (
	DefaultTimeLayout          = time.DateTime
	DefaultTimePrefix          = ""
	DefaultTimePrefixDelimiter = "="
	DefaultCaller              = true
	DefaultCallerSkip          = 3
	DefaultFieldPrefix         = ""
	DefaultLevel               = DebugLevel
	DefaultErrTrace            = true
	DefaultErrTraceBeauty      = true
	DefaultKVDelimiter         = "="
	DefaultDisableColor        = false
	DefaultLevelInfoText       = "INF"
	DefaultLevelWarnText       = "WRN"
	DefaultLevelErrorText      = "ERR"
	DefaultLevelDebugText      = "DBG"
	DefaultLevelPanicText      = "PNC"
	DefaultLevelDPanicText     = "DPC"
	DefaultLevelFatalText      = "FTL"
	DefaultDevMode             = true

	DefaultTimePrefixColor          = ColorDefault
	DefaultTimePrefixDelimiterColor = ColorDefault
	DefaultTimeColor                = ColorDefaultBold
	DefaultLevelDebugColor          = ColorBlue
	DefaultLevelInfoColor           = ColorGreen
	DefaultLevelWarnColor           = ColorBrightYellow
	DefaultLevelErrorColor          = ColorRed
	DefaultLevelPanicColor          = ColorBrightRed
	DefaultLevelDPanicColor         = ColorBrightRed
	DefaultLevelFatalColor          = ColorBrightRed
	DefaultCallerColor              = ColorBrightBlueUnderline
	DefaultErrTraceColor            = ColorBrightBlack
	DefaultKeyColor                 = ColorWhite
	DefaultValueColor               = ColorDefault
	DefaultErrorColor               = ColorRedBold
	DefaultMessageColor             = ColorWhiteBold
)

// NewOptions 创建一个新的日志选项
func NewOptions() *Options {
	return (&Options{}).WithDev()
}

type (
	// Option 日志选项
	Option func(opts *Options)
	// Options 日志选项
	Options struct {
		opts                     []Option
		Handlers                 []slog.Handler                                        // 处理程序
		TimeLayout               string                                                // 时间格式化字符串
		TimePrefix               string                                                // 时间前缀
		TimePrefixDelimiter      string                                                // 时间前缀分隔符
		Caller                   bool                                                  // 是否显示调用者信息
		CallerSkip               int                                                   // 跳过的调用层数
		CallerFormat             func(file string, line int) (repFile, refLine string) // 调用者信息格式化函数
		FieldPrefix              string                                                // 字段前缀
		Level                    slog.Leveler                                          // 日志级别
		ErrTrace                 bool                                                  // 是否显示错误堆栈
		ErrTraceBeauty           bool                                                  // 是否美化错误堆栈
		KVDelimiter              string                                                // 键值对分隔符
		DisableColor             bool                                                  // 是否禁用颜色
		TimePrefixColor          string                                                // 时间前缀颜色
		TimePrefixDelimiterColor string                                                // 时间前缀分隔符颜色
		TimeColor                string                                                // 时间颜色
		LevelText                map[slog.Level]string                                 // 日志级别文本
		LevelColor               map[slog.Level]string                                 // 日志级别颜色
		CallerColor              string                                                // 调用者信息颜色
		ErrTraceColor            string                                                // 错误堆栈颜色
		KeyColor                 string                                                // 键颜色
		ValueColor               string                                                // 值颜色
		ErrorColor               string                                                // 错误颜色
		MessageColor             string                                                // 消息颜色
		DevMode                  bool                                                  // 是否为开发模式
	}
)

// WithDev 设置可选项为开发模式
//   - 开发模式适用于本地开发环境，会以更友好的方式输出日志
func (o *Options) WithDev() *Options {
	o.opts = append(o.opts, func(opts *Options) {
		opts.DevMode = DefaultDevMode
		opts.Handlers = nil
		opts.TimeLayout = DefaultTimeLayout
		opts.TimePrefix = DefaultTimePrefix
		opts.TimePrefixDelimiter = DefaultTimePrefixDelimiter
		opts.Caller = DefaultCaller
		opts.CallerSkip = DefaultCallerSkip
		opts.CallerFormat = CallerBasicFormat
		opts.FieldPrefix = DefaultFieldPrefix
		opts.Level = DefaultLevel
		opts.ErrTrace = DefaultErrTrace
		opts.ErrTraceBeauty = DefaultErrTraceBeauty
		opts.KVDelimiter = DefaultKVDelimiter
		opts.DisableColor = DefaultDisableColor
		opts.TimePrefixColor = DefaultTimePrefixColor
		opts.TimePrefixDelimiterColor = DefaultTimePrefixDelimiterColor
		opts.TimeColor = DefaultTimeColor
		opts.LevelText = map[slog.Level]string{
			DebugLevel:  DefaultLevelDebugText,
			InfoLevel:   DefaultLevelInfoText,
			WarnLevel:   DefaultLevelWarnText,
			ErrorLevel:  DefaultLevelErrorText,
			PanicLevel:  DefaultLevelPanicText,
			DPanicLevel: DefaultLevelDPanicText,
			FatalLevel:  DefaultLevelFatalText,
		}
		opts.LevelColor = map[slog.Level]string{
			DebugLevel:  DefaultLevelDebugColor,
			InfoLevel:   DefaultLevelInfoColor,
			WarnLevel:   DefaultLevelWarnColor,
			ErrorLevel:  DefaultLevelErrorColor,
			PanicLevel:  DefaultLevelPanicColor,
			DPanicLevel: DefaultLevelDPanicColor,
			FatalLevel:  DefaultLevelFatalColor,
		}
		opts.CallerColor = DefaultCallerColor
		opts.ErrTraceColor = DefaultErrTraceColor
		opts.KeyColor = DefaultKeyColor
		opts.ValueColor = DefaultValueColor
		opts.ErrorColor = DefaultErrorColor
		opts.MessageColor = DefaultMessageColor
	})
	return o
}

// WithProd 设置可选项为生产模式
//   - 生产模式适用于生产环境，会以更简洁的方式输出日志
func (o *Options) WithProd() *Options {
	o.WithDev()
	o.opts = append(o.opts, func(opts *Options) {
		opts.DisableColor = true
	})
	return o
}

// WithTest 设置可选项为测试模式
//   - 测试模式适用于测试环境，测试环境与开发环境相似，但是会屏蔽掉一些不必要的信息
func (o *Options) WithTest() *Options {
	o.WithDev()
	// 暂与开发模式相同
	return o
}

// WithDevMode 设置是否为开发模式
//   - 默认值为 DefaultDevMode
//   - 开发模式下将影响部分功能，例如 DPanic
func (o *Options) WithDevMode(enable bool) *Options {
	o.append(func(opts *Options) {
		opts.DevMode = enable
	})
	return o
}

// WithHandler 设置处理程序
func (o *Options) WithHandler(handlers ...slog.Handler) *Options {
	o.append(func(opts *Options) {
		opts.Handlers = handlers
	})
	return o
}

// WithMessageColor 设置消息颜色
//   - 默认消息颜色为 DefaultMessageColor
func (o *Options) WithMessageColor(color string) *Options {
	o.append(func(opts *Options) {
		opts.MessageColor = color
	})
	return o
}

// WithErrorColor 设置错误颜色
//   - 默认错误颜色为 DefaultErrorColor
func (o *Options) WithErrorColor(color string) *Options {
	o.append(func(opts *Options) {
		opts.ErrorColor = color
	})
	return o
}

// WithValueColor 设置值颜色
//   - 默认值颜色为 DefaultValueColor
func (o *Options) WithValueColor(color string) *Options {
	o.append(func(opts *Options) {
		opts.ValueColor = color
	})
	return o
}

// WithKeyColor 设置键颜色
//   - 默认键颜色为 DefaultKeyColor
func (o *Options) WithKeyColor(color string) *Options {
	o.append(func(opts *Options) {
		opts.KeyColor = color
	})
	return o
}

// WithErrTraceColor 设置错误堆栈颜色
//   - 默认错误堆栈颜色为 DefaultErrTraceColor
func (o *Options) WithErrTraceColor(color string) *Options {
	o.append(func(opts *Options) {
		opts.ErrTraceColor = color
	})
	return o
}

// WithCallerColor 设置调用者信息颜色
//   - 默认调用者信息颜色为 DefaultCallerColor
func (o *Options) WithCallerColor(color string) *Options {
	o.append(func(opts *Options) {
		opts.CallerColor = color
	})
	return o
}

// WithLevelColor 设置日志级别颜色
//   - 默认日志级别颜色为 DefaultLevelInfoColor, DefaultLevelWarnColor, DefaultLevelErrorColor, DefaultLevelDebugColor, DefaultLevelPanicColor, DefaultLevelDPanicColor, DefaultLevelFatalColor
func (o *Options) WithLevelColor(level slog.Level, color string) *Options {
	o.append(func(opts *Options) {
		opts.LevelColor[level] = color
	})
	return o
}

// WithLevelText 设置日志级别文本
//   - 默认日志级别文本为 DefaultLevelInfoText, DefaultLevelWarnText, DefaultLevelErrorText, DefaultLevelDebugText, DefaultLevelPanicText, DefaultLevelDPanicText, DefaultLevelFatalText
func (o *Options) WithLevelText(level slog.Level, text string) *Options {
	o.append(func(opts *Options) {
		opts.LevelText[level] = text
	})
	return o
}

// WithTimeColor 设置时间颜色
//   - 默认时间颜色为 DefaultTimeColor
func (o *Options) WithTimeColor(color string) *Options {
	o.append(func(opts *Options) {
		opts.TimeColor = color
	})
	return o
}

// WithTimePrefixDelimiter 设置时间前缀分隔符
//   - 默认时间前缀分隔符为 DefaultTimePrefixDelimiter
func (o *Options) WithTimePrefixDelimiter(delimiter string) *Options {
	o.append(func(opts *Options) {
		opts.TimePrefixDelimiter = delimiter
	})
	return o
}

// WithTimePrefixDelimiterColor 设置时间前缀分隔符颜色
//   - 默认时间前缀分隔符颜色为 DefaultTimePrefixDelimiterColor
func (o *Options) WithTimePrefixDelimiterColor(color string) *Options {
	o.append(func(opts *Options) {
		opts.TimePrefixDelimiterColor = color
	})
	return o
}

// WithTimePrefixColor 设置时间前缀颜色
//   - 默认时间前缀颜色为 DefaultTimePrefixColor
func (o *Options) WithTimePrefixColor(color string) *Options {
	o.append(func(opts *Options) {
		opts.TimePrefixColor = color
	})
	return o
}

// WithDisableColor 设置是否禁用颜色
//   - 默认为 DefaultDisableColor
func (o *Options) WithDisableColor(disable bool) *Options {
	o.append(func(opts *Options) {
		opts.DisableColor = disable
	})
	return o
}

// WithKVDelimiter 设置键值对分隔符
//   - 默认键值对分隔符为 DefaultKVDelimiter
func (o *Options) WithKVDelimiter(delimiter string) *Options {
	o.append(func(opts *Options) {
		opts.KVDelimiter = delimiter
	})
	return o
}

// WithErrTraceBeauty 设置是否美化错误堆栈
//   - 默认为 DefaultErrTraceBeauty
//   - 当启用错误堆栈追踪时，一切 error 字段都会转换为包含错误信息和堆栈信息的组信息
func (o *Options) WithErrTraceBeauty(enable bool) *Options {
	o.append(func(opts *Options) {
		opts.ErrTraceBeauty = enable
	})
	return o
}

// WithErrTrace 设置是否显示错误堆栈
//   - 默认为 DefaultErrTrace
//   - 当启用错误堆栈追踪时，一切 error 字段都会转换为包含错误信息和堆栈信息的组信息
func (o *Options) WithErrTrace(enable bool) *Options {
	o.append(func(opts *Options) {
		opts.ErrTrace = enable
	})
	return o
}

// WithLevel 设置日志级别
//   - 默认日志级别为 DefaultLevel
func (o *Options) WithLevel(level slog.Leveler) *Options {
	o.append(func(opts *Options) {
		opts.Level = level
	})
	return o
}

// WithFieldPrefix 为所有字段设置前缀
//   - 默认字段前缀为 DefaultFieldPrefix
//   - 字段前缀为空时，不会添加字段前缀
//   - 假设字段前缀为 "M"，假设原本输出为 "ID=1"，则日志输出为 "MID=1
func (o *Options) WithFieldPrefix(prefix string) *Options {
	o.append(func(opts *Options) {
		opts.FieldPrefix = prefix
	})
	return o
}

// WithCallerFormat 设置调用者信息格式化函数
//   - 默认格式化函数为 CallerBasicFormat
func (o *Options) WithCallerFormat(format func(file string, line int) (repFile, refLine string)) *Options {
	o.append(func(opts *Options) {
		opts.CallerFormat = format
	})
	return o
}

// WithCaller 设置是否显示调用者信息
//   - 默认为 DefaultCaller，且跳过 DefaultCallerSkip 层调用
//   - 当存在多个 skip 参数时，取首个参数
func (o *Options) WithCaller(enable bool, skip ...int) *Options {
	o.append(func(opts *Options) {
		opts.Caller = enable
		if len(skip) > 0 {
			opts.CallerSkip = skip[0]
		}
	})
	return o
}

// WithTimeLayout 设置时间格式化字符串
//   - 默认时间格式化字符串为 DefaultTimeLayout
//   - 假设时间格式化字符串为 "2006-01-02 15:04:05"，则日志输出为 "2020-01-01 00:00:00 ..."
func (o *Options) WithTimeLayout(layout string) *Options {
	o.append(func(opts *Options) {
		opts.TimeLayout = layout
	})
	return o
}

// WithTimePrefix 设置时间前缀
//   - 默认时间前缀为 DefaultTimePrefix
//   - 时间前缀为空时，不会添加时间前缀
//   - 假设时间前缀为 "TIME="，则日志输出为 "TIME=2020-01-01 00:00:00 ..."
func (o *Options) WithTimePrefix(prefix string) *Options {
	o.append(func(opts *Options) {
		opts.TimePrefix = prefix
	})
	return o
}

// With 初始化日志选项
func (o *Options) With(opts ...*Options) *Options {
	for _, opt := range opts {
		o.append(opt.opts...)
	}
	return o
}

// apply 应用日志选项
func (o *Options) apply() *Options {
	for _, opt := range o.opts {
		opt(o)
	}
	return o
}

// append 添加日志选项
func (o *Options) append(opts ...Option) *Options {
	o.opts = append(o.opts, opts...)
	return o
}
