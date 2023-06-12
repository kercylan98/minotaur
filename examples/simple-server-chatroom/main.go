// 该案例实现了一个简单的聊天室功能
package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/synchronization"
)

func main() {
	connections := synchronization.NewMap[string, *server.Conn]()

	srv := server.New(server.NetworkWebsocket, server.WithWebsocketWriteMessageType(server.WebsocketMessageTypeText))
	srv.RegConnectionOpenedEvent(func(srv *server.Server, conn *server.Conn) {
		for _, c := range connections.Map() {
			c.Write([]byte(fmt.Sprintf("%s 加入了聊天室", conn.GetID())))
		}
		connections.Set(conn.GetID(), conn)
		conn.Write([]byte("欢迎加入"))
	})
	srv.RegConnectionClosedEvent(func(srv *server.Server, conn *server.Conn) {
		if connections.DeleteExist(conn.GetID()) {
			for id, c := range connections.Map() {
				c.Write([]byte(fmt.Sprintf("%s 退出了聊天室", id)))
			}
		}
	})
	srv.RegConnectionReceiveWebsocketPacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte, messageType int) {
		for _, c := range connections.Map() {
			c.Write([]byte(fmt.Sprintf("%s: %s", conn.GetID(), string(packet))))
		}
	})
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
