package main

import "github.com/kercylan98/minotaur/engine/vivid"

func main() {
	vivid.NewActorSystem().Publish("chat", "hello")
}
