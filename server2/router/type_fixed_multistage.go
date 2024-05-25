package router

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"reflect"
)

// NewTypeFixedMultistage 创建一个支持多级分类的路由器
func NewTypeFixedMultistage[T constraints.Ordered, HandleFunc any]() *TypeFixedMultistage[T, HandleFunc] {
	r := &TypeFixedMultistage[T, HandleFunc]{
		routes: make(map[T]HandleFunc),
		subs:   make(map[T]*TypeFixedMultistage[T, HandleFunc]),
	}
	return r
}

// TypeFixedMultistageBind 多级分类路由绑定函数，该路由器与 Multistage 的区别在于仅支持固定类型的路由
type TypeFixedMultistageBind[HandleFunc any] func(HandleFunc)

// Bind 将处理函数绑定到预设的路由中
func (b TypeFixedMultistageBind[HandleFunc]) Bind(handleFunc HandleFunc) {
	b(handleFunc)
}

// TypeFixedMultistage 支持多级分类的路由器
type TypeFixedMultistage[T constraints.Ordered, HandleFunc any] struct {
	parent *TypeFixedMultistage[T, HandleFunc]
	routes map[T]HandleFunc
	subs   map[T]*TypeFixedMultistage[T, HandleFunc]
	route  *T
}

// Register 注册路由是结合 Sub 和 Route 的快捷方式，用于一次性注册多级路由
//   - 该函数将返回一个注册函数，可通过调用其将路由绑定到特定处理函数，例如：router.Register("a", "b").Bind(onExec())
func (m *TypeFixedMultistage[T, HandleFunc]) Register(routes ...T) TypeFixedMultistageBind[HandleFunc] {
	return func(handleFunc HandleFunc) {
		router := m
		for i, route := range routes {
			if i == len(routes)-1 {
				router.Route(route, handleFunc)
			} else {
				router = router.Sub(route)
			}
		}
	}
}

// Route 为特定路由绑定处理函数，被绑定的处理函数将可以通过 Match 函数进行匹配
func (m *TypeFixedMultistage[T, HandleFunc]) Route(route T, handleFunc HandleFunc) {
	trim := route
	if reflect.TypeOf(handleFunc).Kind() != reflect.Func {
		panic(fmt.Errorf("route[%v] registration failed, handle must be a function type", route))
	}

	_, exist := m.routes[trim]
	if exist {
		panic(fmt.Errorf("the route[%v] has already been registered, duplicate registration is not allowed", route))
	}

	_, exist = m.subs[trim]
	if exist {
		panic(fmt.Errorf("the route[%v] has already been registered, duplicate registration is not allowed", route))
	}

	m.routes[trim] = handleFunc
}

// Match 匹配已绑定处理函数的路由，返回处理函数
//   - 如果未找到将会返回空指针
func (m *TypeFixedMultistage[T, HandleFunc]) Match(routes ...T) HandleFunc {
	if len(routes) == 0 {
		panic(fmt.Errorf("the route cannot be empty"))
	}
	var handleFunc HandleFunc
	var exist bool
	var router *TypeFixedMultistage[T, HandleFunc]
	for i, route := range routes {
		handleFunc, exist = m.routes[route]
		if exist {
			return handleFunc
		}
		router, exist = m.subs[route]
		if !exist {
			return handleFunc
		}
		if i == len(routes)-1 {
			return handleFunc
		}
		return router.Match(routes[i+1:]...)
	}
	return handleFunc
}

// Sub 获取子路由器
func (m *TypeFixedMultistage[T, HandleFunc]) Sub(route T) *TypeFixedMultistage[T, HandleFunc] {
	trim := route
	_, exist := m.routes[trim]
	if exist {
		panic(fmt.Errorf("the route[%v] has already been registered, cannot be used as a sub-router", route))
	}

	router, exist := m.subs[trim]
	if !exist {
		router = NewTypeFixedMultistage[T, HandleFunc]()
		router.parent = m
		router.route = &route
		m.subs[trim] = router
	}
	return router
}

// GetPrefixRoutes 获取路由前缀，当该路由器为子路由器时，将会返回所有父级路由
func (m *TypeFixedMultistage[T, HandleFunc]) GetPrefixRoutes() []T {
	var routes []T
	if m.route != nil {
		routes = append(routes, *m.route)
	}
	if m.parent != nil {
		routes = append(m.parent.GetPrefixRoutes(), routes...)
	}
	return routes

}

// GetRoutes 获取所有已注册的路由
func (m *TypeFixedMultistage[T, HandleFunc]) GetRoutes() [][]T {
	var routes [][]T
	for route := range m.routes {
		routes = append(routes, append(m.GetPrefixRoutes(), route))
	}
	for _, sub := range m.subs {
		routes = append(routes, sub.GetRoutes()...)
	}
	return routes
}
