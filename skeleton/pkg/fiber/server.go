package fiber

import (
	"github.com/gofiber/fiber/v2"
)

func New[CTX any](ctx CTX, fiberApp *fiber.App) *Server[CTX] {
	return &Server[CTX]{
		ctx: ctx,
		App: fiberApp,
	}
}

type Server[CTX any] struct {
	*fiber.App
	ctx CTX
}

func (a *Server[CTX]) Use(args ...interface{}) Router[CTX] {
	a.App.Use(args)
	return a
}

func (a *Server[CTX]) Get(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Get(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Head(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Head(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Post(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Post(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Put(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Put(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Delete(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Delete(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Connect(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Connect(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Options(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Options(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Trace(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Trace(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Patch(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Patch(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Add(method, path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.Add(method, path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Static(prefix, root string, config ...fiber.Static) Router[CTX] {
	a.App.Static(prefix, root, config...)
	return a
}

func (a *Server[CTX]) All(path string, handlers ...Handler[CTX]) Router[CTX] {
	a.App.All(path, covertFiberHandler[CTX](a.ctx, handlers)...)
	return a
}

func (a *Server[CTX]) Group(prefix string, handlers ...Handler[CTX]) Router[CTX] {
	group := &Group[CTX]{
		ctx:   a.ctx,
		group: a.App.Group(prefix, covertFiberHandler[CTX](a.ctx, handlers)...),
	}
	return group
}

func (a *Server[CTX]) Route(prefix string, fn func(router Router[CTX]), name ...string) Router[CTX] {
	group := a.Group(prefix)
	if len(name) > 0 {
		group.Name(name[0])
	}
	fn(group)
	return group
}

func (a *Server[CTX]) Mount(prefix string, fiber *Server[CTX]) Router[CTX] {
	a.App.Mount(prefix, fiber.App)
	return a
}

func (a *Server[CTX]) Name(name string) Router[CTX] {
	a.App.Name(name)
	return a
}
