package minotaur_test

import (
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"testing"
)

func TestNewApplication(t *testing.T) {
	app := minotaur.NewApplication(minotaur.WithNetwork(network.WebSocket(":9988")))
	app.GetServer().Protocol().SubscribeConnOpenedEvent("conn_opened", app.ActorSystem(), func(event transport.ServerConnOpenedEvent) {
		t.Log("conn opened")
	})
	app.Launch()
}
