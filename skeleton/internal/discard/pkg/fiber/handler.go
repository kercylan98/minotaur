package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/application"
)

type Handler func(ctx *Context) error

func covertFiberHandler(app *application.Context, handlers []Handler) (result []fiber.Handler) {
	result = make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		result[i] = func(ctx *fiber.Ctx) error {
			return handler(NewContext(app, ctx))
		}
	}
	return
}
