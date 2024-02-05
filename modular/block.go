package modular

// Block 标识模块化服务为阻塞进程的服务，当实现了 Service 且实现了 Block 接口时，模块化应用程序会在 Service.OnMount 阶段完成后执行 OnBlock 函数
//
// 该接口适用于 Http 服务、WebSocket 服务等需要阻塞进程的服务。需要注意的是， OnBlock 的执行不能保证按照 Service 的注册顺序执行
type Block interface {
	Service
	// OnBlock 阻塞进程
	OnBlock()
}
