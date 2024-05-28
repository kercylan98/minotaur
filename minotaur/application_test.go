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
	switch ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		ctx.BindBehavior(vivid.BehaviorOf[transport.ConnOpenedEvent](e.onConnOpened))
		e.app.EventBus().Subscribe(pulse.SubscribeId(ctx.GetReceiver().Id()), ctx.GetReceiver(), transport.ConnOpenedEvent{})
	}
}

func (e *AccountManager) onConnOpened(ctx vivid.MessageContext, message transport.ConnOpenedEvent) {
	conn := vivid.ActorOf[*Account](e.app.ActorSystem(), vivid.NewActorOptions[*Account]().WithInit(func(account *Account) {
		account.app = e.app
		account.conn = message.Conn
	}))
	e.app.EventBus().Subscribe(pulse.SubscribeId(conn.Id()), conn, transport.ConnReceiveEvent{})
}

type Account struct {
	app  *minotaur.Application
	conn transport.Conn
}

func (e *Account) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
	case transport.ConnReceiveEvent:
		e.conn.Write(m.Packet)
	}
}

func TestNewApplication(t *testing.T) {
	app := minotaur.NewApplication("test", minotaur.WithNetwork(network.WebSocket(":9988")))
	vivid.ActorOf[*AccountManager](app.ActorSystem(), vivid.NewActorOptions[*AccountManager]().WithInit(func(manager *AccountManager) {
		manager.app = app
	}))
	app.Launch()
}
