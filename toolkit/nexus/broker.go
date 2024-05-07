package nexus

import "github.com/kercylan98/minotaur/toolkit/constraints"

// Broker 消息核心的接口定义
type Broker[I constraints.Ordered, T comparable] interface {

	// Run 运行消息核心
	Run()

	// Close 关闭消息核心
	Close()

	// Publish 发布消息
	Publish(topic T, event Event[I, T]) error
}
