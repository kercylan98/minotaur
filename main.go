package main

import "github.com/kercylan98/minotaur/server"

// 无意义的测试main入口
func main() {
	srv := server.New(server.NetworkWebsocket, server.WithConnectPacketDiversion(3, 2))
	if err := srv.Run(":8999"); err != nil {
		panic(err)
	}
}
