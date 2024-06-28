package transport_test

import (
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"testing"
	"time"
)

func TestNewWebSocket(t *testing.T) {
	vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewWebSocket(":8080", "/ws"))
	})

	time.Sleep(time.Second * 1000)
}
