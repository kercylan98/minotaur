package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/vivid"
)

type MyActor struct {
}

func (m *MyActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case string:
		fmt.Println(m)
	}
}

func main() {
	system := vivid.NewActorSystem()
	ref := system.ActorOf(func() vivid.Actor {
		return &MyActor{}
	})

	system.Context().Tell(ref, "Hi, minotaur")
	system.ShutdownGracefully()
}
