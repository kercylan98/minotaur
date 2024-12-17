package fiber

import "github.com/gofiber/fiber/v2"

type Router[CTX any] interface {
	Use(args ...interface{}) Router[CTX]

	Get(path string, handlers ...Handler[CTX]) Router[CTX]
	Head(path string, handlers ...Handler[CTX]) Router[CTX]
	Post(path string, handlers ...Handler[CTX]) Router[CTX]
	Put(path string, handlers ...Handler[CTX]) Router[CTX]
	Delete(path string, handlers ...Handler[CTX]) Router[CTX]
	Connect(path string, handlers ...Handler[CTX]) Router[CTX]
	Options(path string, handlers ...Handler[CTX]) Router[CTX]
	Trace(path string, handlers ...Handler[CTX]) Router[CTX]
	Patch(path string, handlers ...Handler[CTX]) Router[CTX]

	Add(method, path string, handlers ...Handler[CTX]) Router[CTX]
	Static(prefix, root string, config ...fiber.Static) Router[CTX]
	All(path string, handlers ...Handler[CTX]) Router[CTX]

	Group(prefix string, handlers ...Handler[CTX]) Router[CTX]

	Route(prefix string, fn func(router Router[CTX]), name ...string) Router[CTX]

	Mount(prefix string, fiber *Server[CTX]) Router[CTX]

	Name(name string) Router[CTX]
}
