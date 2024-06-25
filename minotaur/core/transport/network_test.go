package transport_test

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/minotaur/core/transport"
	"github.com/kercylan98/minotaur/minotaur/core/vivid"
	"testing"
	"time"
)

func TestNewNetwork(t *testing.T) {
	vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewNetwork(":8800"))
	})

	time.Sleep(time.Second * 1111111)
}

func TestNewNetwork2(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewNetwork(":8888"))
	})

	time.Sleep(time.Second)

	server := core.NewProcessRef(core.NewRootAddress("", "", "127.0.0.1", 8800))
	system.Context().Tell(server, "Hello, World!")

	system.Shutdown()

	time.Sleep(time.Second * 1111111)
}
