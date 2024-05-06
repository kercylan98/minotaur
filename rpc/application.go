package rpc

import (
	"context"
	"github.com/kercylan98/minotaur/server/router"
	"github.com/kercylan98/minotaur/toolkit"
)

// NewApplication 创建一个新的 RPC 应用程序
func NewApplication(service Service) *Application {
	var a = &Application{
		service: service,
		router:  router.NewTypeFixedMultistage[Route, any](),
	}
	a.ctx, a.cancel = context.WithCancel(context.Background())

	return a
}

// Application 表示一个 RPC 应用程序
type Application struct {
	ctx      context.Context                         // 应用程序上下文
	cancel   context.CancelFunc                      // 取消函数
	service  Service                                 // 在应用程序中注册的服务信息
	router   *router.TypeFixedMultistage[Route, any] // 路由器
	registry Registry                                // 服务注册器
}

// Register 向应用程序注册给定的路由
func (a *Application) Register(route ...Route) *RouteBinder {
	return &RouteBinder{application: a, route: route}
}

// Run 运行 RPC 应用程序，开始监听服务
func (a *Application) Run() error {
	defer a.cancel()
	go a.onCall()
	return a.registry.OnRegister(a.service)
}

// Close 关闭 RPC 应用程序
func (a *Application) Close() error {
	a.cancel()
	return a.registry.OnUnregister()
}

func (a *Application) onCall() {
	var ch = a.registry.WatchCall()
	for {
		select {
		case <-a.ctx.Done():
			return
		case caller := <-ch:
			handler := a.router.Match(caller.GetRoute()...)
			if handler == nil {
				continue
			}

			switch f := handler.(type) {
			case UnaryHandler:
				resp := f(func(dst any) {
					toolkit.UnmarshalJSON(caller.GetPacket(), dst)
				})
				_ = caller.Respond(toolkit.MarshalJSON(resp))
			case UnaryNotifyHandler:
				f(func(dst any) {
					toolkit.UnmarshalJSON(caller.GetPacket(), dst)
				})
			default:
				continue
			}
		}
	}
}
