package transport_test

import (
	"github.com/fasthttp/websocket"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"net/url"
	"os"
	"testing"
)

func TestWebSocketClient(t *testing.T) {
	actorSystem := vivid.NewActorSystem()
	ref := actorSystem.ActorOf(func() vivid.Actor {
		return transport.NewWebSocketClient(url.URL{
			Scheme: "ws",
			Host:   "localhost:8877",
			Path:   "/ws",
		}, transport.WebSocketClientConfig{
			ConnectionOpenedHandler: func(ctx vivid.ActorContext) {
				t.Log("connection opened")
			},
			ConnectionPacketHandler: func(ctx vivid.ActorContext, packet transport.Packet) {
				t.Log("packet received", string(packet.GetBytes()))
			},
			ConnectionClosedHandler: func(ctx vivid.ActorContext, err error) {
				t.Log("connection closed", err)
			},
		})
	})

	actorSystem.Tell(ref, transport.NewPacket([]byte("hello")).SetContext(websocket.BinaryMessage))

	actorSystem.Signal(func(system *vivid.ActorSystem, signal os.Signal) {
		actorSystem.ShutdownGracefully()
	})
}
