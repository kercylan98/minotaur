package fiber

import (
	"github.com/gofiber/fiber/v2"
)

func NewContext[CTX any](ctx CTX, fiberCtx *fiber.Ctx) *Context[CTX] {
	return &Context[CTX]{
		ctx: ctx,
		Ctx: fiberCtx,
	}
}

type Context[CTX any] struct {
	ctx CTX
	*fiber.Ctx
}
