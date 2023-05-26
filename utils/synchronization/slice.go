package synchronization

import (
	"github.com/kercylan98/minotaur/utils/slice"
	"sync"
)

func NewSlice[T any](options ...SliceOption[T]) *Slice[T] {
	s := &Slice[T]{}
	for _, option := range options {
		option(s)
	}
	return s
}

type Slice[T any] struct {
	rw   sync.RWMutex
	data []T
}

func (slf *Slice[T]) Get(index int) T {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	return slf.data[index]
}

func (slf *Slice[T]) GetWithRange(start, end int) []T {
	return slf.data[start:end]
}

func (slf *Slice[T]) Set(index int, value T) {
	slf.rw.Lock()
	slf.data[index] = value
	slf.rw.Unlock()
}

func (slf *Slice[T]) Append(values ...T) {
	slf.rw.Lock()
	slf.data = append(slf.data, values...)
	slf.rw.Unlock()
}

func (slf *Slice[T]) Release() {
	slf.rw.Lock()
	slf.data = nil
	slf.rw.Unlock()
}

func (slf *Slice[T]) Clear() {
	slf.rw.Lock()
	slf.data = slf.data[:0]
	slf.rw.Unlock()
}

func (slf *Slice[T]) GetData() []T {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	return slice.Copy(slf.data)
}
