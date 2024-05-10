package brokers

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"github.com/kercylan98/minotaur/toolkit/nexus"
)

func NewIndependentGoroutineOptions[I constraints.Ordered, T comparable]() *IndependentGoroutineOptions[I, T] {
	return &IndependentGoroutineOptions[I, T]{
		queueCreatedHook:            nil,
		queueClosedHook:             nil,
		queueBindCounterChangedHook: nil,
	}
}

type IndependentGoroutineOptions[I constraints.Ordered, T comparable] struct {
	queueCreatedHook            func(topic T, queue nexus.Queue[I, T], queueNum int) // 队列创建回调
	queueClosedHook             func(topic T, queue nexus.Queue[I, T], queueNum int) // 队列关闭回调
	queueBindCounterChangedHook func(topic T, count int)                             // 队列绑定计数变更回调
}

// Apply 应用配置
func (i *IndependentGoroutineOptions[I, T]) Apply(options ...*IndependentGoroutineOptions[I, T]) *IndependentGoroutineOptions[I, T] {
	for _, o := range options {
		i.queueCreatedHook = o.queueCreatedHook
		i.queueClosedHook = o.queueClosedHook
		i.queueBindCounterChangedHook = o.queueBindCounterChangedHook
	}

	return i
}

// WithQueueCreatedHook 设置队列创建回调
func (i *IndependentGoroutineOptions[I, T]) WithQueueCreatedHook(hook func(topic T, queue nexus.Queue[I, T], queueNum int)) *IndependentGoroutineOptions[I, T] {
	i.queueCreatedHook = hook
	return i
}

// WithQueueClosedHook 设置队列关闭回调
func (i *IndependentGoroutineOptions[I, T]) WithQueueClosedHook(hook func(topic T, queue nexus.Queue[I, T], queueNum int)) *IndependentGoroutineOptions[I, T] {
	i.queueClosedHook = hook
	return i
}

// WithQueueBindCounterChangedHook 设置队列绑定计数变更回调
func (i *IndependentGoroutineOptions[I, T]) WithQueueBindCounterChangedHook(hook func(topic T, count int)) *IndependentGoroutineOptions[I, T] {
	i.queueBindCounterChangedHook = hook
	return i
}
