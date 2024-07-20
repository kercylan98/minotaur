package transport

import "github.com/gofiber/fiber/v2"

type FiberService interface {
	OnInit(app *fiber.App)
}

type FunctionalFiberService func(app *fiber.App)

func (f FunctionalFiberService) OnInit(app *fiber.App) {
	f(app)
}
