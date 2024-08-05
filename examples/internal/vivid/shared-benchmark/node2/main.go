package main

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"time"
)

//goland:noinspection t
func main() {
	system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared("127.0.0.1:8081")
	}))

	defer system2.Shutdown(true)

	system2.ActorOfF(func() vivid.Actor {
		var ref vivid.ActorRef
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case vivid.ActorRef:
				ref = m
				ctx.Reply(ctx.Ref())
			case *Ping:
				ctx.Tell(ref, &Pong{})
			}
		})
	})

	time.Sleep(time.Hour)
}
