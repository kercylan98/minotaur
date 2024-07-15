package cluster_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/cluster"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"os"
	"testing"
	"time"
)

func TestCluster_Node1(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(log.GetDefault)
		options.WithModule(
			transport.NewNetwork(":8000"),
			cluster.NewNode("node-1", "127.0.0.1", 19000, "localhost:19000", "localhost:19001"),
		)
	})
	system.RegKind("test", func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case string:
				fmt.Println(m)
			}
		})
	}, func(options *vivid.ActorOptions) {
		options.WithNamePrefix("cluster")
	})

	system.Signal(func(system *vivid.ActorSystem, signal os.Signal) {
		system.ShutdownGracefully()
	})
}

func TestCluster_Node2(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(log.GetDefault)
		options.WithModule(
			transport.NewNetwork(":8001"),
			cluster.NewNode("node-2", "127.0.0.1", 19001, "localhost:19000", "localhost:19001"),
		)
	})
	system.RegKind("test", func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case string:
				fmt.Println(m)
			}
		})
	}, func(options *vivid.ActorOptions) {
		options.WithNamePrefix("cluster")
	})

	c := cluster.Invoke(system)

	for i := 0; i < 10; i++ {
		ref := c.KindOf("test")
		system.Context().Tell(ref, "hello world")
		system.Context().TerminateGracefully(ref)
	}

	system.ShutdownGracefully(time.Second * 3)
}
