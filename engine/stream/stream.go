package stream

import "github.com/kercylan98/minotaur/engine/vivid"

type Stream interface {
	Write(packet *Packet) error

	Close() error
}

func New(conn Stream, configurator ...Configurator) *Actor {
	c := &Actor{
		config: newConfiguration(),
		stream: conn,
	}

	for _, conf := range configurator {
		conf.Configure(c.config)
	}

	return c
}

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
	_ = c.stream.Close()
}
