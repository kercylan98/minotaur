package stream_test

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/stream"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/behavior"
	"testing"
	"time"
)

func TestStream(t *testing.T) {
	system := vivid.NewActorSystem()

	fiberApp := fiber.New()
	fiberApp.Get("/ws", stream.NewFiberWebSocketHandler(fiberApp, system, stream.FunctionalConfigurator(func(c *stream.Configuration) {
		var writer vivid.ActorRef
		c.WithPerformance(behavior.FunctionalPerformance[vivid.ActorContext](func(ctx vivid.ActorContext) {
			switch m := ctx.Message().(type) {
			case vivid.ActorRef:
				writer = m
				ctx.Tell(writer, stream.NewPacketDC([]byte("Hello!"), websocket.TextMessage))
			case *stream.Packet:
				ctx.Tell(writer, m) // echo
			}
		}))
	})))

	go func() {
		time.Sleep(time.Second * 3)
		fiberApp.Shutdown()
	}()

	if err := fiberApp.Listen(":8888"); err != nil {
		panic(err)
	}

	time.Sleep(time.Hour)
}
