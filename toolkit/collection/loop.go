package collection

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"sort"
)

// LoopSlice 迭代切片 slice 中的每一个函数，并将索引和值传递给 f 函数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopSlice[S ~[]V, V any](slice S, f func(i int, val V) bool) {
	for i, v := range slice {
		if !f(i, v) {
			break
		}
	}
}

// ReverseLoopSlice 逆序迭代切片 slice 中的每一个函数，并将索引和值传递给 f 函数
//   - 迭代过程将在 f 函数返回 false 时中断
func ReverseLoopSlice[S ~[]V, V any](slice S, f func(i int, val V) bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if !f(i, slice[i]) {
			break
		}
	}
}

// LoopMap 迭代 m 中的每一个函数，并将键和值传递给 f 函数
//   - m 的迭代顺序是不确定的，因此每次迭代的顺序可能不同
//   - 该函数会在 f 中传入一个从 0 开始的索引，用于表示当前迭代的次数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopMap[M ~map[K]V, K comparable, V any](m M, f func(i int, key K, val V) bool) {
	var i int
	for k, v := range m {
		if !f(i, k, v) {
			break
		}
		i++
	}
}

// LoopMapByOrderedKeyAsc 按照键的升序迭代 m 中的每一个函数，并将键和值传递给 f 函数
//   - 该函数会在 f 中传入一个从 0 开始的索引，用于表示当前迭代的次数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopMapByOrderedKeyAsc[M ~map[K]V, K constraints.Ordered, V any](m M, f func(i int, key K, val V) bool) {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return AscBy(keys[i], keys[j])
	})
	for i, k := range keys {
		if !f(i, k, m[k]) {
			break
		}
	}
}

// LoopMapByOrderedKeyDesc 按照键的降序迭代 m 中的每一个函数，并将键和值传递给 f 函数
//   - 该函数会在 f 中传入一个从 0 开始的索引，用于表示当前迭代的次数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopMapByOrderedKeyDesc[M ~map[K]V, K constraints.Ordered, V any](m M, f func(i int, key K, val V) bool) {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return DescBy(keys[i], keys[j])
	})
	for i, k := range keys {
		if !f(i, k, m[k]) {
			break
		}
	}
}

// LoopMapByOrderedValueAsc 按照值的升序迭代 m 中的每一个函数，并将键和值传递给 f 函数
//   - 该函数会在 f 中传入一个从 0 开始的索引，用于表示当前迭代的次数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopMapByOrderedValueAsc[M ~map[K]V, K comparable, V constraints.Ordered](m M, f func(i int, key K, val V) bool) {
	var keys []K
	var values []V
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return AscBy(values[i], values[j])
	})
	for i, v := range values {
		if !f(i, keys[i], v) {
			break
		}
	}
}

// LoopMapByOrderedValueDesc 按照值的降序迭代 m 中的每一个函数，并将键和值传递给 f 函数
//   - 该函数会在 f 中传入一个从 0 开始的索引，用于表示当前迭代的次数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopMapByOrderedValueDesc[M ~map[K]V, K comparable, V constraints.Ordered](m M, f func(i int, key K, val V) bool) {
	var keys []K
	var values []V
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return DescBy(values[i], values[j])
	})
	for i, v := range values {
		if !f(i, keys[i], v) {
			break
		}
	}
}

// LoopMapByKeyGetterAsc 按照键的升序迭代 m 中的每一个函数，并将键和值传递给 f 函数
//   - 该函数会在 f 中传入一个从 0 开始的索引，用于表示当前迭代的次数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopMapByKeyGetterAsc[M ~map[K]V, K comparable, V comparable, N constraints.Ordered](m M, getter func(k K) N, f func(i int, key K, val V) bool) {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return AscBy(getter(keys[i]), getter(keys[j]))
	})
	for i, v := range keys {
		if !f(i, keys[i], m[v]) {
			break
		}
	}
}

// LoopMapByValueGetterAsc 按照值的升序迭代 m 中的每一个函数，并将键和值传递给 f 函数
//   - 该函数会在 f 中传入一个从 0 开始的索引，用于表示当前迭代的次数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopMapByValueGetterAsc[M ~map[K]V, K comparable, V any, N constraints.Ordered](m M, getter func(v V) N, f func(i int, key K, val V) bool) {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return AscBy(getter(m[keys[i]]), getter(m[keys[j]]))
	})
	for i, v := range keys {
		if !f(i, keys[i], m[v]) {
			break
		}
	}
}

// LoopMapByKeyGetterDesc 按照键的降序迭代 m 中的每一个函数，并将键和值传递给 f 函数
//   - 该函数会在 f 中传入一个从 0 开始的索引，用于表示当前迭代的次数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopMapByKeyGetterDesc[M ~map[K]V, K comparable, V comparable, N constraints.Ordered](m M, getter func(k K) N, f func(i int, key K, val V) bool) {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return DescBy(getter(keys[i]), getter(keys[j]))
	})
	for i, v := range keys {
		if !f(i, keys[i], m[v]) {
			break
		}
	}
}

// LoopMapByValueGetterDesc 按照值的降序迭代 m 中的每一个函数，并将键和值传递给 f 函数
//   - 该函数会在 f 中传入一个从 0 开始的索引，用于表示当前迭代的次数
//   - 迭代过程将在 f 函数返回 false 时中断
func LoopMapByValueGetterDesc[M ~map[K]V, K comparable, V any, N constraints.Ordered](m M, getter func(v V) N, f func(i int, key K, val V) bool) {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return DescBy(getter(m[keys[i]]), getter(m[keys[j]]))
	})
	for i, v := range keys {
		if !f(i, keys[i], m[v]) {
			break
		}
	}
}
