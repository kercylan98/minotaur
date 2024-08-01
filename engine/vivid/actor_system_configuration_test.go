package vivid_test

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"sync"
	"testing"
	"time"
)

func TestActorDescriptor_WithSlowProcessingDuration(t *testing.T) {
	wait := new(sync.WaitGroup)
	wait.Add(1)

	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	watcher := system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnSlowProcess:
				wait.Done()
			}
		})
	})

	slow := system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case int:
				time.Sleep(time.Second)
			}
		})
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithSlowProcessingDuration(time.Millisecond*100, watcher)
	})

	system.Tell(slow, 1)
	wait.Wait()
}

func TestActorSystemConfiguration_WithShared(t *testing.T) {
	vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared(":8080")
	})).Shutdown(true)
}

func TestActorSystemConfiguration_WithSharedFutureAsk(t *testing.T) {
	var messageNum = 1000
	var receiverOnce, senderOnce sync.Once

	system1 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared(":8080")
	}))

	system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared(":8081")
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
		config.WithShared(":8080")
	}))

	system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared(":8081")
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

func TestActorSystemConfiguration_WithSharedActorOf(t *testing.T) {
	wait := sync.WaitGroup{}
	wait.Add(1)
	provider := vivid.NewShortcutActorProvider("test", func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnLaunch:
				t.Log("launch", ctx.Ref().URL().String(), "parent", ctx.Parent().URL().String())
			case *vivid.OnTerminated:
				t.Log("terminated", ctx.Ref().URL().String(), "killer", ctx.Sender().URL().String())
				wait.Done()
			case *prc.ProcessId:
				t.Log("hello!")

			}
		})
	})

	system1 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithName("system1")
		config.WithShared(":8080")
		config.WithCluster("127.0.0.1:18080", "127.0.0.1:18080")
		config.WithActorProvider(provider)
	}))

	system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithName("system2")
		config.WithShared(":8081")
		config.WithCluster("127.0.0.1:18081", "127.0.0.1:18080")
	}))

	ref := system2.ActorOf(provider)
	system2.Tell(ref, ref)
	system2.Terminate(ref, true)

	wait.Wait()
	system1.Shutdown(true)
	system2.Shutdown(true)
}

func TestActorDescriptor_WithExpireDuration(t *testing.T) {
	t.Run("except", func(t *testing.T) {
		wait := new(sync.WaitGroup)
		wait.Add(1)
		system := vivid.NewActorSystem()
		system.ActorOfF(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch ctx.Message().(type) {
				case *vivid.OnTerminated:
					wait.Done()
				}
			})
		}, func(descriptor *vivid.ActorDescriptor) {
			descriptor.WithExpireDuration(time.Second)
		})
		wait.Wait()
		system.Shutdown(true)
	})

	t.Run("panic restart", func(t *testing.T) {
		wait := new(sync.WaitGroup)
		wait.Add(2)
		system := vivid.NewActorSystem()
		ref := system.ActorOfF(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch ctx.Message().(type) {
				case *vivid.OnTerminate:
					t.Log("terminate")
				case *vivid.OnTerminated:
					wait.Done()
				case string:
					panic("gg")
				}
			})
		}, func(descriptor *vivid.ActorDescriptor) {
			descriptor.WithExpireDuration(time.Second)
		})
		system.Tell(ref, "gg")
		wait.Wait()
		system.Shutdown(true)
	})
}

func TestActorDescriptor_WithIdleDeadline(t *testing.T) {
	wait := new(sync.WaitGroup)
	wait.Add(1)
	system := vivid.NewActorSystem()
	system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnTerminated:
				wait.Done()
			}
		})
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithIdleDeadline(time.Millisecond * 500)
	})
	wait.Wait()
	system.Shutdown(true)
}
