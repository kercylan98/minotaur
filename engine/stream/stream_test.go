package stream

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/vivid"
	"testing"
)

func TestStream(t *testing.T) {

	system := vivid.NewActorSystem()

	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) (err error) {
		if !websocket.IsWebSocketUpgrade(c) {
			return fiber.ErrUpgradeRequired
		}
		return c.Next()
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		stream := NewStream()
	}))
}
