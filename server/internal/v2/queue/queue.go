package queue

import (
	"github.com/kercylan98/minotaur/utils/buffer"
	"sync"
	"sync/atomic"
)

// New 创建一个并发安全的队列 Queue，该队列支持通过 chanSize 自定义读取 channel 的大小，同支持使用 bufferSize 指定 buffer.Ring 的初始大小
func New[Id, Q comparable, M Message[Q]](id Id, chanSize, bufferSize int) *Queue[Id, Q, M] {
	q := &Queue[Id, Q, M]{
		c:           make(chan MessageHandler[Id, Q, M], chanSize),
		buf:         buffer.NewRing[wrapper[Id, Q, M]](bufferSize),
		condRW:      &sync.RWMutex{},
		identifiers: make(map[Q]int64),
	}
	q.cond = sync.NewCond(q.condRW)
	q.state = &State[Id, Q, M]{
		queue:  q,
		id:     id,
		status: StatusNone,
	}
	return q
}

// Queue 队列是一个适用于消息处理等场景的并发安全的数据结构
//   - 该队列接收自定义的消息 M，并将消息有序的传入 Read 函数所返回的 channel 中以供处理
//   - 该结构主要实现目标为读写分离且并发安全的非阻塞传输队列，当消费阻塞时以牺牲内存为代价换取消息的生产不阻塞，适用于服务器消息处理等
//   - 该队列保证了消息的完整性，确保消息不丢失，在队列关闭后会等待所有消息处理完毕后进行关闭，并提供 SetClosedHandler 函数来监听队列的关闭信号
type Queue[Id, Q comparable, M Message[Q]] struct {
	state         *State[Id, Q, M]                // 队列状态信息
	c             chan MessageHandler[Id, Q, M]   // 消息读取通道
	buf           *buffer.Ring[wrapper[Id, Q, M]] // 消息缓冲区
	cond          *sync.Cond                      // 条件变量
	condRW        *sync.RWMutex                   // 条件变量的读写锁
	closedHandler func(q *Queue[Id, Q, M])        // 关闭处理函数
	identifiers   map[Q]int64                     // 标识符在队列的消息计数映射
}

// Id 获取队列 Id
func (q *Queue[Id, Q, M]) Id() Id {
	return q.state.id
}

// SetId 设置队列 Id
func (q *Queue[Id, Q, M]) SetId(id Id) {
	q.state.id = id
}

// SetClosedHandler 设置队列关闭处理函数，在队列关闭后将执行该函数。此刻队列不再可用
//   - Close 函数为非阻塞调用，调用后不会立即关闭队列，会等待消息处理完毕且处理期间不再有新消息介入
func (q *Queue[Id, Q, M]) SetClosedHandler(handler func(q *Queue[Id, Q, M])) {
	q.closedHandler = handler
}

// Run 阻塞的运行队列，当队列非首次运行时，将会引发来自 ErrorQueueInvalid 的 panic
func (q *Queue[Id, Q, M]) Run() {
	if atomic.LoadInt32(&q.state.status) != StatusNone {
		panic(ErrorQueueInvalid)
	}
	atomic.StoreInt32(&q.state.status, StatusRunning)
	defer func(q *Queue[Id, Q, M]) {
		if q.closedHandler != nil {
			q.closedHandler(q)
		}
	}(q)
	for {
		q.cond.L.Lock()
		for q.buf.IsEmpty() {
			if atomic.LoadInt32(&q.state.status) >= StatusClosing && q.state.total == 0 {
				q.cond.L.Unlock()
				atomic.StoreInt32(&q.state.status, StatusClosed)
				close(q.c)
				return
			}
			q.cond.Wait()
		}
		items := q.buf.ReadAll()
		q.cond.L.Unlock()
		for i := 0; i < len(items); i++ {
			item := items[i]
			q.c <- func(handler func(m M), finisher func(m M, last bool)) {
				defer func(msg M) {
					msg.OnProcessed(item.controller)

					queue := msg.GetQueue()

					q.cond.L.Lock()
					q.state.total--
					curr := q.identifiers[queue] - 1
					if curr != 0 {
						q.identifiers[queue] = curr
					} else {
						delete(q.identifiers, queue)
					}
					if finisher != nil {
						finisher(msg, curr == 0)
					}
					//log.Info("消息总计数", log.Int64("计数", q.state.total))
					q.cond.Signal()
					q.cond.L.Unlock()
				}(item.message)

				handler(item.message)
			}
		}
	}
}

// Push 向队列中推送来自 queue 的消息 m，当队列已关闭时将会返回 ErrorQueueClosed
func (q *Queue[Id, Q, M]) Push(queue Q, m M) error {
	if atomic.LoadInt32(&q.state.status) > StatusClosing {
		return ErrorQueueClosed
	}
	wrapper := newWrapper(q, m)

	q.cond.L.Lock()
	q.identifiers[queue]++
	q.state.total++
	q.buf.Write(wrapper)
	//log.Info("消息总计数", log.Int64("计数", q.state.total))
	q.cond.Signal()
	q.cond.L.Unlock()

	m.OnPublished(wrapper.controller)
	return nil
}

// WaitAdd 向队列增加来自外部的等待计数，在队列关闭时会等待该计数归零
func (q *Queue[Id, Q, M]) WaitAdd(queue Q, delta int64) {
	q.cond.L.Lock()

	currIdent := q.identifiers[queue]
	currIdent += delta
	q.identifiers[queue] = currIdent
	q.state.total += delta
	//log.Info("消息总计数", log.Int64("计数", q.state.total))

	q.cond.Signal()
	q.cond.L.Unlock()
}

// Read 获取队列消息的只读通道，
func (q *Queue[Id, Q, M]) Read() <-chan MessageHandler[Id, Q, M] {
	return q.c
}

// Close 关闭队列
func (q *Queue[Id, Q, M]) Close() {
	atomic.CompareAndSwapInt32(&q.state.status, StatusRunning, StatusClosing)
	q.cond.Broadcast()
}

// State 获取队列状态
func (q *Queue[Id, Q, M]) State() *State[Id, Q, M] {
	return q.state
}

// GetMessageCount 获取消息数量
func (q *Queue[Id, Q, M]) GetMessageCount() (count int64) {
	q.condRW.RLock()
	defer q.condRW.RUnlock()
	for _, i := range q.identifiers {
		count += i
	}
	return
}

// GetMessageCountWithIdent 获取特定消息人的消息数量
func (q *Queue[Id, Q, M]) GetMessageCountWithIdent(queue Q) int64 {
	q.condRW.RLock()
	defer q.condRW.RUnlock()
	return q.identifiers[queue]
}
