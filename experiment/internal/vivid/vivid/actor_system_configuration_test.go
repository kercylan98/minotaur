package vivid_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/prc"
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid"
	"sync"
	"testing"
)

func TestActorSystemConfiguration_WithShared(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Error("expect panic, but not")
			}
		}()
		vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
			config.WithShared(true)
		})).Shutdown(true)
	})

	t.Run("good", func(t *testing.T) {
		vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
			config.
				WithPhysicalAddress(":8080").
				WithShared(true)
		})).Shutdown(true)
	})
}

func TestActorSystemConfiguration_WithSharedFutureAsk(t *testing.T) {
	var messageNum = 1000
	var receiverOnce, senderOnce sync.Once

	system1 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithPhysicalAddress(":8080")
		config.WithShared(true)
	}))

	system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithPhysicalAddress(":8081")
		config.WithShared(true)
	}))

	ref1 := system1.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch v := ctx.Message().(type) {
			case *prc.ProcessId:
				receiverOnce.Do(func() {
					t.Log("receiver start", ctx.Ref().URL().String(), ctx.Sender().URL().String())
				})
				ctx.Reply(v)
			}
		})
	})

	ref2 := ref1.Clone() // 同一进程内，隔离开，避免通过缓存直接调用

	message := prc.NewProcessId("none", "none")
	for i := 1; i <= messageNum; i++ {
		if err := system2.FutureAsk(ref2, message, -1).Wait(); err != nil {
			panic(err)
		}
		senderOnce.Do(func() {
			t.Log("sender start")
		})
	}

	system1.Shutdown(true)
	system2.Shutdown(true)
}

func TestActorSystemConfiguration_WithSharedTerminate(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	system1 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithPhysicalAddress(":8080")
		config.WithShared(true)
	}))

	system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithPhysicalAddress(":8081")
		config.WithShared(true)
	}))

	ref1 := system1.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnTerminate:
				t.Log("terminated", ctx.Ref().URL().String(), "killer", ctx.Sender().URL().String())
				wg.Done()
			}
		})
	})

	ref2 := ref1.Clone() // 同一进程内，隔离开，避免通过缓存直接调用
	system2.Terminate(ref2, true)
	wg.Wait()
	system1.Shutdown(true)
	system2.Shutdown(true)
}
