// 该案例中实现了一个简单的回响服务器
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
