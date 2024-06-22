package vivid

import (
	"fmt"
	"testing"
)

type TestActor struct {
}

func (t *TestActor) OnReceive(ctx ActorContext) {
	switch ctx.Message().(type) {
	case OnBoot:
		fmt.Println("OnBoot", ctx)
	}
}

func TestNewActorSystem(t *testing.T) {
	sys := NewActorSystem("sys")

	sys.ActorOf(func(options *ActorOptions) Actor {
		options.WithName("test")
		return new(TestActor)
	})

	sys.Shutdown()

}
