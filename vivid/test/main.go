package main

import (
	"fmt"
	vivid "github.com/kercylan98/minotaur/vivid"
)

type TestFreeActor struct {
	name string
}

func main() {
	system := vivid.NewActorSystem("test-system")

	system.Shutdown()
	fmt.Println("System shutdown")
}
