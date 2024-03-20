package super

import "context"

func WithCancelContext(ctx context.Context) *CancelContext {
	ctx, cancel := context.WithCancel(ctx)
	return &CancelContext{
		Context: ctx,
		cancel:  cancel,
	}
}

type CancelContext struct {
	context.Context
	cancel func()
}

func (c *CancelContext) Cancel() {
	c.cancel()
}
