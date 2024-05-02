package modular

// BasicService 模块化基本服务接口，所有的服务均需要实现该接口，在服务的生命周期内发生任何错误均应通过 panic 阻止服务继续运行
//
// 在 Golang 中，包与包之间互相引用会导致循环依赖，因此在模块化应用程序中，所有的服务均不应该直接引用其他服务。
//
// 服务应该在 OnInit 阶段将不依赖其他服务的内容初始化完成，并且如果服务需要暴露给其他服务调用，那么也应该在 OnInit 阶段完成对外暴露。
//   - 暴露方式可参考 modular/example
//
// 在 OnPreload 阶段，服务应该完成对其依赖服务的依赖注入，最终在 OnMount 阶段完成对服务功能的定义、路由的声明等。
type BasicService interface {
	// OnInit 服务初始化阶段，该阶段不应该依赖其他任何服务
	OnInit(application *Application)

	// OnPreload 预加载阶段，在进入该阶段时，所有服务已经初始化完成，通常在该阶段注入依赖的服务
	OnPreload(application *Application)

	// OnMount 挂载阶段，该阶段所有服务本身及依赖的服务都已经初始化完成，通常在该阶段进行 OnInit 无法完成的初始化工作
	//
	// 该阶段的意义适用于如下场景：
	//   - 数据库服务依赖于配置服务，由于数据库服务需要在配置服务初始化完成后才能初始化，因此数据库服务的初始化工作无法在 OnInit 阶段完成
	OnMount(application *Application)

	// OnStart 启动阶段，该阶段所有服务均已准备就绪，可在该阶段进行服务功能的定义
	OnStart(application *Application)
}

// BlockService 标识模块化服务为阻塞进程的服务，当实现了 BasicService 且实现了 BlockService 接口时，模块化应用程序会在 BasicService.OnStart 阶段完成后执行 OnBlock 函数
//
// 该接口适用于 Http 服务、WebSocket 服务等需要阻塞进程的服务。需要注意的是， OnBlock 的执行不能保证按照 BasicService 的注册顺序执行
type BlockService interface {
	BasicService
	// OnBlock 阻塞进程逻辑
	OnBlock(application *Application)
}
