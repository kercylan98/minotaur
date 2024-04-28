package modular

type Running interface {
	// OnRunning 开始运行阶段，通常在该阶段处理一些依赖于 OnBlock 阶段的逻辑
	//  - 该阶段不应该阻塞进程，并且发生错误应该通过 panic 阻止服务继续运行
	OnRunning()
}
