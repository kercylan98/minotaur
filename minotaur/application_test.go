package minotaur_test

import (
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"testing"
)

func TestNewApplication(t *testing.T) {
	app := minotaur.NewApplication(
		minotaur.WithNetwork(network.WebSocket(":9988")),
		minotaur.WithLaunchedHook(func(app *minotaur.Application) {
			app.GetServer().
				Protocol().
				SubscribeConnOpenedEvent("conn_opened",
					func(ctx vivid.MessageContext, event transport.ServerConnOpenedEvent) {
						t.Log("conn opened")
					},
				)
		}),
	)

	app.Launch()
}
