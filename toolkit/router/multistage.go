package router

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/log"
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
func (slf MultistageBind[HandleFunc]) Bind(handleFunc HandleFunc) {
	slf(handleFunc)
}

// Multistage 支持多级分类的路由器
type Multistage[HandleFunc any] struct {
	routes map[any]HandleFunc
	subs   map[any]*Multistage[HandleFunc]
	tag    any
	trim   func(route any) any
}

// Register 注册路由是结合 Sub 和 Route 的快捷方式，用于一次性注册多级路由
//   - 该函数将返回一个注册函数，可通过调用其将路由绑定到特定处理函数，例如：router.Register("a", "b").Bind(onExec())
func (slf *Multistage[HandleFunc]) Register(routes ...any) MultistageBind[HandleFunc] {
	return func(handleFunc HandleFunc) {
		router := slf
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
func (slf *Multistage[HandleFunc]) Route(route any, handleFunc HandleFunc) {
	trim := route
	if slf.trim != nil {
		trim = slf.trim(route)
	}
	if reflect.TypeOf(handleFunc).Kind() != reflect.Func {
		panic(fmt.Errorf("route[%v] registration failed, handle must be a function type", route))
	}

	_, exist := slf.routes[trim]
	if exist {
		panic(fmt.Errorf("the route[%v] has already been registered, duplicate registration is not allowed", route))
	}

	_, exist = slf.subs[trim]
	if exist {
		panic(fmt.Errorf("the route[%v] has already been registered, duplicate registration is not allowed", route))
	}

	slf.routes[trim] = handleFunc
	if slf.tag == nil {
		log.Info("Router", log.String("Type", "Multistage"), log.String("Route", fmt.Sprintf("%v", route)))
	} else {
		log.Info("Router", log.String("Type", "Multistage"), log.String("Route", fmt.Sprintf("%v > %v", slf.tag, route)))
	}
}

// Match 匹配已绑定处理函数的路由，返回处理函数
//   - 如果未找到将会返回空指针
func (slf *Multistage[HandleFunc]) Match(routes ...any) HandleFunc {
	if len(routes) == 0 {
		panic(fmt.Errorf("the route cannot be empty"))
	}
	var handleFunc HandleFunc
	var exist bool
	var router *Multistage[HandleFunc]
	for i, route := range routes {
		handleFunc, exist = slf.routes[route]
		if exist {
			return handleFunc
		}
		router, exist = slf.subs[route]
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
func (slf *Multistage[HandleFunc]) Sub(route any) *Multistage[HandleFunc] {
	trim := route
	if slf.trim != nil {
		trim = slf.trim(route)
	}
	_, exist := slf.routes[trim]
	if exist {
		panic(fmt.Errorf("the route[%v] has already been registered, cannot be used as a sub-router", route))
	}

	router, exist := slf.subs[trim]
	if !exist {
		router = NewMultistage[HandleFunc]()
		if slf.tag == nil {
			router.tag = fmt.Sprintf("%v", route)
		} else {
			router.tag = fmt.Sprintf("%v > %v", slf.tag, route)
		}
		slf.subs[trim] = router
	}
	return router
}
