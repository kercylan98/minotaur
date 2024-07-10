package main

import (
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
)

type WebSocketService struct {
}

func (s *WebSocketService) OnInit(kit *transport.GNETKit) {
	kit.ConnectionPacketHook(s.onConnectionPacket)
}

func (s *WebSocketService) onConnectionPacket(kit *transport.GNETKit, conn *transport.Conn, packet transport.Packet) error {
	conn.WritePacket(packet) // echo
	return nil
}

func main() {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(transport.NewWebSocket(":8877", "/ws").BindService(new(WebSocketService)))
	})

	system.ShutdownGracefully()
}
