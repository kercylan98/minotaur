package vivid_test

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"testing"
)

func TestFutureAsk(t *testing.T) {
	system := vivid.NewActorSystem()
	ref := system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case string:
				ctx.Reply(m)
			}
		})
	})

	f := vivid.FutureAsk[string](system, ref, "msg")
	t.Log(f.OnlyResult())
}
