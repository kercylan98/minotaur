package router

import (
	"fmt"
	"reflect"
)

// NewMultistage 创建一个支持多级分类的路由器
func NewMultistage[HandleFunc any](options ...MultistageOption[HandleFunc]) *Multistage[HandleFunc] {
	r := &Multistage[HandleFunc]{
		routes: make(map[any]HandleFunc),
		subs:   make(map[any]*Multistage[HandleFunc]),
	}
	for _, option := range options {
		option(r)
	}
	return r
}

// MultistageBind 多级分类路由绑定函数
type MultistageBind[HandleFunc any] func(HandleFunc)

// Bind 将处理函数绑定到预设的路由中
func (b MultistageBind[HandleFunc]) Bind(handleFunc HandleFunc) {
	b(handleFunc)
}

// Multistage 支持多级分类的路由器
type Multistage[HandleFunc any] struct {
	parent *Multistage[HandleFunc]
	routes map[any]HandleFunc
	subs   map[any]*Multistage[HandleFunc]
	tag    any
	route  any
	trim   func(route any) any
}

// Register 注册路由是结合 Sub 和 Route 的快捷方式，用于一次性注册多级路由
//   - 该函数将返回一个注册函数，可通过调用其将路由绑定到特定处理函数，例如：router.Register("a", "b").Bind(onExec())
func (m *Multistage[HandleFunc]) Register(routes ...any) MultistageBind[HandleFunc] {
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
func (m *Multistage[HandleFunc]) Route(route any, handleFunc HandleFunc) {
	trim := route
	if m.trim != nil {
		trim = m.trim(route)
	}
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
func (m *Multistage[HandleFunc]) Match(routes ...any) HandleFunc {
	if len(routes) == 0 {
		panic(fmt.Errorf("the route cannot be empty"))
	}
	var handleFunc HandleFunc
	var exist bool
	var router *Multistage[HandleFunc]
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
func (m *Multistage[HandleFunc]) Sub(route any) *Multistage[HandleFunc] {
	trim := route
	if m.trim != nil {
		trim = m.trim(route)
	}
	_, exist := m.routes[trim]
	if exist {
		panic(fmt.Errorf("the route[%v] has already been registered, cannot be used as a sub-router", route))
	}

	router, exist := m.subs[trim]
	if !exist {
		router = NewMultistage[HandleFunc]()
		router.parent = m
		router.route = route
		if m.tag == nil {
			router.tag = fmt.Sprintf("%v", route)
		} else {
			router.tag = fmt.Sprintf("%v > %v", m.tag, route)
		}
		m.subs[trim] = router
	}
	return router
}

// GetPrefixRoutes 获取路由前缀，当该路由器为子路由器时，将会返回所有父级路由
func (m *Multistage[HandleFunc]) GetPrefixRoutes() []any {
	var routes []any
	if m.route != nil {
		routes = append(routes, m.route)
	}
	if m.parent != nil {
		routes = append(m.parent.GetPrefixRoutes(), routes...)
	}
	return routes

}

// GetRoutes 获取所有已注册的路由
func (m *Multistage[HandleFunc]) GetRoutes() [][]any {
	var routes [][]any
	for route := range m.routes {
		routes = append(routes, append(m.GetPrefixRoutes(), route))
	}
	for _, sub := range m.subs {
		routes = append(routes, sub.GetRoutes()...)
	}
	return routes
}
