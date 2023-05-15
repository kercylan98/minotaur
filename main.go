package main

import "github.com/kercylan98/minotaur/server"

// 无意义的测试main入口
func main() {
	srv := server.New(server.NetworkWebsocket, server.WithConnectPacketDiversion(3, 2))
	srv.RegConnectionReceiveWebsocketPacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte, messageType int) {
		conn.Write(packet, messageType)
	})
	if err := srv.Run(":8999"); err != nil {
		panic(err)
	}
}
