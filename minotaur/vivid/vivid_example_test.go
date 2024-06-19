package vivid_test

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

func ExampleBehaviorOf() {
	vivid.BehaviorOf(func(ctx vivid.MessageContext, message string) {
		println(message)
	})

	// Output:
}

func ExampleActorOf() {
	system := vivid.NewActorSystem("example")
	defer system.Shutdown()

	vivid.ActorOf[*vivid.IneffectiveActor](&system)

	vivid.ActorOf(&system, vivid.NewActorOptions[*vivid.IneffectiveActor]().
		WithName("actor"),
	)

	// Output:
}

func ExampleActorOfT() {
	system := vivid.NewActorSystem("example")
	defer system.Shutdown()

	typedA := vivid.ActorOfT[*vivid.PrintlnActor, vivid.PrintlnActorTyped](&system)
	typedA.Println("Hello, World!")

	typedB := vivid.ActorOfT[*vivid.PrintlnActor, vivid.PrintlnActorTyped](&system,
		vivid.NewActorOptions[*vivid.PrintlnActor]().
			WithName("actor"),
	)
	typedB.Println("Hello, World!")

	// Output: Hello, World!
	// Hello, World!
}

func ExampleActorOfF() {
	system := vivid.NewActorSystem("example")
	defer system.Shutdown()

	vivid.ActorOfF[*vivid.IneffectiveActor](&system)

	vivid.ActorOfF(&system, func(options *vivid.ActorOptions[*vivid.IneffectiveActor]) {
		options.WithName("actor")
	})

	// Output:
}

func ExampleActorOfFT() {
	system := vivid.NewActorSystem("example")
	defer system.Shutdown()

	typedA := vivid.ActorOfFT[*vivid.PrintlnActor, vivid.PrintlnActorTyped](&system)
	typedA.Println("Hello, World!")

	typedB := vivid.ActorOfFT[*vivid.PrintlnActor, vivid.PrintlnActorTyped](&system,
		func(options *vivid.ActorOptions[*vivid.PrintlnActor]) {
			options.WithName("actor")
		},
	)
	typedB.Println("Hello, World!")

	// Output: Hello, World!
	// Hello, World!
}

func ExampleActorOfI() {
	system := vivid.NewActorSystem("example")
	defer system.Shutdown()

	instance := new(vivid.IneffectiveActor)

	vivid.ActorOfI[*vivid.IneffectiveActor](&system, instance)

	vivid.ActorOfI(&system, instance, func(options *vivid.ActorOptions[*vivid.IneffectiveActor]) {
		options.WithName("actor")
	})

	// Output:
}

func ExampleActorOfIT() {
	system := vivid.NewActorSystem("example")
	defer system.Shutdown()

	instance := new(vivid.PrintlnActor)

	typedA := vivid.ActorOfIT[*vivid.PrintlnActor, vivid.PrintlnActorTyped](&system, instance)
	typedA.Println("Hello, World!")

	typedB := vivid.ActorOfIT[*vivid.PrintlnActor, vivid.PrintlnActorTyped](&system, instance,
		func(options *vivid.ActorOptions[*vivid.PrintlnActor]) {
			options.WithName("actor")
		},
	)
	typedB.Println("Hello, World!")

	// Output: Hello, World!
	// Hello, World!
}
