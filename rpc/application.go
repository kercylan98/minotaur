package rpc

// NewApplication 创建一个 RPC 服务应用程序
func NewApplication(core Core) *Application {
	return &Application{core: core}
}

// Application RPC 服务应用程序
type Application struct {
	core     Core
	services []Service
}

// Register 注册 RPC 服务
func (a *Application) Register(services ...Service) *Application {
	a.services = append(a.services, services...)
	return a
}

// Run 运行 RPC 服务应用程序
func (a *Application) Run(info ServiceInfo) error {
	rpcRouter := new(router).init(a)

	// 绑定路由
	for _, service := range a.services {
		service.OnRPCSetup(rpcRouter)
	}

	// 初始化 RPC 核心
	if err := a.core.OnInit(info, rpcRouter, rpcRouter.GetRoutes()); err != nil {
		return err
	}
	return nil
}
