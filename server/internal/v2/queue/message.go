package queue

// Message 消息接口定义
type Message[Queue comparable] interface {
	// GetQueue 获取消息执行队列
	GetQueue() Queue
	// OnPublished 消息发布成功
	OnPublished(controller Controller)
	// OnProcessed 消息处理完成
	OnProcessed(controller Controller)
}
