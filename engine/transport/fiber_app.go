package transport

import "github.com/gofiber/fiber/v2"

func newFiberApp(app *fiber.App) *FiberApp {
	return &FiberApp{
		App: app,
	}
}

type FiberApp struct {
	*fiber.App
}
