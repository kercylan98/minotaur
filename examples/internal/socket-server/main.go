package main

import (
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/core/vivid"
)

type SocketService struct {
}

func (s *SocketService) OnInit(kit *transport.GNETKit) {
	kit.ConnectionPacketHook(s.onConnectionPacket)
}

func (s *SocketService) onConnectionPacket(kit *transport.GNETKit, conn *transport.Conn, packet transport.Packet) error {
	conn.WritePacket(packet) // echo
	return nil
}

func main() {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithModule(
			transport.NewTCP(":8000").BindService(new(SocketService)),
			transport.NewTCP4(":8001").BindService(new(SocketService)),
			transport.NewTCP6(":8002").BindService(new(SocketService)),
			transport.NewUDP(":8003").BindService(new(SocketService)),
			transport.NewUDP4(":8004").BindService(new(SocketService)),
			transport.NewUDP6(":8005").BindService(new(SocketService)),
			transport.NewWebSocket(":8006", "/ws").BindService(new(SocketService)),
			transport.NewUnix("./unix.sock").BindService(new(SocketService)),
		)
	})

	system.ShutdownGracefully()
}
