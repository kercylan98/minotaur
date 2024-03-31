package server_test

import (
	"github.com/kercylan98/minotaur/server/internal/v2"
	"github.com/kercylan98/minotaur/server/internal/v2/network"
	"testing"
)

func TestNewServer(t *testing.T) {
	srv := server.NewServer(network.WebSocket(":9999"))
	srv.RegisterConnectionReceivePacketEvent(func(srv server.Server, conn server.Conn, packet server.Packet) {
		if err := conn.WritePacket(packet); err != nil {
			panic(err)
		}
	})
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
