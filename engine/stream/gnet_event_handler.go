package stream

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"github.com/panjf2000/gnet/v2"
	"time"
)

var _ gnet.EventHandler = (*GNETEventHandler)(nil)

func NewGNETEventHandler(system *vivid.ActorSystem, configurator ...Configurator) *GNETEventHandler {
	return &GNETEventHandler{
		system:       system,
		configurator: configurator,
	}
}

type GNETEventHandler struct {
	system       *vivid.ActorSystem
	configurator []Configurator
	server       vivid.ActorRef
}

func (g *GNETEventHandler) OnBoot(eng gnet.Engine) (action gnet.Action) {
	// server actor
	g.server = g.system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case gnet.Conn:
				ref := ctx.ActorOfF(func() vivid.Actor {
					return NewStream(newGNETConnWrapper(g.system, m), g.configurator...)
				}, func(descriptor *vivid.ActorDescriptor) {
					descriptor.WithSupervisionStrategyProvider(supervision.FunctionalStrategyProvider(func() supervision.Strategy {
						return supervision.StopStrategy()
					}))
				})
				ctx.Reply(ref)
			}
		})
	})
	return
}

func (g *GNETEventHandler) OnShutdown(eng gnet.Engine) {
	g.system.Terminate(g.server, true)
	return
}

func (g *GNETEventHandler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	message, err := g.system.FutureAsk(g.server, c).Result()
	if err != nil {
		return out, gnet.Close
	}

	reader, ok := message.(vivid.ActorRef)
	if !ok {
		return out, gnet.Close
	}

	c.SetContext(reader)
	return
}

func (g *GNETEventHandler) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	reader := c.Context().(vivid.ActorRef)
	if err != nil {
		g.system.Tell(reader, err)
		return
	}
	g.system.Terminate(reader, true)
	return
}

func (g *GNETEventHandler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	reader := c.Context().(vivid.ActorRef)
	data, err := c.Next(-1)
	if err != nil {
		g.system.Tell(reader, err)
		return
	}

	g.system.Tell(reader, NewPacketD(data))
	return
}

func (g *GNETEventHandler) OnTick() (delay time.Duration, action gnet.Action) {
	return
}
