package main

import (
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
	"os"
)

func main() {
	system := vivid.NewActorSystem()
	system.ActorOf(func() vivid.Actor {
		return transport.NewStreamClient(&transport.StreamClientWebSocketCore{
			Url: "ws://localhost:8877/ws",
		}, transport.StreamClientConfig{
			ConnectionOpenedHandler: func(ctx vivid.ActorContext) {
				fmt.Println("connection opened")
				ctx.Tell(ctx.Ref(), transport.NewPacket([]byte("hello")).SetContext(websocket.BinaryMessage))
			},
			ConnectionPacketHandler: func(ctx vivid.ActorContext, packet transport.Packet) {
				fmt.Println("packet received", string(packet.GetBytes()))
			},
			ConnectionClosedHandler: func(ctx vivid.ActorContext, err error) {
				fmt.Println("connection closed", err)
			},
		})
	})

	system.Signal(func(system *vivid.ActorSystem, signal os.Signal) {
		system.ShutdownGracefully()
	})
}
