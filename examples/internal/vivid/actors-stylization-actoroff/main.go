package main

import "github.com/kercylan98/minotaur/engine/vivid"

func main() {
	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	system.ActorOfF(
		func() vivid.Actor {
			return nil
		}, func(descriptor *vivid.ActorDescriptor) {
			descriptor.WithName("hi")
		},
	)
}
