package listings

import (
	"fmt"
	"sort"
	"sync"
)

// NewSyncPrioritySlice 创建一个并发安全的优先级切片，优先级越低越靠前
func NewSyncPrioritySlice[V any](lengthAndCap ...int) *SyncPrioritySlice[V] {
	p := &SyncPrioritySlice[V]{}
	if len(lengthAndCap) > 0 {
		var length = lengthAndCap[0]
		var c int
		if len(lengthAndCap) > 1 {
			c = lengthAndCap[1]
		}
		p.items = make([]*priorityItem[V], length, c)
	}
	return p
}

// SyncPrioritySlice 是一个优先级切片，优先级越低越靠前
type SyncPrioritySlice[V any] struct {
	rw    sync.RWMutex
	items []*priorityItem[V]
}

// Len 返回切片长度
func (slf *SyncPrioritySlice[V]) Len() int {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	return len(slf.items)
}

// Cap 返回切片容量
func (slf *SyncPrioritySlice[V]) Cap() int {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	return cap(slf.items)
}

// Clear 清空切片
func (slf *SyncPrioritySlice[V]) Clear() {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.items = slf.items[:0]
}

// Append 添加元素
func (slf *SyncPrioritySlice[V]) Append(v V, p int) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.items = append(slf.items, &priorityItem[V]{
		v: v,
		p: p,
	})
	slf.sort()
}

// Appends 添加元素
func (slf *SyncPrioritySlice[V]) Appends(priority int, vs ...V) {
	for _, v := range vs {
		slf.Append(v, priority)
	}
	slf.sort()
}

// AppendByOptionalPriority 添加元素
func (slf *SyncPrioritySlice[V]) AppendByOptionalPriority(v V, priority ...int) {
	if len(priority) == 0 {
		slf.Append(v, 0)
	} else {
		slf.Append(v, priority[0])
	}
}

// Get 获取元素
func (slf *SyncPrioritySlice[V]) Get(index int) (V, int) {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	i := slf.items[index]
	return i.Value(), i.Priority()
}

// GetValue 获取元素值
func (slf *SyncPrioritySlice[V]) GetValue(index int) V {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	return slf.items[index].Value()
}

// GetPriority 获取元素优先级
func (slf *SyncPrioritySlice[V]) GetPriority(index int) int {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	return slf.items[index].Priority()
}

// Set 设置元素
func (slf *SyncPrioritySlice[V]) Set(index int, value V, priority int) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	before := slf.items[index]
	slf.items[index] = &priorityItem[V]{
		v: value,
		p: priority,
	}
	if before.Priority() != priority {
		slf.sort()
	}
}

// SetValue 设置元素值
func (slf *SyncPrioritySlice[V]) SetValue(index int, value V) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.items[index].v = value
}

// SetPriority 设置元素优先级
func (slf *SyncPrioritySlice[V]) SetPriority(index int, priority int) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.items[index].p = priority
	slf.sort()
}

// Action 直接操作切片，如果返回值不为 nil，则替换切片
func (slf *SyncPrioritySlice[V]) Action(action func(items []*priorityItem[V]) []*priorityItem[V]) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	if len(slf.items) == 0 {
		return
	}
	if replace := action(slf.items); replace != nil {
		slf.items = replace
		slf.sort()
	}
}

// Range 遍历切片，如果返回值为 false，则停止遍历
func (slf *SyncPrioritySlice[V]) Range(action func(index int, item *priorityItem[V]) bool) {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	for i, item := range slf.items {
		if !action(i, item) {
			break
		}
	}
}

// RangeValue 遍历切片值，如果返回值为 false，则停止遍历
func (slf *SyncPrioritySlice[V]) RangeValue(action func(index int, value V) bool) {
	slf.Range(func(index int, item *priorityItem[V]) bool {
		return action(index, item.Value())
	})
}

// RangePriority 遍历切片优先级，如果返回值为 false，则停止遍历
func (slf *SyncPrioritySlice[V]) RangePriority(action func(index int, priority int) bool) {
	slf.Range(func(index int, item *priorityItem[V]) bool {
		return action(index, item.Priority())
	})
}

// Slice 返回切片
func (slf *SyncPrioritySlice[V]) Slice() []V {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	var vs []V
	for _, item := range slf.items {
		vs = append(vs, item.Value())
	}
	return vs
}

// String 返回切片字符串
func (slf *SyncPrioritySlice[V]) String() string {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	var vs []V
	for _, item := range slf.items {
		vs = append(vs, item.Value())
	}
	return fmt.Sprint(vs)
}

// sort 排序
func (slf *SyncPrioritySlice[V]) sort() {
	if len(slf.items) <= 1 {
		return
	}
	sort.Slice(slf.items, func(i, j int) bool {
		return slf.items[i].Priority() < slf.items[j].Priority()
	})
	for i := 0; i < len(slf.items); i++ {
		if i == 0 {
			slf.items[i].prev = nil
			slf.items[i].next = slf.items[i+1]
		} else if i == len(slf.items)-1 {
			slf.items[i].prev = slf.items[i-1]
			slf.items[i].next = nil
		} else {
			slf.items[i].prev = slf.items[i-1]
			slf.items[i].next = slf.items[i+1]
		}
	}
}
