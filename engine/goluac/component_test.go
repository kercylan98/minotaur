package goluac_test

import (
	"github.com/kercylan98/minotaur/engine/goluac"
	"github.com/kercylan98/minotaur/engine/vivid"
	"testing"
	"time"
)

func TestLuaComponent(t *testing.T) {

	system := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithComponents(goluac.New())
	}))

	system.ActorOfF(func() vivid.Actor {
		// language=lua
		luaScript := `
			print("hello, world!")
		`
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnLaunch:
				goluac.BindLuaScript(ctx, luaScript)
			}
		})
	})

	time.Sleep(time.Second)
	system.Shutdown(true)
}
