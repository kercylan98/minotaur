package minotaur_test

import (
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"testing"
)

type AccountMod struct {
	bag BagModInterface
}

type AccountModInterface interface {
	vivid.Mod
}

type BagMod struct {
}

type BagModInterface interface {
	vivid.Mod
	PingBag()
}

func (b BagMod) OnLifeCycle(ctx vivid.ActorContext, lifeCycle vivid.ModLifeCycle) {
	switch lifeCycle {
	case vivid.ModLifeCycleOnInit:
		log.Info("BagMod", log.String("lifeCycle", "OnInit"))
	default:
	}
}

func (b BagMod) PingBag() {
	log.Info("BagMod", log.String("ping", "pong"))
}

func (a *AccountMod) OnLifeCycle(ctx vivid.ActorContext, lifeCycle vivid.ModLifeCycle) {
	switch lifeCycle {
	case vivid.ModLifeCycleOnInit:
		log.Info("AccountMod", log.String("lifeCycle", "OnInit"))
	case vivid.ModLifeCycleOnPreload:
		a.bag = vivid.InvokeMod[BagModInterface](ctx)
	case vivid.ModLifeCycleOnStart:
		a.bag.PingBag()
	default:
	}
}

func TestNewApplication(t *testing.T) {
	minotaur.NewApplication(
		minotaur.WithNetwork(network.WebSocket(":9988")),
	).Launch(func(app *minotaur.Application, ctx vivid.MessageContext) {

		switch m := ctx.GetMessage().(type) {
		case vivid.OnBoot:
			app.GetServer().Api().SubscribeConnOpenedEvent(app)
		case transport.ServerConnectionOpenedEvent:
			conn := m.Conn

			conn.Api().LoadMod(vivid.ModOf[AccountModInterface](&AccountMod{}))
			conn.Api().LoadMod(vivid.ModOf[BagModInterface](&BagMod{}))

			conn.Api().ApplyMod()
		}

	})
}
