package socket

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"sync/atomic"
)

// NewFactory 创建一个用于将网络连接转换为支持 Actor 功能的 Socket 对象的 Socket 工厂
func NewFactory(system *vivid.ActorSystem, configurator ...FactoryConfigurator) Factory {
	config := NewFactoryConfiguration()
	for _, c := range configurator {
		c.Configure(config)
	}

	f := &factory{
		config: config,
	}
	system.ActorOfF(func() vivid.Actor {
		return f
	})
	return f
}

// Factory 是用于将网络连接转换为支持 Actor 功能的 Socket 对象的 Socket 工厂，它无需被实现，而是由内部的 factory 结构进行实现及维护
type Factory interface {
	// Produce 创建一个支持 Actor 功能的 Socket 对象
	Produce(actor Actor, writer Writer, closer Closer) Socket

	// GetOnlineSocketCount 获取当前在线的 Socket 数量
	GetOnlineSocketCount() int
}

type factory struct {
	config    *FactoryConfiguration // 配置
	ctx       vivid.ActorContext    // 上下文
	onlineNum atomic.Int32          // 在线数量
}

func (f *factory) GetOnlineSocketCount() int {
	return int(f.onlineNum.Load())
}

func (f *factory) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		f.ctx = ctx
	case *socket:
		f.onInitSocket(ctx, m)
	case *vivid.OnTerminated:
		if !m.TerminatedActor.Equal(ctx.Ref()) {
			f.onOnlineStatusChanged(ctx, m)
		}
	}
}

func (f *factory) Produce(actor Actor, writer Writer, closer Closer) Socket {
	s := newSocket(actor, writer, closer)
	f.ctx.FutureAsk(f.ctx.Ref(), s).AssertWait()
	return s
}

func (f *factory) onOnlineStatusChanged(ctx vivid.ActorContext, m *vivid.OnTerminated) {
	f.onlineNum.Add(-1)
}

func (f *factory) onInitSocket(ctx vivid.ActorContext, m *socket) {
	ref := ctx.ActorOfF(func() vivid.Actor {
		return m
	}, func(descriptor *vivid.ActorDescriptor) {
		// 默认值
		descriptor.WithNamePrefix("socket")

		// 自定义值
		if f.config.socketActorDescriptor != nil {
			f.config.socketActorDescriptor.Configure(descriptor)
		}

		// 不可覆盖值
		descriptor.WithSupervisionStrategyProvider(supervision.FunctionalStrategyProvider(func() supervision.Strategy {
			return supervision.StopStrategy()
		}))
	})
	ctx.Watch(ref)
	ctx.Reply(nil)
	f.onlineNum.Add(1)
}
