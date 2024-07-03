package cluster_test

import (
	"github.com/kercylan98/minotaur/core/cluster"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"testing"
	"time"
)

func TestNewNode_1(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(log.GetDefault)
		options.WithModule(
			transport.NewNetwork(":8000"),
			cluster.NewNode("node-1", "127.0.0.1", 19000, "localhost:19000", "localhost:19001"),
		)
	})

	system.Shutdown(time.Hour)
}

func TestNewNode_2(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(log.GetDefault)
		options.WithModule(
			transport.NewNetwork(":8001"),
			cluster.NewNode("node-2", "127.0.0.1", 19001, "localhost:19000", "localhost:19001"),
		)
	})

	system.Shutdown(time.Second * 5)
}
