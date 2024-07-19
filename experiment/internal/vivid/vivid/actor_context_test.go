package vivid_test

import (
	"errors"
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid"
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid/supervision"
	"sync"
	"testing"
	"time"
)

func TestActorContext_Sender(t *testing.T) {
	system := vivid.NewActorSystem()
	ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case string:
				t.Log("received", m)
				switch m {
				case "tell":
					if ctx.Sender() != nil {
						t.Error("tell sender should be nil")
					}
				default:
					if ctx.Sender() == nil {
						t.Error("ask sender should not be nil")
					}
				}
			}
		})
	}))

	system.Tell(ref, "tell")
	system.Ask(ref, "ask")
	system.FutureAsk(ref, "future-ask")
	system.Broadcast("broadcast")
	system.Shutdown(true)
}

func TestActorContext_Message(t *testing.T) {
	var ok bool
	system := vivid.NewActorSystem()
	ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case string:
				t.Log("received", ctx.Message())
				ok = true
			}
		})
	}))

	system.Tell(ref, "tell")
	system.Shutdown(true)
	if !ok {
		t.Error("tell failed")
	}
}

func TestActorContext_Tell(t *testing.T) {
	wait := new(sync.WaitGroup)
	wait.Add(1)
	system := vivid.NewActorSystem()
	ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case string:
				wait.Done()
			}
		})
	}))

	system.Tell(ref, "hello")
	wait.Wait()
	system.Shutdown(true)
}

func TestActorContext_Ask(t *testing.T) {
	wait := new(sync.WaitGroup)
	wait.Add(1)
	system := vivid.NewActorSystem()
	refA := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case string:
				ctx.Reply("i'm fine")
			}
		})
	}))

	system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case vivid.OnLaunch:
				ctx.Ask(refA, "hello")
			case string:
				wait.Done()
			}
		})
	}))

	wait.Wait()
	system.Shutdown(true)
}

func TestActorContext_FutureAsk(t *testing.T) {
	system := vivid.NewActorSystem()
	ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case string:
				ctx.Reply("reply")
			}
		})
	}))

	f := system.FutureAsk(ref, "ask")
	reply, err := f.Result()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("receive:", reply)
}

func TestActorContext_Broadcast(t *testing.T) {
	wait := new(sync.WaitGroup)
	system := vivid.NewActorSystem()
	for i := 0; i < 10; i++ {
		wait.Add(1)
		system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch ctx.Message().(type) {
				case string:
					wait.Done()
				}
			})
		}))
	}
	system.Broadcast("hello")
	wait.Wait()
	system.Shutdown(true)
}

func TestActorContext_Restart(t *testing.T) {
	wait := new(sync.WaitGroup)
	wait.Add(1)
	system := vivid.NewActorSystem()
	ref := system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case vivid.OnLaunch:
				t.Log("launch")
			case vivid.OnTerminate:
				t.Log("terminate")
			case vivid.OnTerminated:
				t.Log("terminated")
			case vivid.OnRestarting:
				t.Log("restarting")
			case vivid.OnRestarted:
				t.Log("restarted")
				wait.Done()
			case error:
				panic(m)
			}
		})
	}))

	system.Tell(ref, errors.New("restart"))
	wait.Wait()
	system.Shutdown(true)
}

func TestActorContext_RestartN(t *testing.T) {
	wait := new(sync.WaitGroup)
	system := vivid.NewActorSystem()
	restartCount := 3
	wait.Add(restartCount)
	system.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case vivid.OnLaunch:
				panic(m)
			case vivid.OnTerminate:
				t.Log("terminate")
			case vivid.OnTerminated:
				t.Log("terminated")
			case vivid.OnRestarting:
				t.Log("restarting")
			case vivid.OnRestarted:
				t.Log("restarted")
				wait.Done()
			}
		})
	}), vivid.FunctionalActorDescriptorConfigurator(func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithSupervisionStrategy(supervision.OneForOne(restartCount, time.Millisecond*100, time.Millisecond*100, supervision.FunctionalDecide(func(record *supervision.AccidentRecord) supervision.Directive {
			return supervision.DirectiveRestart
		})))
	}))
	wait.Wait()
	system.Shutdown(true)
}
