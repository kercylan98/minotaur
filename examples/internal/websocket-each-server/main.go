package main

import (
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/transport"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

func main() {
	minotaur.NewApplication(
		minotaur.WithNetwork(network.WebSocket(":8080")),
	).Launch(func(app *minotaur.Application, ctx vivid.MessageContext) {
		switch m := ctx.GetMessage().(type) {
		case vivid.OnBoot:
			app.GetServer().SubscribeConnOpenedEvent(app)
		case transport.ServerConnectionOpenedEvent:
			m.Conn.SetPacketHandler(func(ctx vivid.MessageContext, conn transport.Conn, packet transport.ConnectionReactPacketMessage) {
				conn.Write(packet)
			})
		}
	})
}
