package nexus

// Broker 消息核心的接口定义
type Broker[I, T comparable] interface {
	Run()
	Close()
	Publish(topic T, event Event[I, T]) error
}
