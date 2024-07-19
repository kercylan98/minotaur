package vivid_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid"
	"sync"
	"testing"
)

func TestNewActorSystem(t *testing.T) {
	t.Run("functional_config", func(t *testing.T) {
		vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
			config.WithName("test-sys")
		})).Shutdown(true)
	})
}

func TestActorSystem_ActorOf(t *testing.T) {
	t.Run("functional", func(t *testing.T) {
		system := vivid.NewActorSystem()
		system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {})
		}))
		system.Shutdown(true)
	})
}

func TestActorSystem_Terminate(t *testing.T) {
	t.Run("gracefully", func(t *testing.T) {
		var wait sync.WaitGroup
		wait.Add(1)
		system := vivid.NewActorSystem()
		ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch ctx.Message().(type) {
				case vivid.OnLaunch:
					t.Log("launch")
				case vivid.OnTerminate:
					t.Log("terminate")
				case vivid.OnTerminated:
					t.Log("terminated")
					wait.Done()
				}
			})
		}))

		system.Terminate(ref, true)
		wait.Wait()
	})

	t.Run("terminate", func(t *testing.T) {
		var wait sync.WaitGroup
		wait.Add(1)
		system := vivid.NewActorSystem()
		ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch ctx.Message().(type) {
				case vivid.OnLaunch:
					t.Log("launch")
				case vivid.OnTerminate:
					t.Log("terminate")
				case vivid.OnTerminated:
					t.Log("terminated")
					wait.Done()
				}
			})
		}))

		system.Terminate(ref, false)
		wait.Wait()
	})
}

func TestActorSystem_Shutdown(t *testing.T) {
	system := vivid.NewActorSystem()
	system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {})
	}))

	system.Shutdown(true)
}
