package main

import "github.com/kercylan98/minotaur/server"

func main() {
	srv := server.New(server.NetworkWebsocket, server.WithTicker(50, false))
	if err := srv.Run(":9999"); err != nil {
		panic(err)
	}
}
