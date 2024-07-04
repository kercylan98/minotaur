package mappings

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"sync"
)

var _ OrderInterface[int, int] = (*Order[int, int])(nil)

func NewOrderSync[K constraints.Hash, V any]() *OrderSync[K, V] {
	o := &OrderSync[K, V]{}
	return o
}

type OrderSync[K constraints.Hash, V any] struct {
	idx   map[K]int
	value []*orderEntry[K, V]
	rw    sync.RWMutex
}

// Get 获取指定 key 的元素
func (o *OrderSync[K, V]) Get(key K) (value V, exists bool) {
	o.rw.RLock()
	defer o.rw.RUnlock()
	idx, exists := o.idx[key]
	if !exists {
		return
	}
	return o.value[idx].V, true
}

// Add 添加元素
func (o *OrderSync[K, V]) Add(key K, value V) {
	o.rw.Lock()
	defer o.rw.Unlock()
	if _, exists := o.idx[key]; exists {
		return
	}
	if o.idx == nil {
		o.idx = make(map[K]int)
	}

	o.idx[key] = len(o.value)
	o.value = append(o.value, &orderEntry[K, V]{K: key, V: value})
}

// Set 设置指定 key 的元素
func (o *OrderSync[K, V]) Set(key K, value V) {
	o.rw.Lock()
	defer o.rw.Unlock()
	entry, exist := o.idx[key]
	if !exist {
		if o.idx == nil {
			o.idx = make(map[K]int)
		}
		o.idx[key] = len(o.value)
		o.value = append(o.value, &orderEntry[K, V]{K: key, V: value})
	} else {
		o.value[entry].V = value
	}
}

// Len 返回元素数量
func (o *OrderSync[K, V]) Len() int {
	o.rw.RLock()
	defer o.rw.RUnlock()
	return len(o.value)
}

// Del 删除指定 key 的元素
func (o *OrderSync[K, V]) Del(key K) {
	o.rw.Lock()
	defer o.rw.Unlock()
	idx, exists := o.idx[key]
	if !exists {
		return
	}

	if idx < len(o.value)-1 {
		last := o.value[len(o.value)-1]
		o.value[idx] = last
		o.idx[last.K] = idx
	}

	o.value = o.value[:len(o.value)-1]
	delete(o.idx, key)
}

// Range 遍历所有元素，如果 handle 返回 false，则停止遍历
func (o *OrderSync[K, V]) Range(handle func(key K, value V) bool) {
	o.rw.RLock()
	defer o.rw.RUnlock()
	for _, entry := range o.value {
		if !handle(entry.K, entry.V) {
			break
		}
	}
}
