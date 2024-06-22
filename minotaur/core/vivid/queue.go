package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/buffer"
	"github.com/kercylan98/minotaur/toolkit/pools"
	"sync"
)

var unboundedQueuePool = pools.NewObjectPool[unboundedQueue](func() *unboundedQueue {
	return &unboundedQueue{
		Ring: buffer.NewRing[Message](256),
		rw:   &sync.Mutex{},
	}
}, func(data *unboundedQueue) {
	data.Ring.Reset()
})

func releaseDefaultMailbox(mailbox *defaultMailbox) {
	if v, ok := mailbox.queue.(*unboundedQueue); ok {
		unboundedQueuePool.Put(v)
	}
	if v, ok := mailbox.systemQueue.(*unboundedQueue); ok {
		unboundedQueuePool.Put(v)
	}
}

type unboundedQueue struct {
	*buffer.Ring[Message]
	rw *sync.Mutex
}

func (q *unboundedQueue) Enqueue(message core.Message) {
	q.rw.Lock()
	defer q.rw.Unlock()
	q.Write(message)
}

func (q *unboundedQueue) Dequeue() core.Message {
	q.rw.Lock()
	m, _ := q.Read()
	q.rw.Unlock()
	return m
}
