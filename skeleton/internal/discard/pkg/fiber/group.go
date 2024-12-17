package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/application"
)

type Group struct {
	app   *application.Context
	group fiber.Router
}

func (g *Group) Use(args ...interface{}) Router {
	g.group.Use(args)
	return g
}

func (g *Group) Get(path string, handlers ...Handler) Router {
	g.group.Get(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Head(path string, handlers ...Handler) Router {
	g.group.Head(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Post(path string, handlers ...Handler) Router {
	g.group.Post(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Put(path string, handlers ...Handler) Router {
	g.group.Put(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Delete(path string, handlers ...Handler) Router {
	g.group.Delete(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Connect(path string, handlers ...Handler) Router {
	g.group.Connect(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Options(path string, handlers ...Handler) Router {
	g.group.Options(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Trace(path string, handlers ...Handler) Router {
	g.group.Trace(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Patch(path string, handlers ...Handler) Router {
	g.group.Patch(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Add(method, path string, handlers ...Handler) Router {
	g.group.Add(method, path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Static(prefix, root string, config ...fiber.Static) Router {
	g.group.Static(prefix, root, config...)
	return g
}

func (g *Group) All(path string, handlers ...Handler) Router {
	g.group.All(path, covertFiberHandler(g.app, handlers)...)
	return g
}

func (g *Group) Group(prefix string, handlers ...Handler) Router {
	group := &Group{
		app:   g.app,
		group: g.group.Group(prefix, covertFiberHandler(g.app, handlers)...),
	}
	return group
}

func (g *Group) Route(prefix string, fn func(router Router), name ...string) Router {
	group := g.Group(prefix)
	if len(name) > 0 {
		group.Name(name[0])
	}
	fn(group)
	return group
}

func (g *Group) Mount(prefix string, fiber *App) Router {
	g.group.Mount(prefix, fiber.App)
	return g
}

func (g *Group) Name(name string) Router {
	g.group.Name(name)
	return g
}
