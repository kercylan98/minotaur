package log

import (
	"log/slog"
	"sync/atomic"
)

type Logger = slog.Logger

var defaultLogger atomic.Pointer[Logger]

// SetDefault 设置默认日志记录器
func SetDefault(logger *Logger) {
	defaultLogger.Store(logger)
}

// GetDefault 获取默认日志记录器
func GetDefault() *Logger {
	l := defaultLogger.Load()
	if l == nil {
		return slog.Default()
	}
	return l
}

// Debug 使用全局日志记录器在 LevelDebug 级别下记录一条消息
func Debug(msg string, args ...any) {
	l := defaultLogger.Load()
	if l == nil {
		l = slog.Default()
	}
	l.Debug(msg, args...)
}

// Info 使用全局日志记录器在 LevelInfo 级别下记录一条消息
func Info(msg string, args ...any) {
	l := defaultLogger.Load()
	if l == nil {
		l = slog.Default()
	}
	l.Info(msg, args...)
}

// Warn 使用全局日志记录器在 LevelWarn 级别下记录一条消息
func Warn(msg string, args ...any) {
	l := defaultLogger.Load()
	if l == nil {
		l = slog.Default()
	}
	l.Warn(msg, args...)
}

// Error 使用全局日志记录器在 LevelError 级别下记录一条消息
func Error(msg string, args ...any) {
	l := defaultLogger.Load()
	if l == nil {
		l = slog.Default()
	}
	l.Error(msg, args...)
}
