package router

import (
	"fmt"
	"reflect"
)

func NewLevel2Router[Route comparable, Handle any]() *Level2Router[Route, Handle] {
	return &Level2Router[Route, Handle]{
		routes: map[Route]map[Route]Handle{},
	}
}

// Level2Router 支持二级分类的路由器
type Level2Router[Route comparable, Handle any] struct {
	routes map[Route]map[Route]Handle
}

func (slf *Level2Router[Route, Handle]) Route(topRoute Route, route Route, handleFunc Handle) {
	if reflect.TypeOf(handleFunc).Kind() != reflect.Func {
		panic(fmt.Errorf("route[%v] registration failed, handle must be a function type", route))
	}
	routes, exist := slf.routes[topRoute]
	if !exist {
		routes = map[Route]Handle{}
		slf.routes[topRoute] = routes
	}

	_, exist = routes[route]
	if exist {
		panic(fmt.Errorf("the route[%v:%v] has already been registered, duplicate registration is not allowed", topRoute, route))
	}
	routes[route] = handleFunc
}

func (slf *Level2Router[Route, Handle]) Match(topRoute Route, route Route) Handle {
	return slf.routes[topRoute][route]
}
