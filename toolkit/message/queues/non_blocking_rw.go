package queues

import (
	"errors"
	"github.com/kercylan98/minotaur/toolkit/message"
	"github.com/kercylan98/minotaur/utils/buffer"
	"sync"
	"sync/atomic"
	"time"
)

const (
	NonBlockingRWStatusNone    NonBlockingRWState = iota - 1 // 队列未运行
	NonBlockingRWStatusRunning                               // 队列运行中
	NonBlockingRWStatusClosing                               // 队列关闭中
	NonBlockingRWStatusClosed                                // 队列已关闭
)

var (
	ErrorQueueClosed  = errors.New("queue closed")  // 队列已关闭
	ErrorQueueInvalid = errors.New("queue invalid") // 无效的队列
)

type NonBlockingRWState = int32

type nonBlockingRWEventInfo[I, T comparable] struct {
	topic T
	event message.Event[I, T]
	exec  func(handler message.EventHandler[T], finisher message.EventFinisher[I, T])
}

func (e *nonBlockingRWEventInfo[I, T]) GetTopic() T {
	return e.topic
}

func (e *nonBlockingRWEventInfo[I, T]) Exec(handler message.EventHandler[T], finisher message.EventFinisher[I, T]) {
	e.exec(handler, finisher)
}

// NewNonBlockingRW 创建一个并发安全的队列 NonBlockingRW，该队列支持通过 chanSize 自定义读取 channel 的大小，同支持使用 bufferSize 指定 buffer.Ring 的初始大小
//   - closedHandler 可选的设置队列关闭处理函数，在队列关闭后将执行该函数。此刻队列不再可用
func NewNonBlockingRW[I, T comparable](id I, chanSize, bufferSize int) message.Queue[I, T] {
	q := &NonBlockingRW[I, T]{
		id:     id,
		status: NonBlockingRWStatusNone,
		c:      make(chan message.EventInfo[I, T], chanSize),
		buf:    buffer.NewRing[nonBlockingRWEventInfo[I, T]](bufferSize),
		condRW: &sync.RWMutex{},
		topics: make(map[T]int64),
	}
	q.cond = sync.NewCond(q.condRW)
	return q
}

// NonBlockingRW 队列是一个适用于消息处理等场景的并发安全的数据结构
//   - 该队列接收自定义的消息 M，并将消息有序的传入 Read 函数所返回的 channel 中以供处理
//   - 该结构主要实现目标为读写分离且并发安全的非阻塞传输队列，当消费阻塞时以牺牲内存为代价换取消息的生产不阻塞，适用于服务器消息处理等
//   - 该队列保证了消息的完整性，确保消息不丢失，在队列关闭后会等待所有消息处理完毕后进行关闭，并提供 SetClosedHandler 函数来监听队列的关闭信号
type NonBlockingRW[I, T comparable] struct {
	id     I                                          // 队列 ID
	status int32                                      // 状态标志
	total  int64                                      // 消息总计数
	topics map[T]int64                                // 主题对应的消息计数映射
	buf    *buffer.Ring[nonBlockingRWEventInfo[I, T]] // 消息缓冲区
	c      chan message.EventInfo[I, T]               // 消息读取通道
	cs     chan struct{}                              // 关闭信号
	cond   *sync.Cond                                 // 条件变量
	condRW *sync.RWMutex                              // 条件变量的读写锁
}

// GetId 获取队列 Id
func (n *NonBlockingRW[I, T]) GetId() I {
	return n.id
}

// Run 阻塞的运行队列，当队列非首次运行时，将会引发来自 ErrorQueueInvalid 的 panic
func (n *NonBlockingRW[I, T]) Run() {
	if atomic.LoadInt32(&n.status) != NonBlockingRWStatusNone {
		panic(ErrorQueueInvalid)
	}
	atomic.StoreInt32(&n.status, NonBlockingRWStatusRunning)
	for {
		n.cond.L.Lock()
		for n.buf.IsEmpty() {
			if atomic.LoadInt32(&n.status) >= NonBlockingRWStatusClosing && n.total == 0 {
				n.cond.L.Unlock()
				atomic.StoreInt32(&n.status, NonBlockingRWStatusClosed)
				close(n.c)
				close(n.cs)
				return
			}
			n.cond.Wait()
		}
		items := n.buf.ReadAll()
		n.cond.L.Unlock()
		for i := 0; i < len(items); i++ {
			ei := &items[i]
			n.c <- ei
		}
	}
}

// Consume 获取队列消息的只读通道，
func (n *NonBlockingRW[I, T]) Consume() <-chan message.EventInfo[I, T] {
	return n.c
}

// Close 关闭队列
func (n *NonBlockingRW[I, T]) Close() {
	if atomic.CompareAndSwapInt32(&n.status, NonBlockingRWStatusRunning, NonBlockingRWStatusClosing) {
		n.cs = make(chan struct{})
		n.cond.Broadcast()
		<-n.cs
	}
}

// GetMessageCount 获取消息数量
func (n *NonBlockingRW[I, T]) GetMessageCount() (count int64) {
	n.condRW.RLock()
	defer n.condRW.RUnlock()
	return n.total
}

// GetTopicMessageCount 获取特定主题的消息数量
func (n *NonBlockingRW[I, T]) GetTopicMessageCount(topic T) int64 {
	n.condRW.RLock()
	defer n.condRW.RUnlock()
	return n.topics[topic]
}

func (n *NonBlockingRW[I, T]) Publish(topic T, event message.Event[I, T]) error {
	if atomic.LoadInt32(&n.status) > NonBlockingRWStatusClosing {
		return ErrorQueueClosed
	}

	ei := nonBlockingRWEventInfo[I, T]{
		topic: topic,
		event: event,
		exec: func(handler message.EventHandler[T], finisher message.EventFinisher[I, T]) {
			defer func() {
				event.OnProcessed(topic, n, time.Now())

				n.cond.L.Lock()
				n.total--
				curr := n.topics[topic] - 1
				if curr != 0 {
					n.topics[topic] = curr
				} else {
					delete(n.topics, topic)
				}
				if finisher != nil {
					finisher(topic, curr == 0)
				}
				n.cond.Signal()
				n.cond.L.Unlock()
			}()

			handler(topic, func() {
				event.OnProcess(topic, n, time.Now())
			})
			return
		},
	}

	n.cond.L.Lock()
	n.topics[topic]++
	n.total++
	n.buf.Write(ei)
	//log.Info("消息总计数", log.Int64("计数", q.state.total))
	n.cond.Signal()
	n.cond.L.Unlock()

	return nil
}

func (n *NonBlockingRW[I, T]) IncrementCustomMessageCount(topic T, delta int64) {
	n.cond.L.Lock()
	n.total += delta
	n.topics[topic] += delta
	n.cond.Signal()
	n.cond.L.Unlock()
}

// IsClosed 判断队列是否已关闭
func (n *NonBlockingRW[I, T]) IsClosed() bool {
	return atomic.LoadInt32(&n.status) == NonBlockingRWStatusClosed
}

// IsClosing 判断队列是否正在关闭
func (n *NonBlockingRW[I, T]) IsClosing() bool {
	return atomic.LoadInt32(&n.status) == NonBlockingRWStatusClosing
}

// IsRunning 判断队列是否正在运行
func (n *NonBlockingRW[I, T]) IsRunning() bool {
	return atomic.LoadInt32(&n.status) == NonBlockingRWStatusRunning
}
