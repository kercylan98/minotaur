package slice

import (
	"fmt"
	"sort"
)

// NewPriority 创建一个优先级切片
func NewPriority[V any](lengthAndCap ...int) *Priority[V] {
	p := &Priority[V]{}
	if len(lengthAndCap) > 0 {
		var length = lengthAndCap[0]
		var c int
		if len(lengthAndCap) > 1 {
			c = lengthAndCap[1]
		}
		p.items = make([]*PriorityItem[V], length, c)
	}
	return p
}

// Priority 是一个优先级切片
type Priority[V any] struct {
	items []*PriorityItem[V]
}

// Len 返回切片长度
func (slf *Priority[V]) Len() int {
	return len(slf.items)
}

// Cap 返回切片容量
func (slf *Priority[V]) Cap() int {
	return cap(slf.items)
}

// Clear 清空切片
func (slf *Priority[V]) Clear() {
	clear(slf.items)
}

// Append 添加元素
func (slf *Priority[V]) Append(v V, priority int) {
	slf.items = append(slf.items, NewPriorityItem[V](v, priority))
	slf.sort()
}

// Appends 添加元素
func (slf *Priority[V]) Appends(priority int, vs ...V) {
	for _, v := range vs {
		slf.Append(v, priority)
	}
	slf.sort()
}

// Get 获取元素
func (slf *Priority[V]) Get(index int) *PriorityItem[V] {
	return slf.items[index]
}

// GetValue 获取元素值
func (slf *Priority[V]) GetValue(index int) V {
	return slf.items[index].Value()
}

// GetPriority 获取元素优先级
func (slf *Priority[V]) GetPriority(index int) int {
	return slf.items[index].Priority()
}

// Set 设置元素
func (slf *Priority[V]) Set(index int, value V, priority int) {
	before := slf.items[index]
	slf.items[index] = NewPriorityItem[V](value, priority)
	if before.Priority() != priority {
		slf.sort()
	}
}

// SetValue 设置元素值
func (slf *Priority[V]) SetValue(index int, value V) {
	slf.items[index].v = value
}

// SetPriority 设置元素优先级
func (slf *Priority[V]) SetPriority(index int, priority int) {
	slf.items[index].p = priority
	slf.sort()
}

// Action 直接操作切片，如果返回值不为 nil，则替换切片
func (slf *Priority[V]) Action(action func(items []*PriorityItem[V]) []*PriorityItem[V]) {
	if len(slf.items) == 0 {
		return
	}
	if replace := action(slf.items); replace != nil {
		slf.items = replace
		slf.sort()
	}
}

// Range 遍历切片，如果返回值为 false，则停止遍历
func (slf *Priority[V]) Range(action func(index int, item *PriorityItem[V]) bool) {
	for i, item := range slf.items {
		if !action(i, item) {
			break
		}
	}
}

// RangeValue 遍历切片值，如果返回值为 false，则停止遍历
func (slf *Priority[V]) RangeValue(action func(index int, value V) bool) {
	slf.Range(func(index int, item *PriorityItem[V]) bool {
		return action(index, item.Value())
	})
}

// RangePriority 遍历切片优先级，如果返回值为 false，则停止遍历
func (slf *Priority[V]) RangePriority(action func(index int, priority int) bool) {
	slf.Range(func(index int, item *PriorityItem[V]) bool {
		return action(index, item.Priority())
	})
}

// String 返回切片字符串
func (slf *Priority[V]) String() string {
	var vs []V
	for _, item := range slf.items {
		vs = append(vs, item.Value())
	}
	return fmt.Sprint(vs)
}

// sort 排序
func (slf *Priority[V]) sort() {
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
