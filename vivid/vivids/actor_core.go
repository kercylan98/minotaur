package vivids

// ActorCore 是 Actor 的核心接口定义，被用于高级功能的实现
type ActorCore interface {
	ActorRef
	ActorContext

	// GetOptions 获取 Actor 的配置项
	GetOptions() *ActorOptions

	// SetContext 为 Actor 设置一个上下文
	SetContext(key, ctx any)

	// GetContext 获取 Actor 的上下文
	GetContext(key any) any

	// IsPause 判断 Actor 当前是否处于暂停处理消息的状态，返回一个只读的通道
	//  - 当 Actor 未处于暂停状态时，返回 nil
	IsPause() <-chan struct{}

	// BindMessageActorContext 绑定消息的 Actor 上下文为当前 Actor
	BindMessageActorContext(ctx MessageContext)

	// OnReceived 用于处理接收到的消息
	OnReceived(ctx MessageContext) error
}
