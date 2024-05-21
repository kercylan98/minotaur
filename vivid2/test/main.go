package main

import (
	"fmt"
	vivid "github.com/kercylan98/minotaur/vivid2"
)

type TestActor struct {
}

func (t *TestActor) OnReceive(ctx vivid.MessageContext) {
	switch v := ctx.GetMessage().(type) {
	case string:
		fmt.Println(v)
	case int:
		ctx.Reply(v + 1)
	}
}

func main() {
	system := vivid.NewActorSystem("test-system")

	actorRef := vivid.ActorOf[*TestActor](&system)

	actorRef.Tell("Hello, World!")

	fmt.Println(actorRef.Ask(1))

	system.Shutdown()
}
