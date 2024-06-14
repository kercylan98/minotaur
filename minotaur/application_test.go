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
	bag BagModExposer
}

type AccountModExposer interface {
	vivid.Mod
}

type BagMod struct {
}

type BagModExposer interface {
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
		a.bag = vivid.InvokeMod[BagModExposer](ctx)
	case vivid.ModLifeCycleOnStart:
		a.bag.PingBag()
	default:
	}
}

func TestNewApplication(t *testing.T) {
	minotaur.EnableHttpPProf(":9989", "/debug/pprof", func(err error) {
		panic(err)
	})
	minotaur.NewApplication(
		minotaur.WithNetwork(network.WebSocket(":9988")),
	).Launch(func(app *minotaur.Application, ctx vivid.MessageContext) {

		switch m := ctx.GetMessage().(type) {
		case vivid.OnBoot:
			app.GetServer().Api().SubscribeConnOpenedEvent(app)
		case transport.ServerConnectionOpenedEvent:
			conn := m.Conn
			conn.Api().SetPacketHandler(func(ctx vivid.MessageContext, conn transport.Conn, packet transport.ConnectionReactPacketMessage) {
				conn.Api().Write(packet)
			})
			conn.Api().LoadMod(vivid.ModOf[AccountModExposer](&AccountMod{}))
			conn.Api().LoadMod(vivid.ModOf[BagModExposer](&BagMod{}))

			conn.Api().ApplyMod()
		}

	})
}
