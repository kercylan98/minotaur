package fiber

import (
	"github.com/gofiber/fiber/v2"
)

type Handler[CTX any] func(ctx *Context[CTX]) error

func covertFiberHandler[CTX any](ctx CTX, handlers []Handler[CTX]) (result []fiber.Handler) {
	result = make([]fiber.Handler, len(handlers))
	for i, handler := range handlers {
		result[i] = func(fiberCtx *fiber.Ctx) error {
			return handler(NewContext(ctx, fiberCtx))
		}
	}
	return
}

func UseFiberHandler[CTX any](handler fiber.Handler) Handler[CTX] {
	return func(ctx *Context[CTX]) error {
		return handler(ctx.Ctx)
	}
}
