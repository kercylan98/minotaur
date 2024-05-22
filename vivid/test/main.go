package main

import (
	"fmt"
	vivid "github.com/kercylan98/minotaur/vivid"
	"reflect"
)

type TestFreeActor struct {
	name string
}

func main() {
	system := vivid.NewActorSystem("test-system")

	actorRef := vivid.FreeActorOf[*TestFreeActor](&system, vivid.NewFreeActorOptions[*TestFreeActor]().
		WithName("free").
		WithProps(func(f *vivid.FreeActor[*TestFreeActor]) {
			f.GetActor().name = "free"
		}).
		WithMessageHook(func(ctx vivid.MessageContext) bool {
			fmt.Println(vivid.GetActor[*TestFreeActor](ctx).name, reflect.TypeOf(ctx.GetMessage()).String(), ctx.GetMessage())
			return false
		}),
	)
	actorRef.Tell("Hello, World!")

	system.Shutdown()
}
