package log

import (
	"context"
	"github.com/kercylan98/minotaur/utils/super"
	"log/slog"
)

// NewMultiHandler 创建一个新的多处理程序
func NewMultiHandler(handlers ...slog.Handler) slog.Handler {
	return &MultiHandler{
		handlers: handlers,
	}
}

type MultiHandler struct {
	handlers []slog.Handler
}

func (h MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, level) {
			return true
		}
	}

	return false
}

func (h MultiHandler) Handle(ctx context.Context, record slog.Record) (err error) {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, record.Level) {
			err = func() error {
				defer func() {
					err = super.RecoverTransform(recover())
				}()
				return h.handlers[i].Handle(ctx, record.Clone())
			}()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (h MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var handlers = make([]slog.Handler, len(h.handlers))
	for i, s := range h.handlers {
		handlers[i] = s.WithAttrs(attrs)
	}
	return NewMultiHandler(handlers...)
}

func (h MultiHandler) WithGroup(name string) slog.Handler {
	var handlers = make([]slog.Handler, len(h.handlers))
	for i, s := range h.handlers {
		handlers[i] = s.WithGroup(name)
	}
	return NewMultiHandler(handlers...)
}
