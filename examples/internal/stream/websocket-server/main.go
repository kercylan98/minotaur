package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/stream"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/behavior"
)

func main() {
	system := vivid.NewActorSystem()

	fiberApp := fiber.New()
	fiberApp.Get("/ws", stream.NewFiberWebSocketHandler(fiberApp, system, stream.FunctionalConfigurator(func(c *stream.Configuration) {
		c.WithPerformance(vivid.FunctionalStatefulActorPerformance(
			func() behavior.Performance[vivid.ActorContext] {
				var writer stream.Writer
				return vivid.FunctionalActorPerformance(func(ctx vivid.ActorContext) {
					switch m := ctx.Message().(type) {
					case stream.Writer:
						writer = m
						ctx.Tell(writer, stream.NewPacketDC([]byte("Hello!"), websocket.TextMessage))
					case *stream.Packet:
						ctx.Tell(writer, m) // echo
					}
				})
			}).
			Stateful(),
		)
	})))

	if err := fiberApp.Listen(":8888"); err != nil {
		panic(err)
	}
}
