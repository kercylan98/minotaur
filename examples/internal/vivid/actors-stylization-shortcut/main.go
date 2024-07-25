package main

import "github.com/kercylan98/minotaur/engine/vivid"

func main() {
	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	system.ActorOf(vivid.NewShortcutActorProvider("calc", func() vivid.Actor {
		return nil
	}))
}
