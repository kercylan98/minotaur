package vivid

type Receiver interface {
	Tell(v Message, opts ...MessageOption) error
	Ask(v Message, opts ...MessageOption) (any, error)
}

type Future interface {
	// IsCompleted 判断 Future 是否已完成
	IsCompleted() bool
}

// NewFuture 创建一个新的 Future 对象，该对象用于异步获取 Actor 的返回值
func NewFuture(ctx ActorContext, handler func() Message) Future {
	f := &future{
		ctx: ctx,
	}
	c := ctx.(*actorContext)
	var h = func() {
		m := handler()
		f.done = true
		c.core.Tell(m)
	}
	if err := c.system.gp.Submit(h); err != nil {
		go h()
	}
	return f
}

type future struct {
	ctx  ActorContext
	done bool
}

func (f *future) IsCompleted() bool {
	return f.done
}
