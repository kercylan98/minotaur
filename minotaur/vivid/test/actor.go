package test

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"testing"
)

var t *testing.T

func initTest(test *testing.T) {
	t = test
}

type Actor struct {
}

func (a *Actor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnBoot:
		//t.Log("start", ctx.Id().Name())
	case vivid.OnRestart:
		//t.Log("restart", ctx.Id().Name())
	case vivid.OnTerminate:
		//t.Log("terminate", ctx.Id().Name())
	case string:
		switch m {
		case "child":
			ctx.ActorOf(vivid.OfO[*Actor](func(actorOptions *vivid.ActorOptions[*Actor]) {
				actorOptions.WithName("child")
			}))
		}
	case error:
		panic("unknown message")
	}
}
