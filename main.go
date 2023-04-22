package main

import (
	"minotaur/server"
)

func main() {
	ms := server.NewMultipleServer(
		func() (addr string, srv *server.Server) {
			return ":9999", server.New(server.NetworkTCP)
		},
		func() (addr string, srv *server.Server) {
			return ":19999", server.New(server.NetworkGRPC)
		},
	)
	ms.Run()
}
