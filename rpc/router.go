package rpc

import multistageRouter "github.com/kercylan98/minotaur/server/router"

type Route = any
type Reader func(dst any)

func (r Reader) ReadTo(dst any) {
	r(dst)
}

type RouteHandler func(reader Reader)

type Router interface {
	// Route 为特定路由绑定处理函数
	Route(route Route, handler RouteHandler)
	// Register 注册多级路由并返回一个绑定处理函数的函数
	Register(routes ...Route) multistageRouter.MultistageBind[RouteHandler]
	// Call 调用特定路由的处理函数
	Call(routes ...Route) func(request any) error
}

type RouteMatcher interface {
	// Match 匹配特定路由
	Match(routes ...Route) RouteHandler
}

type router struct {
	application *Application
	*multistageRouter.Multistage[RouteHandler]
}

func (r *router) init(application *Application) *router {
	r.application = application
	r.Multistage = multistageRouter.NewMultistage[RouteHandler]()
	return r
}

func (r *router) Call(routes ...Route) func(request any) error {
	return func(request any) error {
		return r.application.core.OnCall(routes...)(request)
	}
}
