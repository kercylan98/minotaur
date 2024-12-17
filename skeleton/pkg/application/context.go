package application

import (
	"context"
	gofiber "github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/skeleton/pkg/fiber"
	"reflect"
)

func New() *Context {
	ctx := &Context{
		ctx: context.Background(),
	}
	ctx.fiber = fiber.New(ctx, gofiber.New())
	return ctx
}

type Context struct {
	ctx            context.Context
	fiber          *fiber.Server[*Context]
	services       map[reflect.Type][]Service
	repositoryList map[reflect.Type][]Repository
	controllers    []Controller
}

func (c *Context) Run() (err error) {
	if err = runRepositoryList(c); err != nil {
		return
	}

	if err = runServices(c); err != nil {
		return
	}

	if err = runControllers(c); err != nil {
		return
	}

	c.fiber.Listen(":8080")

	return nil
}

func (c *Context) Fiber() *fiber.Server[*Context] {
	return c.fiber
}
