package main

import (
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/components/fiber"
	"github.com/kercylan98/minotaur/skeleton/pkg/session"
)

func main() {
	app := application.New()
	app.SetupComponents(
		fiber.NewFiberComponent(":8080"),
		fiber.NewFiberWebSocketComponent("/websocket", socket.FunctionalProvider(func() socket.Actor {
			return session.New()
		})),
	)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
