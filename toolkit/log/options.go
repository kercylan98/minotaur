package log

import (
	""
"github.com/kercylan98/minotaur/toolkit/collection"
"path/filepath"
"strconv"
"sync"
"time"
)

const basicCallSkip = 8

func DefaultOptions() *Options {
	return &Options{
		timeLayout: time.DateTime,
		callerFormatter: func(file string, line int) (repFile, repLine string) {
			return filepath.Base(file), strconv.Itoa(line)
		},
		level: LevelDebug,
		levelText: map[Level]string{
			LevelDebug: "DBG",
			LevelInfo:  "INF",
			LevelWarn:  "WAR",
			LevelError: "ERR",
		},
		levelColor: map[Level]*Color{
			LevelDebug: NewColor(ColorFgCyan),
			LevelInfo:  NewColor(ColorFgGreen),
			LevelWarn:  NewColor(ColorFgHiYellow),
			LevelError: NewColor(ColorFgHiRed),
		},
		callerSkip: basicCallSkip,
		keyColor: map[AttrType]*Color{
			AttrTypeField: NewColor(ColorFgWhite),
			AttrTypeError: NewColor(ColorFgRed),
		},
		delimiterText: map[AttrType]string{
			AttrTypeField: "=",
		},
		valueColor: map[AttrType]*Color{
			AttrTypeCaller:  NewColor(ColorFgHiCyan),
			AttrTypeMessage: NewColor(ColorFgWhite, ColorBold),
			AttrTypeTrace:   NewColor(ColorFgWhite, ColorFaint),
		},
	}
}

type Options struct {
	rw             sync.RWMutex
	level          Level // 日志级别
	timeLayout     string
	keyText        map[AttrType]string // 特定属性类型的前缀字符串
	keyColor       map[AttrType]*Color // 特定属性类型的前缀颜色
	delimiterText  map[AttrType]string // 特定属性类型的分隔符字符串
	delimiterColor map[AttrType]*Color // 特定属性类型的分隔符颜色
	valueColor     map[AttrType]*Color // 特定属性类型的值颜色
	levelText      map[Level]string    // 特定级别的字符串
	levelColor     map[Level]*Color    // 特定级别的颜色

	disabledColor bool // 是否禁用颜色

	callerSkip int // 调用者跳过数量

	disabledCaller   bool                                                  // 是否禁用调用者
	callerFormatter  func(file string, line int) (repFile, repLine string) // 调用者格式化函数
	stackTrace       map[Level]bool                                        // 是否开启特定级别的堆栈追踪
	stackTraceBeauty map[Level]bool                                        // 是否开启特定级别的堆栈追踪美化
}

func (opt *Options) Apply(opts ...*Options) *Options {
	opt.rw.Lock()
	defer opt.rw.Unlock()
	for _, o := range opts {
		o.getMany(func(o *Options) {
			opt.level = o.level
			opt.keyText = collection.CloneMap(o.keyText)
			opt.delimiterText = collection.CloneMap(o.delimiterText)
			opt.keyColor = map[AttrType]*Color{}
			for attrType, color := range o.keyColor {
				opt.keyColor[attrType] = color.clone()
			}
			opt.delimiterColor = map[AttrType]*Color{}
			for attrType, color := range o.delimiterColor {
				opt.delimiterColor[attrType] = color.clone()
			}
			opt.valueColor = map[AttrType]*Color{}
			for attrType, color := range o.valueColor {
				opt.valueColor[attrType] = color.clone()
			}
			opt.timeLayout = o.timeLayout
			opt.disabledColor = o.disabledColor
			opt.levelColor = make(map[Level]*Color)
			for level, color := range o.levelColor {
				opt.levelColor[level] = color.clone()
			}
			opt.callerSkip = o.callerSkip
			opt.disabledCaller = o.disabledCaller
			opt.levelText = make(map[Level]string)
			for level, text := range o.levelText {
				opt.levelText[level] = text
			}
			opt.callerFormatter = o.callerFormatter
			opt.stackTrace = collection.CloneMap(o.stackTrace)
			opt.stackTraceBeauty = collection.CloneMap(o.stackTraceBeauty)
		})
	}
	return opt
}

