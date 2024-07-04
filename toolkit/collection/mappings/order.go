package mappings

import "github.com/kercylan98/minotaur/toolkit/constraints"

var _ OrderInterface[int, int] = (*Order[int, int])(nil)

func NewOrder[K constraints.Hash, V any]() *Order[K, V] {
	o := &Order[K, V]{}
	return o
}

type orderEntry[K constraints.Hash, V any] struct {
	K K
	V V
}

type Order[K constraints.Hash, V any] struct {
	idx   map[K]int
	value []*orderEntry[K, V]
}

// Get 获取指定 key 的元素
func (o *Order[K, V]) Get(key K) (value V, exists bool) {
	idx, exists := o.idx[key]
	if !exists {
		return
	}
	return o.value[idx].V, true
}

// Add 添加元素
func (o *Order[K, V]) Add(key K, value V) {
	if _, exists := o.idx[key]; exists {
		return
	}
	if o.idx == nil {
		o.idx = make(map[K]int)
	}

	entry := &orderEntry[K, V]{K: key, V: value}
	o.idx[key] = len(o.value)
	o.value = append(o.value, entry)
}

// Set 设置指定 key 的元素
func (o *Order[K, V]) Set(key K, value V) {
	entry, exist := o.idx[key]
	if !exist {
		o.Add(key, value)
	} else {
		o.value[entry].V = value
	}
}

// Len 返回元素数量
func (o *Order[K, V]) Len() int {
	return len(o.value)
}

// Del 删除指定 key 的元素
func (o *Order[K, V]) Del(key K) {
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
func (o *Order[K, V]) Range(handle func(key K, value V) bool) {
	for _, entry := range o.value {
		if !handle(entry.K, entry.V) {
			break
		}
	}
}
