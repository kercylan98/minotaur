package vivids

// Dispatcher 消息调度器接口
type Dispatcher interface {
	// OnInit 用于初始化调度器
	OnInit(system ActorSystemExternal)

	// Attach 用于将一个 Actor 添加到调度器中
	Attach(actor ActorCore) error

	// Detach 用于将一个 Actor 从调度器中移除
	Detach(actor ActorCore) error

	// Send 用于向一个 Actor 发送消息
	Send(receiver ActorCore, msg MessageContext) error

	// Stop 用于停止调度器
	Stop()
}
