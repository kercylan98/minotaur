package minotaur_test

import (
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/pulse"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"testing"
)

type AccountManager struct {
	app *minotaur.Application
}

func (e *AccountManager) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		e.app.EventBus().Subscribe(pulse.SubscribeId(ctx.GetReceiver().Id()), ctx.GetReceiver(), transport.ConnOpenedEvent{})
	case transport.ConnOpenedEvent:
		m.Conn.Write(transport.NewPacket([]byte("hello")))
	}

}

func TestNewApplication(t *testing.T) {
	app := minotaur.NewApplication("test", minotaur.WithNetwork(network.WebSocket(":9988")))
	vivid.ActorOf[*AccountManager](app.ActorSystem(), vivid.NewActorOptions[*AccountManager]().WithConstruct(func() *AccountManager {
		return &AccountManager{app: app}
	}()))
	app.Launch()
}
