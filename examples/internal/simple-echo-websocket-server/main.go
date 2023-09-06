package main

import "github.com/kercylan98/minotaur/server"

func main() {
	srv := server.New(server.NetworkWebsocket)
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
