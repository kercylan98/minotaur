package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
)

func main() {
	type ChatMessage string

	vivid.NewActorSystem().ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *vivid.OnLaunch:
				ctx.Subscribe("chat")
				ctx.Publish("chat", ChatMessage("hi"))
			case ChatMessage:
				fmt.Println(m)
			}
		})
	})
}
