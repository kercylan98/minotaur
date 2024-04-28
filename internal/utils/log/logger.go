package log

import (
	"log/slog"
	"os"
)

// NewLogger 创建一个新的日志记录器
func NewLogger(handlers ...slog.Handler) *Logger {
	var h slog.Handler
	switch len(handlers) {
	case 0:
		h = NewHandler(os.Stdout, nil)
	case 1:
		h = handlers[0]
	default:
		h = NewMultiHandler(handlers...)
	}
	return &Logger{
		Logger: slog.New(h),
	}
}

type Logger struct {
	*slog.Logger
}
