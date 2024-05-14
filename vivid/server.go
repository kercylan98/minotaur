package vivid

// Server 是 ActorSystem 远程调用的服务端
type Server interface {
	// Run 用于启动一个 ActorSystem 服务端
	Run() error

	// Shutdown 用于关闭一个 ActorSystem 服务端
	Shutdown() error

	// C 用于获取一个用于接收远程消息的通道
	C() <-chan RemoteMessageEvent
}
