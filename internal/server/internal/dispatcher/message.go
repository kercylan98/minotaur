package dispatcher

type Message[P comparable] interface {
	// GetProducer 获取消息生产者
	GetProducer() P
}
