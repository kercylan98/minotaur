package minotaur_test

import (
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"testing"
)

type AccountManager struct {
	*minotaur.Application
}

func (e *AccountManager) OnReceive(ctx vivid.MessageContext) {
	switch ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		ctx.BindBehavior(vivid.BehaviorOf[transport.ServerConnOpenedEvent](e.onConnOpened))
		e.EventBus().Subscribe("conn-opened", ctx, transport.ServerConnOpenedEvent{})
	}
}

func (e *AccountManager) onConnOpened(ctx vivid.MessageContext, message transport.ServerConnOpenedEvent) {
	vivid.ActorOf[*Account](e.ActorSystem(), vivid.NewActorOptions[*Account]().WithInit(func(account *Account) {
		account.ConnActor = message.ConnActor
	}))
}

type Account struct {
	*transport.ConnActor
}

func (c *Account) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnPreStart:
		c.ConnActor.OnReceive(ctx)
	case transport.ConnReceivePacketMessage:
		c.Write(m.Packet)
	}
}

func TestNewApplication(t *testing.T) {
	app := minotaur.NewApplication("test", minotaur.WithNetwork(network.WebSocket(":9988")))
	vivid.ActorOf[*AccountManager](app.ActorSystem(), vivid.NewActorOptions[*AccountManager]().WithInit(func(manager *AccountManager) {
		manager.Application = app
	}))
	app.Launch()
}
