package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/stream"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/examples/internal/chat-room/internal/manager"
	"github.com/kercylan98/minotaur/examples/internal/chat-room/internal/user"
)

func main() {
	fiberApp := fiber.New()
	system := vivid.NewActorSystem()
	defer system.Shutdown(true)

	roomManager := system.ActorOfF(func() vivid.Actor {
		return manager.New()
	})

	fiberApp.Get("/ws", stream.NewFiberWebSocketHandler(fiberApp, system, stream.FunctionalConfigurator(func(c *stream.Configuration) {
		c.WithPerformance(vivid.FunctionalActorPerformance(user.New(roomManager).OnReceive))
	})))

	if err := fiberApp.Listen(":8080"); err != nil {
		panic(err)
	}
}
