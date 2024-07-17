package vivid_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/random"
	"sync"
	"testing"
	"time"
)

func TestActorContext_DeadLetter(t *testing.T) {
	system := vivid.NewTestActorSystem().Add(1)
	system.ActorOf(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case vivid.OnLaunch:
				if ctx.DeadLetter() == nil {
					t.Error("not found dead letter")
				}
				ctx.System().Done()
			}
		})
	})
	system.Wait().Shutdown()
}

func TestActorContext_PersistSnapshot(t *testing.T) {
	system := vivid.NewActorSystem().Add(1)

	var once sync.Once
	var initState string

	system.ActorOf(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case vivid.OnLaunch:
				randomStatus := random.HostName()
				if initState == "" {
					initState = randomStatus
				}
				ctx.PersistSnapshot(initState)
				once.Do(func() {
					panic("restart")
				})
			case string: // recover snapshot
				if initState != m {
					t.Errorf("status not match, init: %s, persist:%s", initState, m)
				}
				ctx.System().Done()
			}
		})
	})

	system.Wait().Shutdown()
}

func TestActorContext_StatusChanged(t *testing.T) {
	system := vivid.NewTestActorSystem().Add(1)

	var once sync.Once
	var stateLimit = 2000

	ref := system.ActorOf(func() vivid.Actor {
		state := 0
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case vivid.OnLaunch:
				t.Log("curr statue: ", state)
			case bool: // event
				state++
				ctx.StatusChanged(m)
				if state == stateLimit {
					once.Do(func() {
						t.Log("restart before statue: ", state)
						panic("restart")
					})
					ctx.System().Done()
				}
			}
		})
	})

	for i := 0; i < stateLimit; i++ {
		system.Context().Tell(ref, true)
	}

	system.Wait().Shutdown()
}

func TestActorContext_Sender(t *testing.T) {
	t.Run("Tell", func(t *testing.T) {
		system := vivid.NewTestActorSystem().Add(1)
		system.Context().Tell(system.ActorOf(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch ctx.Message().(type) {
				case *testing.T:
					if ctx.Sender().Address() != ctx.System().DeadLetter().Ref().Address() {
						t.Error("tell function should not own sender")
					}
					ctx.System().Done()
				}
			})
		}), t)
		system.Wait().Shutdown()
	})

	t.Run("Ask", func(t *testing.T) {
		system := vivid.NewTestActorSystem().Add(1)
		system.Context().Ask(system.ActorOf(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch ctx.Message().(type) {
				case *testing.T:
					if ctx.Sender().Address() == ctx.System().DeadLetter().Ref().Address() {
						t.Error("ask function must have sender")
					}
					ctx.System().Done()
				}
			})
		}), t)
		system.Wait().Shutdown()
	})

	t.Run("FutureAsk", func(t *testing.T) {
		system := vivid.NewTestActorSystem().Add(1)
		system.Context().FutureAsk(system.ActorOf(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch m := ctx.Message().(type) {
				case *testing.T:
					if ctx.Sender().Address() == ctx.System().DeadLetter().Ref().Address() {
						t.Error("ask function must have sender")
					}
					ctx.System().Done()
					ctx.Reply(m)
				}
			})
		}), t).AssertWait()
		system.Wait().Shutdown()
	})
}

func TestActorContext_ActorOf(t *testing.T) {
	system := vivid.NewTestActorSystem().Add(1)
	system.ActorOf(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case vivid.OnLaunch:
				ctx.System().Done()
			}
		})
	})

	system.Wait().Shutdown()
}

func TestActorContext_KindOf(t *testing.T) {
	system := vivid.NewTestActorSystem().Add(1)
	system.RegKind("test", func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case vivid.OnLaunch:
				ctx.System().Done()
			}
		})
	})

	system.KindOf("test")
	system.Wait().Shutdown()
}

type StateActor struct {
	counter int
}

func (s *StateActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case int:
		ctx.StatusChanged(m)
		s.counter++
	case vivid.OnPersistenceSnapshot:
		// 持久化快照
		ctx.PersistSnapshot(s)
	case *StateActor:
		// 恢复状态
		s.counter = m.counter
	case string:
		fmt.Println(s.counter)
	}
}

func TestActorContext_Persist(t *testing.T) {
	system := vivid.NewTestActorSystem()

	// 创建 Actor
	ref := system.ActorOf(func() vivid.Actor {
		return &StateActor{}
	}, func(options *vivid.ActorOptions) {
		options.WithPersistenceEventLimit(10).WithPersistence(vivid.NewMemoryStorage(), "state")
	})

	// 发送消息
	for i := 0; i < 11; i++ {
		system.Context().Tell(ref, i)
	}

	time.Sleep(time.Second)

	// 创建 Actor
	ref2 := system.ActorOf(func() vivid.Actor {
		return &StateActor{}
	}, func(options *vivid.ActorOptions) {
		options.WithPersistenceEventLimit(10).WithPersistence(vivid.NewMemoryStorage(), "state")
	})
	system.Context().Tell(ref2, "print")

	time.Sleep(time.Second)
}
