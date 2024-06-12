package log

import (
	"log/slog"
	"os"
	"sync/atomic"
)

type Logger = slog.Logger

var defaultLogger atomic.Pointer[Logger]

func init() {
	SetDefault(New(NewHandler(os.Stdout, NewDevHandlerOptions())))
}

// New 创建一个新的日志记录器
func New(handler slog.Handler) *Logger {
	return slog.New(handler)
}

// SetDefault 设置默认日志记录器
func SetDefault(logger *Logger) {
	defaultLogger.Store(logger)
}

// GetDefault 获取默认日志记录器
func GetDefault() *Logger {
	l := defaultLogger.Load()
	return l
}

// Debug 使用全局日志记录器在 LevelDebug 级别下记录一条消息
func Debug(msg string, args ...any) {
	l := defaultLogger.Load()
	l.Debug(msg, args...)
}

// Info 使用全局日志记录器在 LevelInfo 级别下记录一条消息
func Info(msg string, args ...any) {
	l := defaultLogger.Load()
	l.Info(msg, args...)
}

// Warn 使用全局日志记录器在 LevelWarn 级别下记录一条消息
func Warn(msg string, args ...any) {
	l := defaultLogger.Load()
	l.Warn(msg, args...)
}

// Error 使用全局日志记录器在 LevelError 级别下记录一条消息
func Error(msg string, args ...any) {
	l := defaultLogger.Load()
	l.Error(msg, args...)
}
