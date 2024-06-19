package vivid_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

func ExampleActorOf() {
	system := vivid.NewActorSystem("example")
	defer system.Shutdown()
	ref := vivid.ActorOf[*vivid.IneffectiveActor](&system)
	fmt.Println(ref != nil)
	// Output: true
}

func ExampleActorOfT() {
	system := vivid.NewActorSystem("example")
	defer system.Shutdown()
	typed := vivid.ActorOfT[*vivid.PrintlnActor, vivid.PrintlnActorTyped](&system)
	typed.Println("Hello, World!")
	// Output: Hello, World!
}
