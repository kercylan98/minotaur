package reactor

import (
	"errors"
	"github.com/kercylan98/minotaur/utils/buffer"
	"sync"
	"sync/atomic"
)

func newQueue[M any](idx, chanSize, bufferSize int) *queue[M] {
	q := &queue[M]{
		c:    make(chan queueMessage[M], chanSize),
		buf:  buffer.NewRing[queueMessage[M]](bufferSize),
		cond: sync.NewCond(&sync.Mutex{}),
	}
	q.QueueState = &QueueState[M]{
		queue:  q,
		idx:    idx,
		status: QueueStatusNone,
	}
	return q
}

type queue[M any] struct {
	*QueueState[M]
	c             chan queueMessage[M]          // 通道
	buf           *buffer.Ring[queueMessage[M]] // 缓冲区
	cond          *sync.Cond                    // 条件变量
	closedHandler func(q *queue[M])             // 关闭处理函数
}

func (q *queue[M]) Id() int {
	return q.idx
}

func (q *queue[M]) setClosedHandler(handler func(q *queue[M])) {
	q.closedHandler = handler
}

func (q *queue[M]) run() {
	atomic.StoreInt32(&q.status, QueueStatusRunning)
	defer func(q *queue[M]) {
		atomic.StoreInt32(&q.status, QueueStatusClosed)
		if q.closedHandler != nil {
			q.closedHandler(q)
		}
	}(q)
	for {
		q.cond.L.Lock()
		for q.buf.IsEmpty() {
			if atomic.LoadInt32(&q.status) >= QueueStatusClosing {
				q.cond.L.Unlock()
				close(q.c)
				return
			}
			q.cond.Wait()
		}
		items := q.buf.ReadAll()
		q.cond.L.Unlock()
		for _, item := range items {
			q.c <- item
		}
	}
}

func (q *queue[M]) push(ident *identifiable, m M) error {
	if atomic.LoadInt32(&q.status) > QueueStatusRunning {
		return errors.New("queue status exception")
	}
	q.cond.L.Lock()
	q.buf.Write(queueMessage[M]{
		ident: ident,
		msg:   m,
	})
	q.cond.Signal()
	q.cond.L.Unlock()
	return nil
}

func (q *queue[M]) read() <-chan queueMessage[M] {
	return q.c
}
