package vivid_test

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	"testing"
	"time"
)

func TestActorOptions_WithDeadline(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(log.GetDefault)
	})
	system.Add(1)

	system.ActorOf(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case vivid.OnTerminated:
				system.Done()
			}
		})
	}, func(options *vivid.ActorOptions) {
		options.WithDeadline(time.Millisecond * 100)
	})

	system.Wait()
	system.Shutdown()
}

func TestActorOptions_WithMessageDeadline(t *testing.T) {
	system := vivid.NewActorSystem(func(options *vivid.ActorSystemOptions) {
		options.WithLoggerProvider(log.GetDefault)
	})
	system.Add(1)

	system.ActorOf(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case vivid.OnTerminated:
				system.Done()
			}
		})
	}, func(options *vivid.ActorOptions) {
		options.WithMessageDeadline(time.Millisecond * 2000)
	})

	system.Wait()
	system.Shutdown()
}
