package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
)

type ExampleMessage struct {
	Content string
}

func main() {
	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	tell(system)
	ask(system)
	futureAsk(system)
}

func tell(system *vivid.ActorSystem) {
	receiver := system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *ExampleMessage:
				fmt.Println(m.Content)
			}
		})
	})

	system.Tell(receiver, &ExampleMessage{
		Content: "Hi",
	})
}

func ask(system *vivid.ActorSystem) {
	receiver := system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *ExampleMessage:
				switch m.Content {
				case "你好":
					ctx.Reply(&ExampleMessage{Content: "你好"})
				case "能帮我一下吗？":
					ctx.Reply(&ExampleMessage{Content: "不能，谢谢！"})
				}
			}
		})
	})

	system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *vivid.OnLaunch:
				ctx.Ask(receiver, &ExampleMessage{Content: "你好"})
			case *ExampleMessage:
				switch m.Content {
				case "你好":
					ctx.Reply(&ExampleMessage{Content: "能帮我一下吗？"})
				}
			}
		})
	})
}

func futureAsk(system *vivid.ActorSystem) {
	receiver := system.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case *ExampleMessage:
				ctx.Reply(m)
			}
		})
	})

	future := system.FutureAsk(receiver, &ExampleMessage{
		Content: "Hi",
	})

	if _, err := future.Result(); err != nil {
		panic(err)
	}
}