// WithStackTraceBeauty 设置堆栈追踪美化
//   - 该函数支持运行时设置
func (opt *Options) WithStackTraceBeauty(level Level, enable bool) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		if opt.stackTraceBeauty == nil {
			opt.stackTraceBeauty = make(map[Level]bool)
		}

		opt.stackTraceBeauty[level] = enable
	})
}

// GetStackTraceBeauty 获取堆栈追踪是否美化
func (opt *Options) GetStackTraceBeauty(level Level) bool {
	return getOptionsValue(opt, func(opt *Options) bool {
		return opt.stackTraceBeauty[level]
	})
}

// WithStackTrace 设置堆栈追踪，当日志记录器中包含 error 时，将会打印堆栈信息
//   - 该函数支持运行时设置
func (opt *Options) WithStackTrace(level Level, enable bool) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		if opt.stackTrace == nil {
			opt.stackTrace = make(map[Level]bool)
		}

		opt.stackTrace[level] = enable
	})
}

// GetStackTrace 获取堆栈追踪
func (opt *Options) GetStackTrace(level Level) bool {
	return getOptionsValue(opt, func(opt *Options) bool {
		return opt.stackTrace[level]
	})
}

// WithCallerFormatter 设置调用者格式化函数
//   - 该函数支持运行时设置
func (opt *Options) WithCallerFormatter(formatter func(file string, line int) (repFile, repLine string)) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.callerFormatter = formatter
	})
}

// GetCallerFormatter 获取调用者格式化函数
func (opt *Options) GetCallerFormatter() func(file string, line int) (repFile, repLine string) {
	return getOptionsValue(opt, func(opt *Options) func(file string, line int) (repFile, repLine string) {
		return opt.callerFormatter
	})
}

// WithDisableCaller 设置是否禁用调用者
//   - 该函数支持运行时设置
func (opt *Options) WithDisableCaller(disable bool) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.disabledCaller = disable
	})
}

// IsDisabledCaller 获取是否已经禁用调用者
func (opt *Options) IsDisabledCaller() bool {
	return getOptionsValue(opt, func(opt *Options) bool {
		return opt.disabledCaller
	})
}

// WithAttrPrefix 设置属性前缀
//   - 该函数支持运行时设置
func (opt *Options) WithAttrPrefix(attrType AttrType, prefix string) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		if opt.keyText == nil {
			opt.keyText = map[AttrType]string{}
		}
		opt.keyText[attrType] = prefix
	})
}

// GetAttrPrefix 获取属性前缀
func (opt *Options) GetAttrPrefix(attrType AttrType) string {
	return getOptionsValue(opt, func(opt *Options) string {
		return opt.keyText[attrType]
	})
}

// WithAttrDelimiter 设置属性分隔符
//   - 该函数支持运行时设置
func (opt *Options) WithAttrDelimiter(attrType AttrType, delimiter string) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		if opt.delimiterText == nil {
			opt.delimiterText = map[AttrType]string{}
		}
		opt.delimiterText[attrType] = delimiter
	})
}

// GetAttrDelimiter 获取属性分隔符
func (opt *Options) GetAttrDelimiter(attrType AttrType) string {
	return getOptionsValue(opt, func(opt *Options) string {
		return opt.delimiterText[attrType]
	})
}

// WithAttrPrefixColor 设置属性前缀颜色
//   - 该函数支持运行时设置
func (opt *Options) WithAttrPrefixColor(attrType AttrType, color *Color) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		if opt.keyColor == nil {
			opt.keyColor = map[AttrType]*Color{}
		}
		opt.keyColor[attrType] = color
	})
}

// GetAttrPrefixColor 获取属性前缀颜色
func (opt *Options) GetAttrPrefixColor(attrType AttrType) *Color {
	return getOptionsValue(opt, func(opt *Options) *Color {
		return opt.keyColor[attrType]
	})
}

// WithAttrDelimiterColor 设置属性分隔符颜色
//   - 该函数支持运行时设置
func (opt *Options) WithAttrDelimiterColor(attrType AttrType, color *Color) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		if opt.delimiterColor == nil {
			opt.delimiterColor = map[AttrType]*Color{}
		}
		opt.delimiterColor[attrType] = color
	})
}

