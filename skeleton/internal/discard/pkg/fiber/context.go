package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/application"
)

func NewContext(app *application.Context, ctx *fiber.Ctx) *Context {
	return &Context{
		Context: app,
		Ctx:     ctx,
	}
}

type Context struct {
	*application.Context
	*fiber.Ctx
}
