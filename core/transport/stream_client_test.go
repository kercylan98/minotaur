package transport_test

import (
	"github.com/fasthttp/websocket"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"os"
	"testing"
)

type WebSocketService struct {
}

func (f *WebSocketService) OnInit(kit *transport.FiberKit) {
	kit.WebSocket("/ws").ConnectionPacketHook(f.onConnectionPacket)
}

func (f *WebSocketService) onConnectionPacket(kit *transport.FiberKit, ctx *transport.FiberContext, conn *transport.Conn, packet transport.Packet) error {
	conn.WritePacket(packet) // echo
	return nil
}

func TestWebSocketServer(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewFiber(":8877").BindService(new(WebSocketService)))
	})

	system.Signal(func(system *vivid.ActorSystem, signal os.Signal) {
		system.ShutdownGracefully()
	})
}

func TestWebSocketClient(t *testing.T) {
	actorSystem := vivid.NewActorSystem()
	ref := actorSystem.ActorOf(func() vivid.Actor {
		return transport.NewStreamClient(&transport.StreamClientTCPCore{
			Url: "ws://localhost:8877/ws",
		}, transport.StreamClientConfig{
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
