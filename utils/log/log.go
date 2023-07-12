package log

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/kercylan98/minotaur/utils/times"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
)

// NewLog 创建一个日志记录器
func NewLog(options ...Option) *Log {
	log := &Log{
		filename: func(level Level) string {
			return fmt.Sprintf("%s.log", level.String())
		},
		rotateFilename: func(level Level) string {
			return strings.Join([]string{level.String(), "%Y%m%d.log"}, ".")
		},
		levelPartition: defaultLevelPartition,
	}

	for _, option := range options {
		option(log)
	}

	if len(log.rotateOptions) == 0 {
		log.rotateOptions = []rotateLogs.Option{
			rotateLogs.WithMaxAge(times.Week),
			rotateLogs.WithRotationTime(times.Day),
		}
	}

	if len(log.cores) == 0 {
		var encoder = NewEncoder()

		switch log.mode {
		case RunModeDev:
			var partition LevelEnablerFunc = func(lvl Level) bool {
				return true
			}
			log.cores = append(log.cores, zapcore.NewCore(encoder, os.Stdout, partition))
		case RunModeTest, RunModeProd:
			if log.mode == RunModeTest {
				infoRotate, err := rotateLogs.New(
					filepath.Join(log.rotateLogDir, log.rotateFilename(InfoLevel)),
					append([]rotateLogs.Option{rotateLogs.WithLinkName(filepath.Join(log.logDir, log.filename(InfoLevel)))}, log.rotateOptions...)...,
				)
				if err != nil {
					panic(err)
				}
				errRotate, err := rotateLogs.New(
					filepath.Join(log.rotateLogDir, log.rotateFilename(ErrorLevel)),
					append([]rotateLogs.Option{rotateLogs.WithLinkName(filepath.Join(log.logDir, log.filename(ErrorLevel)))}, log.rotateOptions...)...,
				)
				if err != nil {
					panic(err)
				}
				if log.logDir != str.None {
					log.cores = append(log.cores, zapcore.NewCore(encoder, zapcore.AddSync(infoRotate), LevelEnablerFunc(func(lvl Level) bool { return lvl < ErrorLevel })))
					log.cores = append(log.cores, zapcore.NewCore(encoder, zapcore.AddSync(errRotate), LevelEnablerFunc(func(lvl Level) bool { return lvl >= ErrorLevel })))
					log.cores = append(log.cores, zapcore.NewCore(encoder, os.Stdout, LevelEnablerFunc(func(lvl Level) bool { return lvl < ErrorLevel })))
					log.cores = append(log.cores, zapcore.NewCore(encoder, os.Stdout, LevelEnablerFunc(func(lvl Level) bool { return lvl >= ErrorLevel })))
				}
			} else {
				infoRotate, err := rotateLogs.New(
					filepath.Join(log.rotateLogDir, log.rotateFilename(InfoLevel)),
					append([]rotateLogs.Option{rotateLogs.WithLinkName(filepath.Join(log.logDir, log.filename(InfoLevel)))}, log.rotateOptions...)...,
				)
				if err != nil {
					panic(err)
				}
				errRotate, err := rotateLogs.New(
					filepath.Join(log.rotateLogDir, log.rotateFilename(ErrorLevel)),
					append([]rotateLogs.Option{rotateLogs.WithLinkName(filepath.Join(log.logDir, log.filename(ErrorLevel)))}, log.rotateOptions...)...,
				)
				if err != nil {
					panic(err)
				}
				if log.logDir != str.None {
					log.cores = append(log.cores, zapcore.NewCore(encoder, zapcore.AddSync(infoRotate), LevelEnablerFunc(func(lvl Level) bool { return lvl == InfoLevel })))
					log.cores = append(log.cores, zapcore.NewCore(encoder, zapcore.AddSync(errRotate), LevelEnablerFunc(func(lvl Level) bool { return lvl >= ErrorLevel })))
				}
			}
		}
	}

	log.zap = zap.New(zapcore.NewTee(log.cores...), zap.AddCaller(), zap.AddCallerSkip(1))
	log.sugar = log.zap.Sugar()
	return log
}

type Log struct {
	zap            *zap.Logger
	sugar          *zap.SugaredLogger
	filename       func(level Level) string
	rotateFilename func(level Level) string
	rotateOptions  []rotateLogs.Option
	levelPartition map[Level]func() LevelEnablerFunc
	cores          []Core
	mode           RunMode
	logDir         string
	rotateLogDir   string
}

func (slf *Log) Debugf(format string, args ...interface{}) {
	slf.sugar.Debugf(format, args...)
}

func (slf *Log) Infof(format string, args ...interface{}) {
	slf.sugar.Infof(format, args...)
}

func (slf *Log) Warnf(format string, args ...interface{}) {
	slf.sugar.Warnf(format, args...)
}

func (slf *Log) Errorf(format string, args ...interface{}) {
	slf.sugar.Errorf(format, args...)
}

func (slf *Log) Fatalf(format string, args ...interface{}) {
	slf.sugar.Fatalf(format, args...)
}

func (slf *Log) Printf(format string, args ...interface{}) {
	slf.sugar.Infof(format, args...)
}

// Debug 在 DebugLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func (slf *Log) Debug(msg string, fields ...Field) {
	slf.zap.Debug(msg, fields...)
}

// Info 在 InfoLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func (slf *Log) Info(msg string, fields ...Field) {
	slf.zap.Info(msg, fields...)
}

// Warn 在 WarnLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func (slf *Log) Warn(msg string, fields ...Field) {
	slf.zap.Warn(msg, fields...)
}

// Error 在 ErrorLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func (slf *Log) Error(msg string, fields ...Field) {
	slf.zap.Error(msg, fields...)
}

// DPanic 在 DPanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
//   - 如果记录器处于开发模式，它就会出现 panic（DPanic 的意思是“development panic”）。这对于捕获可恢复但不应该发生的错误很有用
func (slf *Log) DPanic(msg string, fields ...Field) {
	slf.zap.DPanic(msg, fields...)
}

// Panic 在 PanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
//   - 即使禁用了 PanicLevel 的日志记录，记录器也会出现 panic
func (slf *Log) Panic(msg string, fields ...Field) {
	slf.zap.Panic(msg, fields...)
}

// Fatal 在 FatalLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
//   - 然后记录器调用 os.Exit(1)，即使 FatalLevel 的日志记录被禁用
func (slf *Log) Fatal(msg string, fields ...Field) {
	slf.zap.Fatal(msg, fields...)
}
