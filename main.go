package main

import "minotaur/server"

func main() {
	s := server.New(server.NetworkKcp)
	s.Run(":9999")
}
