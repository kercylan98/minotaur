package vivid_test

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"sync"
	"testing"
)

func TestUserGuardActor_Supervisor(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)

	vivid.ActorOfF(vivid.GetTestActorSystem(), func(options *vivid.ActorOptions[*vivid.ExportUserGuardActor]) {
		options.WithActorContextHook(func(ctx vivid.ActorContext) {
			ref := vivid.ActorOfF[*vivid.PanicActor](ctx, func(options *vivid.ActorOptions[*vivid.PanicActor]) {
				options.WithInit(func(actor *vivid.PanicActor) {
					actor.PanicHook = func(err error) {
						t.Log(fmt.Errorf("panic: %w", err))
					}
					actor.RestartHook = func() {
						t.Log("restart")
						wait.Done()
					}
				})
			})
			ref.Tell(errors.New("boom"))
		})
	})

	// wait restart
	wait.Wait()
	vivid.GetTestActorSystem().Shutdown()
}
