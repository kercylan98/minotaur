package socket

import "github.com/kercylan98/minotaur/engine/vivid"

// NewFactory 创建一个用于将网络连接转换为支持 Actor 功能的 Socket 对象的 Socket 工厂
func NewFactory(system *vivid.ActorSystem) Factory {
	f := &factory{}
	system.ActorOfF(func() vivid.Actor {
		return f
	})
	return f
}

// Factory 是用于将网络连接转换为支持 Actor 功能的 Socket 对象的 Socket 工厂，它无需被实现，而是由内部的 factory 结构进行实现及维护
type Factory interface {
	// Produce 创建一个支持 Actor 功能的 Socket 对象
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
