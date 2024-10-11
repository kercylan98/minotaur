package vivid_test

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"sync"
	"testing"
)

func TestSubscription(t *testing.T) {
	wg := new(sync.WaitGroup)
	start := new(sync.WaitGroup)
	start.Add(3)

	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	var refs []vivid.ActorRef
	for i := 0; i < 3; i++ {
		refs = append(refs, system.ActorOfF(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch m := ctx.Message().(type) {
				case *vivid.OnLaunch:
					ctx.Subscribe("chat")
					wg.Add(1)
					start.Done()
				case string:
					if m == "start" {
						ctx.Publish("chat", "hello")
						t.Log(ctx.Ref().URL().String(), "receive", m, "from", ctx.Sender().URL().String())
						wg.Done()
						return
					}
				}
			})
		}))
	}

	start.Wait()
	for _, ref := range refs {
		system.Tell(ref, "start")
	}

	wg.Wait()
}

func TestSharedSubscription(t *testing.T) {
	start := new(sync.WaitGroup)
	start.Add(2)
	wg := new(sync.WaitGroup)
	wg.Add(4)
	system1 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared("127.0.0.1:8888")
	}))
	system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared("127.0.0.1:9999")
	}))

	// 连接中便能散播
	system2.Tell(system1.Context().Ref().Clone(), system2.Context().Ref())

	system1.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnLaunch:
				ctx.Subscribe("chat")
				start.Done()
				start.Wait()
				ctx.Publish("chat", ctx.Ref())
			case vivid.ActorRef:
				t.Log(ctx.Ref().URL().String(), "receive", "message", "from", ctx.Sender().URL().String())
				wg.Done()
			}
		})
	})

	system2.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnLaunch:
				ctx.Subscribe("chat")
				start.Done()
				start.Wait()
				ctx.Publish("chat", ctx.Ref())
			case vivid.ActorRef:
				t.Log(ctx.Ref().URL().String(), "receive", "message", "from", ctx.Sender().URL().String())
				wg.Done()
			}
		})
	})

	wg.Wait()
}
