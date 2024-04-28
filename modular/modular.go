package modular

import (
	"github.com/samber/do/v2"
	"reflect"
)

// Run 运行应用程序
func Run(application *Application) {
	application.run()
}

// RegisterService 注冊一个全局服务
//   - Instance 是服务的实例类型，该类型必须是一个实现了 GlobalService 接口的结构体指针类型，用于提供服务的具体实现
//   - Exposer 是服务的暴露接口，用于提供服务的对外接口
//
// 通过该函数注册的服务将会在应用程序启动时被实例化，并且在整个应用程序的生命周期中只存在一个实例
func RegisterService[Instance GlobalService, Exposer any](application *Application) {
	tof := reflect.TypeOf((*Instance)(nil)).Elem()
	vof := reflect.New(tof).Interface()
	service := vof.(GlobalService)
	exposer := vof.(Exposer)
	application.services = append(application.services, service)
	do.ProvideValue[Exposer](application.injector, exposer)
}

// InvokeService 获取特定全局服务的实例
func InvokeService[Expose any](application *Application) Expose {
	return do.MustInvoke[Expose](application.injector)
}
