package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/vivid"
)

func main() {
	system := vivid.NewActorSystem()
	system.ActorOf(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case vivid.OnLaunch:
				fmt.Println("Actor launched")
			}
		})
	})
}
