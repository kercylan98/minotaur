package test

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/typed/tm"
	"sync"
	"testing"
)

type TestActor struct {
}

func (t *TestActor) OnReceive(ctx vivid.ActorContext) {
}

func (t *TestActor) Say(ctx vivid.ActorContext, message *Request) {
	fmt.Println(message.Message)
}

func (t *TestActor) Call(ctx vivid.ActorContext, message *Request, responder tm.AskResponder[*Response]) {
	responder.Reply(&Response{
		Message: "hello",
	})
}

func (t *TestActor) Ping(ctx vivid.ActorContext, message *Request) (*Response, error) {
	return &Response{
		Message: "pong",
	}, nil
}

func TestTyped(t *testing.T) {
	wait := new(sync.WaitGroup)
	wait.Add(2)
	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	ref := system.ActorOf(FunctionalTestActorTypedProvider(func() TestActorTypedInterface {
		return new(TestActor)
	}))

	typed := NewTestActorTyped(ref)

	system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *vivid.OnLaunch:
				typed.Say(ctx, &Request{Message: "hello"})
				typed.Call(ctx, &Request{Message: "hi"})
				result, err := typed.Ping(ctx, &Request{Message: "ping"})
				if err != nil {
					panic(err)
				}

				fmt.Println(result.Message)
				wait.Done()
			case *Response:
				fmt.Println(m.Message)
				wait.Done()
			}
		})
	})

	wait.Wait()

}
