package synchronization

import (
	"go.uber.org/zap"
	"minotaur/utils/log"
	"sync"
)

func NewPool[T any](bufferSize int, generator func() T, releaser func(data T)) *Pool[T] {
	pool := &Pool[T]{
		bufferSize: bufferSize,
		generator:  generator,
		releaser:   releaser,
	}
	for i := 0; i < bufferSize; i++ {
		pool.put(generator())
	}
	return pool
}

// Pool 线程安全的对象缓冲池
//   - 一些高频临时生成使用的对象可以通过 Pool 进行管理，例如属性计算等
//   - 缓冲区内存在可用对象时直接返回，否则新建一个进行返回
//   - 通过 Release 将使用完成的对象放回缓冲区，超出缓冲区大小的对象将被放弃
type Pool[T any] struct {
	mutex      sync.Mutex
	buffers    []T
	bufferSize int
	generator  func() T
	releaser   func(data T)
}

func (slf *Pool[T]) Get() T {
	slf.mutex.Lock()
	if len(slf.buffers) > 0 {
		data := slf.buffers[0]
		slf.buffers = slf.buffers[1:]
		slf.mutex.Unlock()
		return data
	}
	slf.mutex.Unlock()
	log.Warn("Pool", zap.String("Get", "the number of buffer members is insufficient, consider whether it is due to unreleased or inappropriate buffer size"))
	return slf.generator()
}

func (slf *Pool[T]) Release(data T) {
	slf.releaser(data)
	slf.put(data)
}

func (slf *Pool[T]) Close() {
	slf.mutex.Lock()
	slf.buffers = nil
	slf.bufferSize = 0
	slf.generator = nil
	slf.releaser = nil
	slf.mutex.Unlock()
}

func (slf *Pool[T]) put(data T) {
	slf.mutex.Lock()
	if len(slf.buffers) > slf.bufferSize {
		slf.mutex.Unlock()
		return
	}
	slf.buffers = append(slf.buffers, data)
	slf.mutex.Unlock()
}
