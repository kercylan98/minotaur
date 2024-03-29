package router_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server/router"
	"testing"
)

func TestMultistage_Match(t *testing.T) {
	r := router.NewMultistage[func()]()

	r.Sub("System").Route("Heartbeat", func() {
		fmt.Println("Heartbeat")
	})

	r.Route("ServerTime", func() {
		fmt.Println("ServerTime")
	})

	r.Register("System", "Network", "Ping")(func() {
		fmt.Println("Ping")
	})

	r.Register("System", "Network", "Echo").Bind(onEcho)

	r.Match("System", "Heartbeat")()
	r.Match("ServerTime")()
	r.Match("System", "Network", "Ping")()
	r.Match("System", "Network", "Echo")()
	fmt.Println(r.Match("None") == nil)
}

func onEcho() {
	fmt.Println("Echo")
}
