package listings

import (
	"github.com/kercylan98/minotaur/utils/collection"
	"sync"
)

// NewSyncSlice 创建一个 SyncSlice
func NewSyncSlice[V any](length, cap int) *SyncSlice[V] {
	s := &SyncSlice[V]{}
	if length > 0 || cap > 0 {
		s.data = make([]V, length, cap)
	}
	return s
}

// SyncSlice 是基于 sync.RWMutex 实现的线程安全的 slice
type SyncSlice[V any] struct {
	rw   sync.RWMutex
	data []V
}

func (slf *SyncSlice[V]) Get(index int) V {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	return slf.data[index]
}

func (slf *SyncSlice[V]) GetWithRange(start, end int) []V {
	return slf.data[start:end]
}

func (slf *SyncSlice[V]) Set(index int, value V) {
	slf.rw.Lock()
	slf.data[index] = value
	slf.rw.Unlock()
}

func (slf *SyncSlice[V]) Append(values ...V) {
	slf.rw.Lock()
	slf.data = append(slf.data, values...)
	slf.rw.Unlock()
}

func (slf *SyncSlice[V]) Release() {
	slf.rw.Lock()
	slf.data = nil
	slf.rw.Unlock()
}

func (slf *SyncSlice[V]) Clear() {
	slf.rw.Lock()
	slf.data = slf.data[:0]
	slf.rw.Unlock()
}

func (slf *SyncSlice[V]) GetData() []V {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	return collection.CloneSlice(slf.data)
}
