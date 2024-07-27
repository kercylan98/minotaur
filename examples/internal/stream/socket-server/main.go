package main

import (
	"github.com/kercylan98/minotaur/engine/stream"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/behavior"
	"github.com/panjf2000/gnet/v2"
)

func main() {
	system := vivid.NewActorSystem()

	if err := gnet.Run(stream.NewGNETEventHandler(system, stream.FunctionalConfigurator(func(c *stream.Configuration) {
		c.WithPerformance(vivid.FunctionalStatefulActorPerformance(
			func() behavior.Performance[vivid.ActorContext] {
				var writer stream.Writer
				return vivid.FunctionalActorPerformance(func(ctx vivid.ActorContext) {
					switch m := ctx.Message().(type) {
					case stream.Writer:
						writer = m
					case *stream.Packet:
						ctx.Tell(writer, m) // echo
					}
				})
			}).
			Stateful(),
		)
	})), "tcp://127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
