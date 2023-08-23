package stream

import (
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/slice"
	"reflect"
)

// WithSlice 创建一个 Slice
//   - 该函数不会影响到传入的 slice
func WithSlice[V any](values []V) Slice[V] {
	return slice.Copy(values)
}

// Slice 提供了 slice 的链式操作
type Slice[V any] []V

// Filter 过滤 handle 返回 false 的元素
func (slf Slice[V]) Filter(handle func(index int, value V) bool) Slice[V] {
	var ns = make([]V, 0, len(slf))
	for i, v := range slf {
		if handle(i, v) {
			ns = append(ns, v)
		}
	}
	return ns
}

// FilterValue 过滤特定的 value
func (slf Slice[V]) FilterValue(values ...V) Slice[V] {
	return slf.Filter(func(index int, value V) bool {
		for _, v := range values {
			if reflect.DeepEqual(v, value) {
				return false
			}
		}
		return true
	})
}

// FilterIndex 过滤特定的 index
func (slf Slice[V]) FilterIndex(indexes ...int) Slice[V] {
	return slf.Filter(func(index int, value V) bool {
		return !slice.Contains(indexes, index)
	})
}

// RandomKeep 随机保留 n 个元素
func (slf Slice[V]) RandomKeep(n int) Slice[V] {
	length := len(slf)
	if n >= length {
		return slf
	}
	var hit = make([]int, length)
	for i := 0; i < n; i++ {
		hit[i] = 1
	}
	slice.Shuffle(hit)
	var ns = make([]V, 0, n)
	for i, v := range slf {
		if hit[i] == 1 {
			ns = append(ns, v)
		}
	}
	return ns
}

// RandomDelete 随机删除 n 个元素
func (slf Slice[V]) RandomDelete(n int) Slice[V] {
	length := len(slf)
	if n >= length {
		return slf[:0]
	}
	var hit = make([]int, length)
	for i := 0; i < n; i++ {
		hit[i] = 1
	}
	slice.Shuffle(hit)
	var ns = make([]V, 0, n)
	for i, v := range slf {
		if hit[i] == 0 {
			ns = append(ns, v)
		}
	}
	return ns
}

// Shuffle 随机打乱
func (slf Slice[V]) Shuffle() Slice[V] {
	slice.Shuffle(slf)
	return slf
}

// Reverse 反转
func (slf Slice[V]) Reverse() Slice[V] {
	slice.Reverse(slf)
	return slf
}

// Clear 清空
func (slf Slice[V]) Clear() Slice[V] {
	return slf[:0]
}

// Distinct 去重，如果 handle 返回 true 则认为是重复元素
func (slf Slice[V]) Distinct() Slice[V] {
	return slice.Distinct(slf)
}

// Merge 合并
func (slf Slice[V]) Merge(values ...V) Slice[V] {
	return append(slf, values...)
}

// GetStartPart 获取前 n 个元素
func (slf Slice[V]) GetStartPart(n int) Slice[V] {
	return slf[:n]
}

// GetEndPart 获取后 n 个元素
func (slf Slice[V]) GetEndPart(n int) Slice[V] {
	return slice.GetEndPart(slf, n)
}

// GetPart 获取指定区间的元素
func (slf Slice[V]) GetPart(start, end int) Slice[V] {
	return slice.GetPart(slf, start, end)
}

// ContainsHandle 如果包含指定的元素则执行 handle
func (slf Slice[V]) ContainsHandle(value V, handle func(slice Slice[V]) Slice[V]) Slice[V] {
	if slice.ContainsAny(slf, value) {
		return handle(slf)
	}
	return slf
}

// Set 设置指定位置的元素
func (slf Slice[V]) Set(index int, value V) Slice[V] {
	slf[index] = value
	return slf
}

// Delete 删除指定位置的元素
func (slf Slice[V]) Delete(index int) Slice[V] {
	return append(slf, slf[index+1:]...)
}

// ToMapStream 将当前的 Slice stream 转换为 Map stream
func (slf Slice[V]) ToMapStream() Map[int, V] {
	return hash.ToMap(slf)
}
