package main

import (
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
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

func main() {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewFiber(":8877").BindService(new(WebSocketService)))
	})

	system.ShutdownGracefully()
}
