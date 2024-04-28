package server

import (
	"github.com/kercylan98/minotaur/utils/log"
	"reflect"
)

// MultipleService 兼容传统 service 设计模式的接口，通过该接口可以实现更简洁、更具有可读性的服务绑定
type MultipleService interface {
	// OnInit 初始化服务，该方法将会在 Server 初始化时执行
	//   - 通常来说，该阶段发生任何错误都应该 panic 以阻止 Server 启动
	OnInit(srv *MultipleServer)
	// OnPreloading 预加载阶段，该方法将会在所有服务的 OnInit 函数执行完毕后执行
	//   - 通常来说，该阶段发生任何错误都应该 panic 以阻止 Server 启动
	OnPreloading(srv *MultipleServer)
}

// BindServiceToMultipleServer 绑定服务到多个 MultipleServer
func BindServiceToMultipleServer(server *MultipleServer, services ...MultipleService) {
	for i := 0; i < len(services); i++ {
		service := services[i]
		server.preload = append(server.preload, func() {
			name := reflect.TypeOf(service).String()
			defer func(name string) {
				if err := recover(); err != nil {
					log.Error("MultipleServer", log.String("service", name), log.String("status", "preloading"), log.Any("err", err))
					panic(err)
				}
			}(name)
			service.OnPreloading(server)
			log.Info("MultipleServer", log.String("service", name), log.String("status", "preloaded"))
		})
	}

	for i := 0; i < len(services); i++ {
		service := services[i]
		server.services = append(server.services, func() {
			name := reflect.TypeOf(service).String()
			defer func(name string) {
				if err := recover(); err != nil {
					log.Error("MultipleServer", log.String("service", name), log.String("status", "initialization"), log.Any("err", err))
					panic(err)
				}
			}(name)
			service.OnInit(server)
			log.Info("MultipleServer", log.String("service", name), log.String("status", "initialized"))
		})
	}

}
