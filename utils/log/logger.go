package log

import (
	"log/slog"
	"os"
)

// NewLogger 创建一个新的日志记录器
func NewLogger(options ...*Options) *Logger {
	opts := NewOptions().With(options...)
	return &Logger{
		Logger: slog.New(NewHandler(os.Stdout, opts)),
		opts:   opts,
	}
}

type Logger struct {
	*slog.Logger
	opts *Options
}
