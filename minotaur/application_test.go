package minotaur_test

import (
	"github.com/gobwas/ws"
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"testing"
)

func TestNewApplication(t *testing.T) {
	minotaur.NewApplication(
		minotaur.WithNetwork(network.WebSocket(":9988")),
	).Launch(func(app *minotaur.Application, ctx vivid.MessageContext) {

		switch m := ctx.GetMessage().(type) {
		case vivid.OnBoot:
			app.GetServer().Api().SubscribeConnOpenedEvent(app)
		case transport.ServerConnectionOpenedEvent:
			m.Conn.Api().Write(transport.NewPacket([]byte("Hello, World!")).SetContext(ws.OpText))
			m.Conn.Api().BecomeReactPacketMessage(func(context vivid.MessageContext, message transport.ConnectionReactPacketMessage) {
				m.Conn.Api().Write(message.Packet) // Echo
			})
		}

	})
}
