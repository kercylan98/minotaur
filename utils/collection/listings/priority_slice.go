package listings

import (
	"fmt"
	"sort"
)

// NewPrioritySlice 创建一个优先级切片，优先级越低越靠前
func NewPrioritySlice[V any](lengthAndCap ...int) *PrioritySlice[V] {
	p := &PrioritySlice[V]{}
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

// PrioritySlice 是一个优先级切片，优先级越低越靠前
type PrioritySlice[V any] struct {
	items []*priorityItem[V]
}

// Len 返回切片长度
func (slf *PrioritySlice[V]) Len() int {
	return len(slf.items)
}

// Cap 返回切片容量
func (slf *PrioritySlice[V]) Cap() int {
	return cap(slf.items)
}

// Clear 清空切片
func (slf *PrioritySlice[V]) Clear() {
	slf.items = slf.items[:0]
}

// Append 添加元素
func (slf *PrioritySlice[V]) Append(v V, p int) {
	slf.items = append(slf.items, &priorityItem[V]{
		v: v,
		p: p,
	})
	slf.sort()
}

// Appends 添加元素
func (slf *PrioritySlice[V]) Appends(priority int, vs ...V) {
	for _, v := range vs {
		slf.Append(v, priority)
	}
	slf.sort()
}

// Get 获取元素
func (slf *PrioritySlice[V]) Get(index int) (V, int) {
	i := slf.items[index]
	return i.Value(), i.Priority()
}

// GetValue 获取元素值
func (slf *PrioritySlice[V]) GetValue(index int) V {
	return slf.items[index].Value()
}

// GetPriority 获取元素优先级
func (slf *PrioritySlice[V]) GetPriority(index int) int {
	return slf.items[index].Priority()
}

// Set 设置元素
func (slf *PrioritySlice[V]) Set(index int, value V, priority int) {
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
func (slf *PrioritySlice[V]) SetValue(index int, value V) {
	slf.items[index].v = value
}

// SetPriority 设置元素优先级
func (slf *PrioritySlice[V]) SetPriority(index int, priority int) {
	slf.items[index].p = priority
	slf.sort()
}

// Action 直接操作切片，如果返回值不为 nil，则替换切片
func (slf *PrioritySlice[V]) Action(action func(items []*priorityItem[V]) []*priorityItem[V]) {
	if len(slf.items) == 0 {
		return
	}
	if replace := action(slf.items); replace != nil {
		slf.items = replace
		slf.sort()
	}
}

// Range 遍历切片，如果返回值为 false，则停止遍历
func (slf *PrioritySlice[V]) Range(action func(index int, item *priorityItem[V]) bool) {
	for i, item := range slf.items {
		if !action(i, item) {
			break
		}
	}
}

// RangeValue 遍历切片值，如果返回值为 false，则停止遍历
func (slf *PrioritySlice[V]) RangeValue(action func(index int, value V) bool) {
	slf.Range(func(index int, item *priorityItem[V]) bool {
		return action(index, item.Value())
	})
}

// RangePriority 遍历切片优先级，如果返回值为 false，则停止遍历
func (slf *PrioritySlice[V]) RangePriority(action func(index int, priority int) bool) {
	slf.Range(func(index int, item *priorityItem[V]) bool {
		return action(index, item.Priority())
	})
}

// SyncSlice 返回切片
func (slf *PrioritySlice[V]) Slice() []V {
	var vs []V
	for _, item := range slf.items {
		vs = append(vs, item.Value())
	}
	return vs
}

// String 返回切片字符串
func (slf *PrioritySlice[V]) String() string {
	var vs []V
	for _, item := range slf.items {
		vs = append(vs, item.Value())
	}
	return fmt.Sprint(vs)
}

// sort 排序
func (slf *PrioritySlice[V]) sort() {
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

// priorityItem 是一个优先级切片元素
type priorityItem[V any] struct {
	next *priorityItem[V]
	prev *priorityItem[V]
	v    V
	p    int
}

// Value 返回元素值
func (p *priorityItem[V]) Value() V {
	return p.v
}

// Priority 返回元素优先级
func (p *priorityItem[V]) Priority() int {
	return p.p
}

// Next 返回下一个元素
func (p *priorityItem[V]) Next() *priorityItem[V] {
	return p.next
}

// Prev 返回上一个元素
func (p *priorityItem[V]) Prev() *priorityItem[V] {
	return p.prev
}
