package stream_test

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/stream"
	"github.com/kercylan98/minotaur/engine/vivid"
	"testing"
)

func TestStream(t *testing.T) {
	system := vivid.NewActorSystem()

	fiberApp := fiber.New()
	fiberApp.Get("/ws", stream.NewFiberWebSocketHandler(fiberApp, system, stream.FunctionalConfigurator(func(c *stream.Configuration) {
		var writer vivid.ActorRef
		c.WithPerformance(vivid.ActorFunctionalPerformance(func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case stream.Writer:
				writer = m
				ctx.Tell(writer, stream.NewPacketDC([]byte("Hello!"), websocket.TextMessage))
			case *stream.Packet:
				ctx.Tell(writer, m) // echo
			}
		}))
	})))

	if err := fiberApp.Listen(":8888"); err != nil {
		panic(err)
	}
}
