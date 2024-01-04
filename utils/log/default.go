package log

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

var logger atomic.Pointer[Logger]

func init() {
	logger.Store(NewLogger())
}

// Default 获取默认的日志记录器
func Default() *Logger {
	return logger.Load()
}

// SetDefault 设置默认的日志记录器
func SetDefault(l *Logger) {
	logger.Store(l)
}

// Debug 在 DebugLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func Debug(msg string, args ...any) {
	handle(DebugLevel, msg, args...)
}

// Info 在 InfoLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func Info(msg string, args ...any) {
	handle(InfoLevel, msg, args...)
}

// Warn 在 WarnLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func Warn(msg string, args ...any) {
	handle(WarnLevel, msg, args...)
}

// Error 在 ErrorLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func Error(msg string, args ...any) {
	handle(ErrorLevel, msg, args...)
}

// DPanic 在 DPanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
//   - 如果记录器处于开发模式，它就会出现 panic（DPanic 的意思是“development panic”）。这对于捕获可恢复但不应该发生的错误很有用
func DPanic(msg string, args ...any) {
	handle(DPanicLevel, msg, args...)
}

// Panic 在 PanicLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
//   - 即使禁用了 PanicLevel 的日志记录，记录器也会出现 panic
func Panic(msg string, args ...any) {
	handle(PanicLevel, msg, args...)
	panic(errors.New(msg))
}

// Fatal 在 FatalLevel 记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
//   - 然后记录器调用 os.Exit(1)，即使 FatalLevel 的日志记录被禁用
func Fatal(msg string, args ...any) {
	handle(FatalLevel, msg, args...)
	os.Exit(1)
}

// handle 在指定的级别记录一条消息。该消息包括在日志站点传递的任何字段以及记录器上累积的任何字段
func handle(level slog.Level, msg string, args ...any) {
	d := Default()
	pcs := make([]uintptr, 1)
	runtime.CallersFrames(pcs[:runtime.Callers(d.opts.CallerSkip, pcs)])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(args...)
	_ = d.Handler().Handle(context.Background(), r)
	if level == DPanicLevel && d.opts.DevMode {
		panic(errors.New(msg))
	}
}
