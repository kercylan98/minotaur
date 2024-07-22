package stream_test

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/stream"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/behavior"
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"testing"
)

type FiberWebSocketWrapper struct {
	*websocket.Conn
}

func (f *FiberWebSocketWrapper) Write(packet *stream.Packet) error {
	return f.WriteMessage(packet.Context().(int), packet.Data())
}

func TestStream(t *testing.T) {
	system := vivid.NewActorSystem()

	fiberApp := fiber.New()
	fiberApp.Use("/ws", func(c *fiber.Ctx) (err error) {
		if !websocket.IsWebSocketUpgrade(c) {
			return fiber.ErrUpgradeRequired
		}
		return c.Next()
	})

	fiberApp.Get("/ws", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)

		// 危险：非并发安全
		ref := system.ActorOfF(func() vivid.Actor {
			return stream.New(&FiberWebSocketWrapper{c}, stream.FunctionalConfigurator(func(c *stream.Configuration) {
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
			}))
		}, func(descriptor *vivid.ActorDescriptor) {
			descriptor.WithSupervisionStrategyProvider(supervision.FunctionalStrategyProvider(func() supervision.Strategy {
				return supervision.StopStrategy()
			}))
		})

		defer system.Terminate(ref, true)

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				system.Tell(ref, err)
				break
			}
			system.Tell(ref, stream.NewPacketDC(msg, mt))
		}
	}))

	if err := fiberApp.Listen(":8888"); err != nil {
		panic(err)
	}
}
