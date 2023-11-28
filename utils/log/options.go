package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

type (
	Config           = zap.Config
	Level            = zapcore.Level
	LevelEncoder     = zapcore.LevelEncoder
	TimeEncoder      = zapcore.TimeEncoder
	DurationEncoder  = zapcore.DurationEncoder
	CallerEncoder    = zapcore.CallerEncoder
	NameEncoder      = zapcore.NameEncoder
	ReflectedEncoder = zapcore.ReflectedEncoder
	WriteSyncer      = zapcore.WriteSyncer
	LoggerOption     = zap.Option
	Core             = zapcore.Core
	LevelEnabler     = zapcore.LevelEnabler
	Option           func(config *Config)
)

func Default(opts ...Option) *Encoder {
	config := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableStacktrace: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "Time",
			LevelKey:      "Level",
			NameKey:       "Name",
			CallerKey:     "Caller",
			MessageKey:    "Msg",
			StacktraceKey: "Stack",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(t.Format(time.DateTime))
			},
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	// 应用选项
	for _, opt := range opts {
		opt(config)
	}

	if len(config.Encoding) == 0 {
		if config.Development {
			config.Encoding = "console"
		} else {
			config.Encoding = "json"
		}
	}

	var encoder = new(Encoder)
	encoder.conf = config
	switch config.Encoding {
	case "console":
		encoder.e = zapcore.NewConsoleEncoder(config.EncoderConfig)
	case "json":
		encoder.e = zapcore.NewJSONEncoder(config.EncoderConfig)
	default:
		panic("unknown encoding")
	}

	return encoder
}

// WithLevel 设置日志级别
func WithLevel(level Level) Option {
	return func(config *Config) {
		config.Level.SetLevel(level)
	}
}

// WithDevelopment 设置是否为开发模式
func WithDevelopment(development bool) Option {
	return func(config *Config) {
		config.Development = development
	}
}

// WithDisableCaller 设置是否禁用调用者
func WithDisableCaller(disableCaller bool) Option {
	return func(config *Config) {
		config.DisableCaller = disableCaller
	}
}

// WithDisableStacktrace 设置是否禁用堆栈跟踪
func WithDisableStacktrace(disableStacktrace bool) Option {
	return func(config *Config) {
		config.DisableStacktrace = disableStacktrace
	}
}

// WithSampling 设置采样策略
func WithSampling(sampling *zap.SamplingConfig) Option {
	return func(config *Config) {
		config.Sampling = sampling
	}
}

// WithEncoding 设置编码器
func WithEncoding(encoding string) Option {
	return func(config *Config) {
		config.Encoding = encoding
	}
}

// WithEncoderMessageKey 设置消息键
func WithEncoderMessageKey(encoderMessageKey string) Option {
	return func(config *Config) {
		config.EncoderConfig.MessageKey = encoderMessageKey
	}
}

// WithEncoderLevelKey 设置级别键
func WithEncoderLevelKey(encoderLevelKey string) Option {
	return func(config *Config) {
		config.EncoderConfig.LevelKey = encoderLevelKey
	}
}

// WithEncoderTimeKey 设置时间键
func WithEncoderTimeKey(encoderTimeKey string) Option {
	return func(config *Config) {
		config.EncoderConfig.TimeKey = encoderTimeKey
	}
}

// WithEncoderNameKey 设置名称键
func WithEncoderNameKey(encoderNameKey string) Option {
	return func(config *Config) {
		config.EncoderConfig.NameKey = encoderNameKey
	}
}

// WithEncoderCallerKey 设置调用者键
func WithEncoderCallerKey(encoderCallerKey string) Option {
	return func(config *Config) {
		config.EncoderConfig.CallerKey = encoderCallerKey
	}
}

// WithEncoderFunctionKey 设置函数键
func WithEncoderFunctionKey(encoderFunctionKey string) Option {
	return func(config *Config) {
		config.EncoderConfig.FunctionKey = encoderFunctionKey
	}
}

// WithEncoderStacktraceKey 设置堆栈跟踪键
func WithEncoderStacktraceKey(encoderStacktraceKey string) Option {
	return func(config *Config) {
		config.EncoderConfig.StacktraceKey = encoderStacktraceKey
	}
}

// WithEncoderLineEnding 设置行尾
func WithEncoderLineEnding(encoderLineEnding string) Option {
	return func(config *Config) {
		config.EncoderConfig.LineEnding = encoderLineEnding
	}
}

// WithEncoderLevel 设置级别编码器
func WithEncoderLevel(encoderLevel LevelEncoder) Option {
	return func(config *Config) {
		config.EncoderConfig.EncodeLevel = encoderLevel
	}
}

// WithEncoderTime 设置时间编码器
func WithEncoderTime(encoderTime TimeEncoder) Option {
	return func(config *Config) {
		config.EncoderConfig.EncodeTime = encoderTime
	}
}

// WithEncoderDuration 设置持续时间编码器
func WithEncoderDuration(encoderDuration DurationEncoder) Option {
	return func(config *Config) {
		config.EncoderConfig.EncodeDuration = encoderDuration
	}
}

// WithEncoderCaller 设置调用者编码器
func WithEncoderCaller(encoderCaller CallerEncoder) Option {
	return func(config *Config) {
		config.EncoderConfig.EncodeCaller = encoderCaller
	}
}

// WithEncoderName 设置名称编码器
func WithEncoderName(encoderName NameEncoder) Option {
	return func(config *Config) {
		config.EncoderConfig.EncodeName = encoderName
	}
}

// WithEncoderNewReflectedEncoder 设置反射编码器
func WithEncoderNewReflectedEncoder(encoderNewReflectedEncoder func(io.Writer) ReflectedEncoder) Option {
	return func(config *Config) {
		config.EncoderConfig.NewReflectedEncoder = encoderNewReflectedEncoder
	}
}

// WithOutputPaths 设置输出路径
func WithOutputPaths(outputPaths ...string) Option {
	return func(config *Config) {
		config.OutputPaths = outputPaths
	}
}

// WithErrorOutputPaths 设置错误输出路径
func WithErrorOutputPaths(errorOutputPaths ...string) Option {
	return func(config *Config) {
		config.ErrorOutputPaths = errorOutputPaths
	}
}

// WithInitialFields 设置初始字段
func WithInitialFields(initialFields map[string]interface{}) Option {
	return func(config *Config) {
		config.InitialFields = initialFields
	}
}
