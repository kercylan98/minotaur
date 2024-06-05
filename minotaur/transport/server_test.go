package transport_test

import (
	"github.com/kercylan98/minotaur/minotaur/pulse"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

type Conn struct {
	*transport.ConnActor
}

func (c *Conn) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		c.ConnActor.OnReceive(ctx)
	case transport.ConnReceivePacketMessage:
		c.Write(m.Packet)
	}
}

type AccountManager struct {
	eventBus *pulse.Pulse
}

func (e *AccountManager) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		e.eventBus.Subscribe(pulse.SubscribeId(ctx.GetRef().Id()), ctx.GetRef(), transport.ServerConnOpenedEvent{})
	case transport.ServerConnOpenedEvent:
		vivid.ActorOf[*Conn](ctx, vivid.NewActorOptions[*Conn]().WithInit(func(conn *Conn) {
			conn.ConnActor = m.ConnActor
		}))

	}
}

func TestNewServer(t *testing.T) {
	system := vivid.NewActorSystem("test-srv")
	eventBus := pulse.NewPulse(&system)
	srv := transport.NewServer(&system)
	srv.Launch(&eventBus, network.WebSocket(":8899"))

	vivid.ActorOf[*AccountManager](&system, vivid.NewActorOptions[*AccountManager]().WithInit(func(manager *AccountManager) {
		manager.eventBus = &eventBus
	}))

	var s = make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-s

	system.Shutdown()
	t.Log("Server shutdown")
}
