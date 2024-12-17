package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/application"
)

func NewApp(app *application.Context, fiberApp *fiber.App) *App {
	return &App{
		App: fiberApp,
		app: app,
	}
}

type App struct {
	*fiber.App
	app *application.Context
}

func (a *App) Use(args ...interface{}) Router {
	a.App.Use(args)
	return a
}

func (a *App) Get(path string, handlers ...Handler) Router {
	a.App.Get(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Head(path string, handlers ...Handler) Router {
	a.App.Head(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Post(path string, handlers ...Handler) Router {
	a.App.Post(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Put(path string, handlers ...Handler) Router {
	a.App.Put(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Delete(path string, handlers ...Handler) Router {
	a.App.Delete(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Connect(path string, handlers ...Handler) Router {
	a.App.Connect(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Options(path string, handlers ...Handler) Router {
	a.App.Options(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Trace(path string, handlers ...Handler) Router {
	a.App.Trace(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Patch(path string, handlers ...Handler) Router {
	a.App.Patch(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Add(method, path string, handlers ...Handler) Router {
	a.App.Add(method, path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Static(prefix, root string, config ...fiber.Static) Router {
	a.App.Static(prefix, root, config...)
	return a
}

func (a *App) All(path string, handlers ...Handler) Router {
	a.App.All(path, covertFiberHandler(a.app, handlers)...)
	return a
}

func (a *App) Group(prefix string, handlers ...Handler) Router {
	group := &Group{
		app:   a.app,
		group: a.App.Group(prefix, covertFiberHandler(a.app, handlers)...),
	}
	return group
}

func (a *App) Route(prefix string, fn func(router Router), name ...string) Router {
	group := a.Group(prefix)
	if len(name) > 0 {
		group.Name(name[0])
	}
	fn(group)
	return group
}

func (a *App) Mount(prefix string, fiber *App) Router {
	a.App.Mount(prefix, fiber.App)
	return a
}

func (a *App) Name(name string) Router {
	a.App.Name(name)
	return a
}
