package fiber

import "github.com/gofiber/fiber/v2"

type Router interface {
	Use(args ...interface{}) Router

	Get(path string, handlers ...Handler) Router
	Head(path string, handlers ...Handler) Router
	Post(path string, handlers ...Handler) Router
	Put(path string, handlers ...Handler) Router
	Delete(path string, handlers ...Handler) Router
	Connect(path string, handlers ...Handler) Router
	Options(path string, handlers ...Handler) Router
	Trace(path string, handlers ...Handler) Router
	Patch(path string, handlers ...Handler) Router

	Add(method, path string, handlers ...Handler) Router
	Static(prefix, root string, config ...fiber.Static) Router
	All(path string, handlers ...Handler) Router

	Group(prefix string, handlers ...Handler) Router

	Route(prefix string, fn func(router Router), name ...string) Router

	Mount(prefix string, fiber *App) Router

	Name(name string) Router
}
