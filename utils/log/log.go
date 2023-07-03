package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"runtime/debug"
	"time"
)

var (
	logger      *zap.Logger
	prod        bool
	logPath     string
	logDevWrite bool
	logTime     = 7
)

func init() {
	logger = newLogger()
	if prod && len(logPath) == 0 {
		Warn("Logger", zap.String("Tip", "in production mode, if the log file output directory is not set, only the console will be output"))
	}
}

func newLogger() *zap.Logger {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(time.DateTime))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.FatalLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	var cores zapcore.Core

	if !prod {
		if len(logPath) > 0 && logDevWrite {
			infoWriter := getWriter(fmt.Sprintf("%s/info.log", logPath), logTime)
			errorWriter := getWriter(fmt.Sprintf("%s/error.log", logPath), logTime)
			cores = zapcore.NewTee(
				zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
				zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
				zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), debugLevel),
			)
		} else {
			cores = zapcore.NewTee(
				zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), debugLevel),
			)
		}
	} else {
		if len(logPath) == 0 {
			cores = zapcore.NewTee(
				zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), debugLevel),
			)
		} else {
			infoWriter := getWriter(fmt.Sprintf("%s/info.log", logPath), logTime)
			errorWriter := getWriter(fmt.Sprintf("%s/error.log", logPath), logTime)
			cores = zapcore.NewTee(
				zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
				zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
				zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), errorLevel),
			)
		}
	}

	return zap.New(cores, zap.AddCaller(), zap.AddCallerSkip(1))
}

func getWriter(filename string, times int) io.Writer {
	hook, err := rotatelogs.New(
		filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(times)),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

func Logger() *zap.Logger {
	return logger
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
	fmt.Println(string(debug.Stack()))
}

func ErrorHideStack(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// ErrorWithStack 通过额外的堆栈信息打印错误日志
func ErrorWithStack(msg, stack string, fields ...zap.Field) {
	logger.Error(msg, fields...)
	var stackMerge string
	if len(stack) > 0 {
		stackMerge = stack
	}
	stackMerge += string(debug.Stack())
	fmt.Println(stackMerge)
}

// SetProd 设置生产环境模式
func SetProd(isProd bool) {
	if prod == isProd {
		return
	}
	prod = isProd
	if logger != nil {
		_ = logger.Sync()
	}
	logger = newLogger()
}

// SetLogDir 设置日志输出目录
func SetLogDir(dir string) {
	logPath = dir
	if logger != nil {
		_ = logger.Sync()
	}
	logger = newLogger()
}

// SetWriteFileWithDev 设置开发环境下写入文件
func SetWriteFileWithDev(isWrite bool) {
	if isWrite == logDevWrite {
		return
	}
	logDevWrite = isWrite
	if logger != nil {
		_ = logger.Sync()
	}
	logger = newLogger()
}

// SetLogRotate 设置日志切割时间
func SetLogRotate(t int) {
	logTime = t
	if logger != nil {
		_ = logger.Sync()
	}
	logger = newLogger()
}
