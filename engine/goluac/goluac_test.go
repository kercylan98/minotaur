package goluac_test

import (
	_ "embed"
	"fmt"
	"github.com/kercylan98/minotaur/engine/goluac"
	"github.com/kercylan98/minotaur/engine/goluac/internal/libs"
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

	system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *libs.LuaMessage:
				fmt.Println("go vivid actor receive message: " + string(m.Data))
				ctx.Reply(m)
				fmt.Println("go vivid actor reply message: " + string(m.Data))
			}
		})
	})

	system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *vivid.OnLaunch:
				goluac.Bind(ctx, goluacTestLua)
			}
		})
	})

	time.Sleep(time.Second * 10)
}
