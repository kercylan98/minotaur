package goluac_test

import (
	_ "embed"
	"github.com/kercylan98/minotaur/engine/goluac"
	"github.com/kercylan98/minotaur/engine/vivid"
	"testing"
	"time"
)

//go:embed goluac_test.lua
var goluacTestLua string

func TestGoluac(t *testing.T) {
	system := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithComponents(goluac.NewComponent())
	}))

	luaActorRef := system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnLaunch:
				goluac.AttachActorContext(ctx, goluacTestLua)
			}
		})
	})

	system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *vivid.OnLaunch:
				ctx.Ask(luaActorRef, &goluac.LuaMessage{Data: []byte(`{"type": "test", "data": "hello goluac" }`)})
				ctx.Ask(luaActorRef, map[string]any{
					"type": "test",
					"data": "hello map",
				})
			case *goluac.LuaMessage:
				t.Log(string(m.Data))
			}
		})
	})

	time.Sleep(time.Second * 10)
}
