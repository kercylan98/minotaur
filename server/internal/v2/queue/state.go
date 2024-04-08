package queue

import (
	"sync/atomic"
)

const (
	StatusNone    = iota - 1 // 队列未运行
	StatusRunning            // 队列运行中
	StatusClosing            // 队列关闭中
	StatusClosed             // 队列已关闭
)

type State[Id, Q comparable, M Message[Q]] struct {
	queue  *Queue[Id, Q, M]
	id     Id    // 队列 ID
	status int32 // 状态标志
	total  int64 // 消息总计数
}

// IsClosed 判断队列是否已关闭
func (q *State[Id, Q, M]) IsClosed() bool {
	return atomic.LoadInt32(&q.status) == StatusClosed
}

// IsClosing 判断队列是否正在关闭
func (q *State[Id, Q, M]) IsClosing() bool {
	return atomic.LoadInt32(&q.status) == StatusClosing
}

// IsRunning 判断队列是否正在运行
func (q *State[Id, Q, M]) IsRunning() bool {
	return atomic.LoadInt32(&q.status) == StatusRunning
}
