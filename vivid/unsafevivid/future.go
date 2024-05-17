package unsafevivid

import vivid "github.com/kercylan98/minotaur/vivid/vivids"

// NewFuture 创建一个新的 Future 对象，该对象用于异步获取 Actor 的返回值
func NewFuture(ctx vivid.ActorContext, handler func() vivid.Message) vivid.Future {
	f := &future{
		ctx: ctx,
	}
	c := ctx.(*ActorContext)
	var h = func() {
		m := handler()
		f.done = true
		c.Core.Tell(m)
	}
	if err := c.System.gp.Submit(h); err != nil {
		go h()
	}
	return f
}

type future struct {
	ctx  vivid.ActorContext
	done bool
}

func (f *future) IsCompleted() bool {
	return f.done
}
