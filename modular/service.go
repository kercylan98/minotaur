package modular

import "reflect"

// Service 模块化服务接口，所有的服务均需要实现该接口，在服务的生命周期内发生任何错误均应通过 panic 阻止服务继续运行
//   - 生命周期示例： OnInit -> OnPreload -> OnMount
//
// 在 Golang 中，包与包之间互相引用会导致循环依赖，因此在模块化应用程序中，所有的服务均不应该直接引用其他服务。
//
// 服务应该在 OnInit 阶段将不依赖其他服务的内容初始化完成，并且如果服务需要暴露给其他服务调用，那么也应该在 OnInit 阶段完成对外暴露。
//   - 暴露方式可参考 modular/example
//
// 在 OnPreload 阶段，服务应该完成对其依赖服务的依赖注入，最终在 OnMount 阶段完成对服务功能的定义、路由的声明等。
type Service interface {
	// OnInit 服务初始化阶段，该阶段不应该依赖其他任何服务
	OnInit()

	// OnPreload 预加载阶段，在进入该阶段时，所有服务已经初始化完成，可在该阶段注入其他服务的依赖
	OnPreload()

	// OnMount 挂载阶段，该阶段所有服务本身及依赖的服务都已经初始化完成，可在该阶段进行服务功能的定义
	OnMount()
}

// RegisterServices 注册服务
func RegisterServices(s ...Service) {
	application.RegisterServices(s...)
}

func newService(instance Service) *service {
	vof := reflect.ValueOf(instance)
	return &service{
		name:     vof.Type().String(),
		instance: instance,
		vof:      vof,
	}
}

type service struct {
	name     string
	instance Service
	vof      reflect.Value
}
