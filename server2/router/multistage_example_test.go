package router_test

import "github.com/kercylan98/minotaur/server/router"

func ExampleNewMultistage() {
	router.NewMultistage[func()]()

	// Output:
}

func ExampleMultistage_Register() {
	r := router.NewMultistage[func()]()

	r.Register("System", "Network", "Ping")(func() {
		// ...
	})

	// Output:
}

func ExampleMultistage_Sub() {
	r := router.NewMultistage[func()]()

	r.Sub("System").Route("Heartbeat", func() {
		// ...
	})

	// Output:
}

func ExampleMultistage_Route() {
	r := router.NewMultistage[func()]()

	r.Route("ServerTime", func() {
		// ...
	})

	// Output:
}

func ExampleMultistage_Match() {
	r := router.NewMultistage[func()]()

	r.Route("ServerTime", func() {})
	r.Register("System", "Network", "Ping").Bind(func() {})

	r.Match("ServerTime")()
	r.Match("System", "Network", "Ping")()

	// Output:
}
