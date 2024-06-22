package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/buffer"
	"github.com/kercylan98/minotaur/toolkit/pools"
)

var unboundedQueuePool = pools.NewObjectPool[unboundedQueue](func() *unboundedQueue {
	return &unboundedQueue{
		Ring: buffer.NewRing[Message](1024),
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
}

func (q *unboundedQueue) Enqueue(message core.Message) {
	q.Write(message)
}

func (q *unboundedQueue) Dequeue() core.Message {
	m, _ := q.Read()
	return m
}
