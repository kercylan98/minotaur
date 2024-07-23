package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
)

func main() {
	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnLaunch:
				fmt.Println("Hello, World!")
			}
		})
	})
}
