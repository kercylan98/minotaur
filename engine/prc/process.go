package prc

type UnboundProcess interface {
	// DeliveryUserMessage 投递用户消息，用户消息应该是优先级低于系统消息的消息，具体情况根据进程实现方式有所不同。
	//   - 通常情况下，该 sender 均是作为发送方的存在，不排除一些特殊情况，需要重定向发送方以确保消息响应时候被到目标进程，那么 forward 将表示重定向后的发送方
	//   - receiver 在大多数情况下应该都表示进程本身，在一些跨节点转发的情况下，可能会是其他进程
	DeliveryUserMessage(receiver, sender, forward *ProcessRef, message Message)

	// DeliverySystemMessage 投递具有最高优先级的系统消息，具体情况根据进程实现方式有所不同。
	//   - 通常情况下，该 sender 均是作为发送方的存在，不排除一些特殊情况，需要重定向发送方以确保消息响应时候被到目标进程，那么 forward 将表示重定向后的发送方
	//   - receiver 在大多数情况下应该都表示进程本身，在一些跨节点转发的情况下，可能会是其他进程
	DeliverySystemMessage(receiver, sender, forward *ProcessRef, message Message)
}

type Process interface {
	// Initialize 初始化进程
	Initialize(rc *ResourceController, id *ProcessId)

	UnboundProcess

	// IsTerminated 告知进程是否已经终止，如果已死，那么该进程的引用将会尝试将缓存更新为存活的进程
	IsTerminated() bool

	// Terminate 终止进程，通常情况下，该方法将会向进程投递一个包含发起方的终止消息，具体情况根据进程实现方式有所不同。
	//  - 该函数在进程从资源控制器中取消注册时将被调用
	Terminate(source *ProcessRef)
}
