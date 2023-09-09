package server_test

import (
	"github.com/kercylan98/minotaur/server"
	"time"
)

func ExampleNew() {
	srv := server.New(server.NetworkWebsocket, server.WithLimitLife(time.Millisecond))
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}

	// Output:
}

func ExampleServer_Run() {
	srv := server.New(server.NetworkWebsocket, server.WithLimitLife(time.Millisecond))
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}

	// Output:
}
