package reactor

import (
	"errors"
	"github.com/kercylan98/minotaur/utils/buffer"
	"sync"
	"sync/atomic"
)

func newQueue[M any](idx, chanSize, bufferSize int) *queue[M] {
	q := &queue[M]{
		c:   make(chan M, chanSize),
		buf: buffer.NewRing[M](bufferSize),
		rw:  sync.NewCond(&sync.Mutex{}),
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
	c   chan M          // 通道
	buf *buffer.Ring[M] // 缓冲区
	rw  *sync.Cond      // 读写锁
}

func (q *queue[M]) Id() int {
	return q.idx
}

func (q *queue[M]) run() {
	atomic.StoreInt32(&q.status, QueueStatusRunning)
	defer func(q *queue[M]) {
		atomic.StoreInt32(&q.status, QueueStatusClosed)
	}(q)
	for {
		q.rw.L.Lock()
		for q.buf.IsEmpty() {
			if atomic.LoadInt32(&q.status) >= QueueStatusClosing {
				q.rw.L.Unlock()
				close(q.c)
				return
			}
			q.rw.Wait()
		}
		items := q.buf.ReadAll()
		q.rw.L.Unlock()
		for _, item := range items {
			q.c <- item
		}
	}
}

func (q *queue[M]) push(m M) error {
	if atomic.LoadInt32(&q.status) != QueueStatusRunning {
		return errors.New("queue status exception")
	}
	q.rw.L.Lock()
	q.buf.Write(m)
	q.rw.Signal()
	q.rw.L.Unlock()
	return nil
}

func (q *queue[M]) read() <-chan M {
	return q.c
}
