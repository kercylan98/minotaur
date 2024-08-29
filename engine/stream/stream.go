// Package stream
package stream

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
type Stream interface {
	Write(packet *Packet) error

	Close() error
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
type WriterCreatedHook interface {
	Stream

	OnWriterCreated(writer Writer)
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
func NewStream(conn Stream, configurator ...Configurator) *Actor {
	c := &Actor{
		config: newConfiguration(),
		stream: conn,
	}

	for _, conf := range configurator {
		conf.Configure(c.config)
	}

	return c
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
type Actor struct {
	config *Configuration
	stream Stream
	writer vivid.ActorRef
}

func (c *Actor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		c.onLaunch(ctx)
	case *vivid.OnTerminate:
		c.onTerminate(ctx)
	case error:
		c.onError(ctx, m)
	default:
		if c.config.performance != nil {
			c.config.performance.Perform(ctx)
		}
	}
}

func (c *Actor) onLaunch(ctx vivid.ActorContext) {
	// 启动写 Actor
	c.writer = ctx.ActorOfF(func() vivid.Actor {
		return newStreamWriter(c.stream)
	})
	if c.config.performance != nil {
		c.config.performance.Perform(ctx)
	}
	ctx.Tell(ctx.Ref(), c.writer)

	if hook, ok := c.stream.(WriterCreatedHook); ok {
		hook.OnWriterCreated(c.writer)
	}
}

func (c *Actor) onError(ctx vivid.ActorContext, err error) {
	if err == nil {
		return
	}
	if c.config.performance != nil {
		c.config.performance.Perform(ctx)
	}
	ctx.ReportAbnormal(err)
}

func (c *Actor) onTerminate(ctx vivid.ActorContext) {
	if c.config.performance != nil {
		c.config.performance.Perform(ctx)
	}
}
