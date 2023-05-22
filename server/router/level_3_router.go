package router

import (
	"fmt"
	"reflect"
)

func NewLevel3Router[Route comparable, Handle any]() *Level3Router[Route, Handle] {
	return &Level3Router[Route, Handle]{
		routes: map[Route]map[Route]map[Route]Handle{},
	}
}

// Level3Router 支持三级分类的路由器
type Level3Router[Route comparable, Handle any] struct {
	routes map[Route]map[Route]map[Route]Handle
}

func (slf *Level3Router[Route, Handle]) Route(topRoute Route, level2Route Route, route Route, handleFunc Handle) {
	if reflect.TypeOf(handleFunc).Kind() != reflect.Func {
		panic(fmt.Errorf("route[%v] registration failed, handle must be a function type", route))
	}
	routes, exist := slf.routes[topRoute]
	if !exist {
		routes = map[Route]map[Route]Handle{}
		slf.routes[topRoute] = routes
	}
	level2Routes, exist := routes[level2Route]
	if !exist {
		level2Routes = map[Route]Handle{}
		routes[level2Route] = level2Routes
	}
	_, exist = level2Routes[route]
	if exist {
		panic(fmt.Errorf("the route[%v:%v:%v] has already been registered, duplicate registration is not allowed", topRoute, level2Route, route))
	}
	level2Routes[route] = handleFunc
}

func (slf *Level3Router[Route, Handle]) Match(topRoute Route, level2Route Route, route Route) Handle {
	return slf.routes[topRoute][level2Route][route]
}
