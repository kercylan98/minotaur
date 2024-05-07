package rpc

// Discovery 表示一个 RPC 服务发现器
type Discovery interface {
	// WatchRegister 监听服务注册事件
	WatchRegister() <-chan CallableService

	// WatchUnregister 监听服务注销事件
	WatchUnregister() <-chan CallableService

	// Close 关闭服务发现器
	Close() error
}
