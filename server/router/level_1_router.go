package router

import (
	"fmt"
	"reflect"
)

func NewLevel1Router[Route comparable, Handle any]() *Level1Router[Route, Handle] {
	return &Level1Router[Route, Handle]{
		routes: map[Route]Handle{},
	}
}

// Level1Router 支持一级分类的路由器
type Level1Router[Route comparable, Handle any] struct {
	routes map[Route]Handle
}

func (slf *Level1Router[Route, Handle]) Route(route Route, handleFunc Handle) {
	if reflect.TypeOf(handleFunc).Kind() != reflect.Func {
		panic(fmt.Errorf("route[%v] registration failed, handle must be a function type", route))
	}
	_, exist := slf.routes[route]
	if exist {
		panic(fmt.Errorf("the route[%v] has already been registered, duplicate registration is not allowed", route))
	}
	slf.routes[route] = handleFunc
}

func (slf *Level1Router[Route, Handle]) Match(route Route) Handle {
	return slf.routes[route]
}
