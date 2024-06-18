package minotaur_test

import (
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/cluster"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"testing"
	"time"
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

type ClusterActor struct {
}

func (c *ClusterActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case string:
		log.Info("ClusterActor", log.String("actor", ctx.Id().String()), log.String("message", m))
	}
}

func TestNewApplication(t *testing.T) {
	//minotaur.EnableHttpPProf(":9989", "/debug/pprof", func(err error) {
	//	panic(err)
	//})
	minotaur.NewApplication(
		minotaur.WithNetwork(network.WebSocket(":10000")),
		minotaur.WithActorSystemOptions(vivid.NewActorSystemOptions().
			WithCluster(
				cluster.WithClusterName("test"),
				cluster.WithBindAddr("0.0.0.0"),
				cluster.WithAdvertiseAddr("192.168.2.112"),
				cluster.WithBindPort(1000),
				cluster.WithAdvertisePort(1000),
			),
		),
	).Launch(func(app *minotaur.Application, ctx vivid.MessageContext) {

		switch m := ctx.GetMessage().(type) {
		case vivid.OnBoot:
			app.GetServer().SubscribeConnOpenedEvent(app)
			ctx.ActorOf(vivid.OfO(func(options *vivid.ActorOptions[*ClusterActor]) {
				options.WithName("cluster")
			}))
		case transport.ServerConnectionOpenedEvent:
			conn := m.Conn
			conn.SetPacketHandler(func(ctx vivid.MessageContext, conn transport.Conn, packet transport.ConnectionReactPacketMessage) {
				conn.Write(packet)
			})
			conn.LoadMod(vivid.ModOf[AccountModExposer](&AccountMod{}))
			conn.LoadMod(vivid.ModOf[BagModExposer](&BagMod{}))

			conn.ApplyMod()
		}

	})
}

func TestNewApplication2(t *testing.T) {
	minotaur.NewApplication(
		minotaur.WithNetwork(network.WebSocket(":10001")),
		minotaur.WithActorSystemOptions(vivid.NewActorSystemOptions().
			WithCluster(
				cluster.WithClusterName("test"),
				cluster.WithBindAddr("0.0.0.0"),
				cluster.WithAdvertiseAddr("192.168.2.112"),
				cluster.WithBindPort(1001),
				cluster.WithAdvertisePort(1001),
				cluster.WithDefaultJoinAddresses("192.168.2.112:1000"),
			),
		),
	).Launch(func(app *minotaur.Application, ctx vivid.MessageContext) {

		switch m := ctx.GetMessage().(type) {
		case vivid.OnBoot:
			app.GetServer().SubscribeConnOpenedEvent(app)
			ctx.ActorOf(vivid.OfO(func(options *vivid.ActorOptions[*ClusterActor]) {
				options.WithName("cluster")
			}))
			go func() {
				for {
					time.Sleep(time.Millisecond * 100)
					id := vivid.NewActorId("test", "192.168.2.112", 1000, "minotaur", "/user/app/cluster")
					ref := app.ActorSystem().GetActorRef(id)
					ref.Tell("hello")
				}
			}()
		case transport.ServerConnectionOpenedEvent:
			conn := m.Conn
			conn.SetPacketHandler(func(ctx vivid.MessageContext, conn transport.Conn, packet transport.ConnectionReactPacketMessage) {
				conn.Write(packet)
			})
			conn.LoadMod(vivid.ModOf[AccountModExposer](&AccountMod{}))
			conn.LoadMod(vivid.ModOf[BagModExposer](&BagMod{}))

			conn.ApplyMod()
		}

	})
}
