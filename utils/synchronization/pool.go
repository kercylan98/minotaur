package synchronization

import "sync"

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
	return slf.generator()
}

func (slf *Pool[T]) Release(data T) {
	slf.releaser(data)
	slf.put(data)
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
