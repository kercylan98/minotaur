package rpc

import (
	"fmt"
)

// Route 是一个 RPC 路由的接口，该接口用于定义一个 RPC 路由
type Route = string

// RouteHandler 是一个 RPC 路由处理器的类型，该类型用于定义一个 RPC 路由处理器
type RouteHandler func(ctx Context) error

// Router 是 RPC 路由器的接口，该接口被用于 RPC 服务路由的注册及调用路由的匹配
type Router interface {
	// Register 用于注册一个 RPC 服务路由，当存在相同路由时会发生 panic，该路由处理函数接收一个上下文参数
	Register(route Route, handler RouteHandler)

	// Match 用于根据 RPC 服务名匹配对应的路由处理器，如果找到则返回对应的处理器，否则返回 nil
	Match(route Route) RouteHandler
}

// NewRouter 用于创建一个新的 RPC 路由器
func NewRouter() Router {
	return &router{
		handlers: make(map[Route]RouteHandler),
	}
}

type router struct {
	handlers map[Route]RouteHandler
}

func (r *router) Register(route Route, handler RouteHandler) {
	if _, ok := r.handlers[route]; ok {
		panic(fmt.Errorf("rpc: route %s already registered", route))
	}
	r.handlers[route] = handler
}

func (r *router) Match(route Route) RouteHandler {
	return r.handlers[route]
}
