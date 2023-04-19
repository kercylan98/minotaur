package main

import "minotaur/server"

func main() {
	s := server.New(server.NetworkWebsocket)

	s.Run(":9999")
}
