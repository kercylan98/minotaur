package fiber

import (
	"github.com/gofiber/fiber/v2"
)

type Group[CTX any] struct {
	ctx   CTX
	group fiber.Router
}

func (g *Group[CTX]) Use(args ...interface{}) Router[CTX] {
	g.group.Use(args)
	return g
}

func (g *Group[CTX]) Get(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Get(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Head(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Head(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Post(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Post(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Put(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Put(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Delete(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Delete(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Connect(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Connect(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Options(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Options(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Trace(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Trace(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Patch(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Patch(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Add(method, path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.Add(method, path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Static(prefix, root string, config ...fiber.Static) Router[CTX] {
	g.group.Static(prefix, root, config...)
	return g
}

func (g *Group[CTX]) All(path string, handlers ...Handler[CTX]) Router[CTX] {
	g.group.All(path, covertFiberHandler(g.ctx, handlers)...)
	return g
}

func (g *Group[CTX]) Group(prefix string, handlers ...Handler[CTX]) Router[CTX] {
	group := &Group[CTX]{
		ctx:   g.ctx,
		group: g.group.Group(prefix, covertFiberHandler(g.ctx, handlers)...),
	}
	return group
}

func (g *Group[CTX]) Route(prefix string, fn func(router Router[CTX]), name ...string) Router[CTX] {
	group := g.Group(prefix)
	if len(name) > 0 {
		group.Name(name[0])
	}
	fn(group)
	return group
}

func (g *Group[CTX]) Mount(prefix string, fiber *Server[CTX]) Router[CTX] {
	g.group.Mount(prefix, fiber.App)
	return g
}

func (g *Group[CTX]) Name(name string) Router[CTX] {
	g.group.Name(name)
	return g
}
