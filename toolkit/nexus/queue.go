package nexus

import (
	"github.com/kercylan98/minotaur/toolkit/balancer"
	"github.com/kercylan98/minotaur/toolkit/constraints"
)

type Queue[I constraints.Ordered, T comparable] interface {
	// GetId 获取队列 Id
	GetId() I
	// Publish 向队列中推送消息
	Publish(topic T, event Event[I, T]) error
	// IncrementCustomMessageCount 增加自定义消息计数
	IncrementCustomMessageCount(topic T, delta int64)
	// Run 运行队列
	Run()
	// Consume 消费消息
	Consume() <-chan EventInfo[I, T]
	// Close 关闭队列
	Close()

	balancer.Item[I]
}
