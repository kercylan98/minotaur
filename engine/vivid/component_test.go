package vivid_test

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"sync"
	"testing"
)

func TestComponent(t *testing.T) {
	wait := new(sync.WaitGroup)
	wait.Add(1)

	system := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithComponents(vivid.FunctionalActorContextCaptureComponent(func(actorSystem *vivid.ActorSystem, ctx vivid.ActorContext) {
			ctx.Tell(ctx.Ref(), 1)
		}))
	}))

	system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case int:
				t.Log("Receive message")
				wait.Done()
			}
		})
	})

	wait.Wait()
	system.Shutdown(true)
}
