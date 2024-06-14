package log

import (
	"context"
	"log/slog"
)

// NewSilentLogger 创建一个静默日志记录器，该记录器不会输出任何日志
func NewSilentLogger() *Logger {
	return slog.New(new(SilentHandler))
}

type SilentHandler struct {
}

func (s SilentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return false
}

func (s SilentHandler) Handle(ctx context.Context, record slog.Record) error {
	return nil
}

func (s SilentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return nil
}

func (s SilentHandler) WithGroup(name string) slog.Handler {
	return nil
}
