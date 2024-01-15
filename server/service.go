package server

import (
	"github.com/kercylan98/minotaur/utils/log"
	"reflect"
)

// Service 兼容传统 service 设计模式的接口，通过该接口可以实现更简洁、更具有可读性的服务绑定
//   - 在这之前，我们在实现功能上会将 Server 进行全局存储，之后通过 init 函数进行初始化，这样的顺序是不可控的。
type Service interface {
	// OnInit 初始化服务，该方法将会在 Server 初始化时执行
	//   - 通常来说，该阶段发生任何错误都应该 panic 以阻止 Server 启动
	OnInit(srv *Server)
}

// BindService 绑定服务到特定 Server，被绑定的服务将会在 Server 初始化时执行 Service.OnInit 方法
func BindService(srv *Server, services ...Service) {
	for i := 0; i < len(services); i++ {
		service := services[i]
		srv.services = append(srv.services, func() {
			name := reflect.TypeOf(service).String()
			defer func(name string) {
				if err := recover(); err != nil {
					log.Error("Server", log.String("service", name), log.String("status", "initialization"), log.Any("err", err))
					panic(err)
				}
			}(name)
			service.OnInit(srv)
			log.Info("Server", log.String("service", name), log.String("status", "initialized"))
		})
	}
}
