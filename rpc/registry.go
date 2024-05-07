package rpc

// Registry 表示一个 RPC 服务注册器，用于注册服务到注册中心
type Registry interface {
	// OnRegister 服务注册时触发
	OnRegister(service Service) error

	// OnUnregister 服务注销时触发
	OnUnregister() error

	// WatchCall 监听服务调用事件
	WatchCall() <-chan Caller
}
