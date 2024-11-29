package goluac

import (
	"github.com/kercylan98/minotaur/engine/goluac/internal/libs"
	"github.com/kercylan98/minotaur/engine/vivid"
	lua "github.com/yuin/gopher-lua"
)

func newGoluac(ctx vivid.ActorContext, code string) *goluac {
	l := lua.NewState()
	c := &goluac{lib: libs.NewLibrary(l, ctx)}
	if err := l.DoString(code); err != nil {
		panic(err)
	}
	return c
}

type goluac struct {
	lib *libs.Library
}
