package goluac_test

import (
	_ "embed"
	"github.com/kercylan98/minotaur/engine/vivid"
	goluac2 "github.com/kercylan98/minotaur/engine/vivid/goluac"
	"testing"
	"time"
)

//go:embed goluac_test.lua
var goluacTestLua string

func TestGoluac(t *testing.T) {
	system := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithComponents(goluac2.NewComponent())
	}))

	luaActorRef := system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnLaunch:
				goluac2.AttachActorContext(ctx, goluacTestLua)
			}
		})
	})

	system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *vivid.OnLaunch:
				ctx.Ask(luaActorRef, goluac2.NewLuaMessage("test", map[string]any{
					"sender": ctx.Sender().URL().String(),
					"data":   "i'm received a message from goluac",
				}))
			case *goluac2.LuaMessage:
				switch m.Name {
				case "reply":
					var data = make(map[string]any)
					m.UnmarshalP(&data)
					t.Log(data)
				}
			}
		})
	})

	time.Sleep(time.Second * 10)
}
