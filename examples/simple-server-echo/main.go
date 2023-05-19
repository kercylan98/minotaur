package main

import (
	"github.com/kercylan98/minotaur/server"
)

func main() {
	srv := server.New(server.NetworkWebsocket)
	srv.RegConnectionReceiveWebsocketPacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte, messageType int) {
		conn.Write(packet, messageType)
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
