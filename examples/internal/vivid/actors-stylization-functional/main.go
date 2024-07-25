package main

import "github.com/kercylan98/minotaur/engine/vivid"

func main() {
	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	system.ActorOf(
		vivid.FunctionalActorProvider(func() vivid.Actor {
			return nil
		}), vivid.FunctionalActorDescriptorConfigurator(func(descriptor *vivid.ActorDescriptor) {
			descriptor.WithName("hi")
		}),
	)

	system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			// none
		})
	}))
}
