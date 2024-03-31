package reactor

import (
	"sync/atomic"
)

const (
	QueueStatusNone    = iota - 1 // 队列未运行
	QueueStatusRunning            // 队列运行中
	QueueStatusClosing            // 队列关闭中
	QueueStatusClosed             // 队列已关闭
)

type QueueState[M any] struct {
	queue  *queue[M]
	idx    int   // 队列索引
	status int32 // 状态标志
}

// IsClosed 判断队列是否已关闭
func (q *QueueState[M]) IsClosed() bool {
	return atomic.LoadInt32(&q.status) == QueueStatusClosed
}

// IsClosing 判断队列是否正在关闭
func (q *QueueState[M]) IsClosing() bool {
	return atomic.LoadInt32(&q.status) == QueueStatusClosing
}

// IsRunning 判断队列是否正在运行
func (q *QueueState[M]) IsRunning() bool {
	return atomic.LoadInt32(&q.status) == QueueStatusRunning
}

// Close 关闭队列
func (q *QueueState[M]) Close() {
	atomic.CompareAndSwapInt32(&q.status, QueueStatusRunning, QueueStatusClosing)
}
