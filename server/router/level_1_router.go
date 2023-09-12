package router

import (
	"fmt"
	"reflect"
)

// NewLevel1Router 创建支持一级分类的路由器
//
// Deprecated: 从 Minotaur 0.1.7 开始，由于该路由器设计不合理，局限性大，已弃用。建议使用 Multistage 进行代替。
func NewLevel1Router[Route comparable, Handle any]() *Level1Router[Route, Handle] {
	return &Level1Router[Route, Handle]{
		routes: map[Route]Handle{},
	}
}

// Level1Router 支持一级分类的路由器
//
// Deprecated: 从 Minotaur 0.1.7 开始，由于该路由器设计不合理，局限性大，已弃用。建议使用 Multistage 进行代替。
type Level1Router[Route comparable, Handle any] struct {
	routes map[Route]Handle
}

// Route 创建路由
//
// Deprecated: 从 Minotaur 0.1.7 开始，由于该路由器设计不合理，局限性大，已弃用。建议使用 Multistage 进行代替。
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

// Match 匹配路由
//
// Deprecated: 从 Minotaur 0.1.7 开始，由于该路由器设计不合理，局限性大，已弃用。建议使用 Multistage 进行代替。
func (slf *Level1Router[Route, Handle]) Match(route Route) Handle {
	return slf.routes[route]
}
