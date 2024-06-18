package vivid_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"testing"
)

type TestActor struct {
	vivid.ActorRef
}

type TestActorTyped interface {
	vivid.ActorTyped
	Println(string)
}

func (t *TestActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case vivid.OnBoot:
		t.ActorRef = ctx
	case string:
		fmt.Println(m)
	}
}

func (t *TestActor) Println(s string) {
	t.Tell(s)
}

func TestActorOfT(t *testing.T) {
	system := vivid.NewActorSystem("test")
	ref := vivid.ActorOfT[*TestActor, TestActorTyped](&system)

	ref.Println("Hello, World!")

	system.Shutdown()
}