// GetAttrDelimiterColor 获取属性分隔符颜色
func (opt *Options) GetAttrDelimiterColor(attrType AttrType) *Color {
	return getOptionsValue(opt, func(opt *Options) *Color {
		return opt.delimiterColor[attrType]
	})
}

// WithAttrTextColor 设置属性文本颜色
//   - 该函数支持运行时设置
func (opt *Options) WithAttrTextColor(attrType AttrType, color *Color) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		if opt.valueColor == nil {
			opt.valueColor = map[AttrType]*Color{}
		}
		opt.valueColor[attrType] = color
	})
}

// GetAttrTextColor 获取属性前缀颜色
func (opt *Options) GetAttrTextColor(attrType AttrType) *Color {
	return getOptionsValue(opt, func(opt *Options) *Color {
		return opt.valueColor[attrType]
	})
}

// WithCallerSkip 设置调用者跳过数量
//   - 该函数支持运行时设置
func (opt *Options) WithCallerSkip(skip int) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.callerSkip = basicCallSkip + skip
	})
}

// GetCallerSkip 获取调用者跳过数量
func (opt *Options) GetCallerSkip() int {
	return getOptionsValue(opt, func(opt *Options) int {
		return basicCallSkip - opt.callerSkip
	})
}

// WithLevelText 设置日志级别文本
//   - 该函数支持运行时设置
func (opt *Options) WithLevelText(level Level, text string) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		if opt.levelText == nil {
			opt.levelText = make(map[Level]string)
		}
		opt.levelText[level] = text
	})
}

// GetLevelText 获取日志级别文本
func (opt *Options) GetLevelText(level Level) string {
	return getOptionsValue(opt, func(opt *Options) string {
		return opt.levelText[level]
	})
}

// WithLevelColor 设置日志级别颜色
//   - 该函数支持运行时设置
func (opt *Options) WithLevelColor(level Level, color *Color) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		if opt.levelColor == nil {
			opt.levelColor = make(map[Level]*Color)
		}
		opt.levelColor[level] = color
	})
}

// GetLevelColor 获取日志级别颜色
func (opt *Options) GetLevelColor(level Level) *Color {
	return getOptionsValue(opt, func(opt *Options) *Color {
		return opt.levelColor[level]
	})
}

// WithDisableColor 设置禁用颜色
//   - 该函数支持运行时设置
func (opt *Options) WithDisableColor(disable bool) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.disabledColor = disable
	})
}

// IsDisabledColor 获取是否已经禁用颜色
func (opt *Options) IsDisabledColor() bool {
	return getOptionsValue(opt, func(opt *Options) bool {
		return opt.disabledColor
	})
}

// WithLevel 设置日志级别
//   - 该函数支持运行时设置
func (opt *Options) WithLevel(level Level) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.level = level
	})
}

// GetLevel 获取当前日志级别
func (opt *Options) GetLevel() Level {
	return getOptionsValue(opt, func(opt *Options) Level {
		return opt.level
	})
}

// WithTimeLayout 设置日志的时间布局，默认为 time.DateTime
//   - 该函数支持运行时设置
func (opt *Options) WithTimeLayout(layout string) *Options {
	return opt.modifyOptionsValue(func(opt *Options) {
		opt.timeLayout = layout
	})
}

// GetTimeLayout 获取当前日志的时间布局
func (opt *Options) GetTimeLayout() string {
	return getOptionsValue(opt, func(opt *Options) string {
		return opt.timeLayout
	})
}

func (opt *Options) modifyOptionsValue(handler func(opt *Options)) *Options {
	opt.rw.Lock()
	handler(opt)
	opt.rw.Unlock()
	return opt
}

func (opt *Options) getMany(handler func(opt *Options)) {
	opt.rw.RLock()
	defer opt.rw.RUnlock()
	handler(opt)
}

func getOptionsValue[V any](opt *Options, handler func(opt *Options) V) V {
	opt.rw.RLock()
	defer opt.rw.RUnlock()
	return handler(opt)
}
