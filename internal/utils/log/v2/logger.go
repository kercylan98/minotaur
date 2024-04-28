package log

import (
	"log/slog"
	"os"
	"sync/atomic"
)

var logger = func() *atomic.Pointer[Logger] {
	var p atomic.Pointer[Logger]
	p.Store(slog.New(NewHandler(os.Stdout)))
	return &p
}()

// NewLogger 创建一个新的日志记录器
func NewLogger(handler Handler) *Logger {
	return slog.New(handler)
}

// NewDefaultLogger 创建一个新的默认日志记录器
func NewDefaultLogger() *Logger {
	return NewLogger(NewHandler(os.Stdout))
}

// GetLogger 并发安全的获取当前全局日志记录器
func GetLogger() *Logger {
	l := logger.Load()
	if h := cloneHandler(l.Handler()); h != nil {
		return NewLogger(h)
	}
	return l
}

// SetLogger 并发安全的设置全局日志记录器
func SetLogger(l *Logger) {
	logger.Store(l)
}

// ResetLogger 并发安全的重置全局日志记录器
func ResetLogger() {
	logger.Store(slog.New(NewHandler(os.Stdout)))
}

// Debug 使用全局日志记录器在 LevelDebug 级别下记录一条消息
func Debug(msg string, args ...any) {
	logger.Load().Debug(msg, args...)
}

// Info 使用全局日志记录器在 LevelInfo 级别下记录一条消息
func Info(msg string, args ...any) {
	logger.Load().Info(msg, args...)
}

// Warn 使用全局日志记录器在 LevelWarn 级别下记录一条消息
func Warn(msg string, args ...any) {
	logger.Load().Warn(msg, args...)
}

// Error 使用全局日志记录器在 LevelError 级别下记录一条消息
func Error(msg string, args ...any) {
	logger.Load().Error(msg, args...)
}

// Log 按照指定级别记录日志消息
func Log(level Level, msg string, args ...any) {
	switch level {
	case LevelDebug:
		logger.Load().Debug(msg, args...)
	case LevelInfo:
		logger.Load().Info(msg, args...)
	case LevelWarn:
		logger.Load().Warn(msg, args...)
	case LevelError:
		logger.Load().Error(msg, args...)
	default:
	}
}

func cloneHandler(h Handler) Handler {
	switch h := h.(type) {
	case *MinotaurHandler:
		cloneHandler := h.clone()
		cloneHandler.GetOptions().WithCallerSkip(-1)
		return cloneHandler
	case *MultiHandler:
		var handlers = make([]Handler, 0, len(h.handlers))
		for _, handler := range h.handlers {
			handlers = append(handlers, cloneHandler(handler))
		}
		return NewMultiHandler(handlers...)
	}
	return nil
}
