package message

// Broker 消息核心的接口定义
type Broker[P Producer, Q Queue] interface {
	PublishMessage(message Message[P, Q])
}
