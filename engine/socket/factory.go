package socket

import "github.com/kercylan98/minotaur/engine/vivid"

// NewFactory 创建一个用于将 socket 转换为 Actor 的 socket 工厂
func NewFactory(system *vivid.ActorSystem) Factory {
	f := &factory{}
	system.ActorOfF(func() vivid.Actor {
		return f
	})
	return f
}

type Factory interface {
	Produce(actor Actor, writer Writer, closer Closer) Socket
}

type factory struct {
	ctx vivid.ActorContext
}

func (f *factory) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		f.ctx = ctx
	case *socket:
		ctx.ActorOfF(func() vivid.Actor {
			return m
		})
		ctx.Reply(nil)
	}
}

func (f *factory) Produce(actor Actor, writer Writer, closer Closer) Socket {
	s := newSocket(actor, writer, closer)
	f.ctx.FutureAsk(f.ctx.Ref(), s).AssertWait()
	return s
}
