package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/cmd/gateway/internal/app/pkg/session"
)

func main() {
	actorSystem := vivid.NewActorSystem()
	socketFactory := socket.NewFactory(actorSystem)
	fiberApp := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	fiberApp.Get("/ws", websocket.New(func(conn *websocket.Conn) {
		socket.ProduceFiberSocketV2(socketFactory, conn, session.New(), socket.FunctionalFiberV2Configurator(func(config *socket.FiberV2Configuration) {
			config.WithContextInitializer(socket.FunctionalContextEditor(func(ctx *socket.Context) {
				//ctx.Set(AuthToken, conn.Query("auth_token"))
			}))
		}))
	}))
}
