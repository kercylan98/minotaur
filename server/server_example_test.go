package server_test

import (
	"github.com/kercylan98/minotaur/server"
	"time"
)

func ExampleNew() {
	srv := server.New(server.NetworkWebsocket,
		server.WithDeadlockDetect(time.Second*5),
	)

	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet server.Packet) {
		conn.Write(packet)
	})

	go func() { time.Sleep(1 * time.Second); srv.Shutdown() }()
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}

	// Output:
}
