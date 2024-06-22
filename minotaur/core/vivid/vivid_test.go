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
	case int:
		ctx.Reply(ctx.Message())
	}
}

func TestNewActorSystem(t *testing.T) {
	sys := NewActorSystem("sys")

	ref := sys.ActorOf(func(options *ActorOptions) Actor {
		options.WithName("test")
		return new(TestActor)
	})

	m, err := sys.Context().FutureAsk(ref, 123).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(m)

	sys.Shutdown()

}
