package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
)

type HelloActor struct{}

func (m *HelloActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case string:
		fmt.Println("Hello world!")
		ctx.Reply(m)
	}
}

func main() {
	system := vivid.NewActorSystem()
	ref := system.ActorOfF(func() vivid.Actor {
		return &HelloActor{}
	})

	reply := system.FutureAsk(ref, "Hey, sao ju~").AssertResult()
	fmt.Println(reply)
}
