package collection

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
)

// FindLoopedNextInSlice 返回 i 的下一个数组成员，当 i 达到数组长度时从 0 开始
//   - 当 i 为负数时将返回第一个元素
func FindLoopedNextInSlice[S ~[]V, V any](slice S, i int) (next int, value V) {
	if i < 0 {
		return 0, slice[0]
	}
	next = i + 1
	if next == len(slice) {
		next = 0
	}
	return next, slice[next]
}

// FindLoopedPrevInSlice 返回 i 的上一个数组成员，当 i 为 0 时从数组末尾开始
//   - 当 i 为负数时将返回最后一个元素
func FindLoopedPrevInSlice[S ~[]V, V any](slice S, i int) (prev int, value V) {
	if i < 0 {
		return len(slice) - 1, slice[len(slice)-1]
	}
	prev = i - 1
	if prev == -1 {
		prev = len(slice) - 1
	}
	return prev, slice[prev]
}

// FindCombinationsInSliceByRange 获取给定数组的所有组合，且每个组合的成员数量限制在指定范围内
func FindCombinationsInSliceByRange[S ~[]V, V any](s S, minSize, maxSize int) []S {
	n := len(s)
	if n == 0 || minSize <= 0 || maxSize <= 0 || minSize > maxSize {
		return nil
	}

	var result []S
	var currentCombination S

	var backtrack func(startIndex int, currentSize int)
	backtrack = func(startIndex int, currentSize int) {
		if currentSize >= minSize && currentSize <= maxSize {
			combination := make(S, len(currentCombination))
			copy(combination, currentCombination)
			result = append(result, combination)
		}

		for i := startIndex; i < n; i++ {
			currentCombination = append(currentCombination, s[i])
			backtrack(i+1, currentSize+1)
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	backtrack(0, 0)
	return result
}

// FindFirstOrDefaultInSlice 判断切片中是否存在元素，返回第一个元素，不存在则返回默认值
func FindFirstOrDefaultInSlice[S ~[]V, V any](slice S, defaultValue V) V {
	if len(slice) == 0 {
		return defaultValue
	}
	return slice[0]
}

// FindOrDefaultInSlice 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则返回默认值
func FindOrDefaultInSlice[S ~[]V, V any](slice S, defaultValue V, handler func(v V) bool) (t V) {
	if len(slice) == 0 {
		return defaultValue
	}
	for _, v := range slice {
		if handler(v) {
			return v
		}
	}
	return defaultValue
}

// FindOrDefaultInComparableSlice 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则返回默认值
func FindOrDefaultInComparableSlice[S ~[]V, V comparable](slice S, v V, defaultValue V) (t V) {
	if len(slice) == 0 {
		return defaultValue
	}
	for _, value := range slice {
		if value == v {
			return value
		}
	}
	return defaultValue
}

// FindInSlice 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则索引返回 -1
func FindInSlice[S ~[]V, V any](slice S, handler func(v V) bool) (i int, t V) {
	if len(slice) == 0 {
		return -1, t
	}
	for i, v := range slice {
		if handler(v) {
			return i, v
		}
	}
	return -1, t
}

// FindIndexInSlice 判断切片中是否存在某个元素，返回第一个匹配的索引，不存在则索引返回 -1
func FindIndexInSlice[S ~[]V, V any](slice S, handler func(v V) bool) int {
	if len(slice) == 0 {
		return -1
	}
	for i, v := range slice {
		if handler(v) {
			return i
		}
	}
	return -1
}

// FindInComparableSlice 判断切片中是否存在某个元素，返回第一个匹配的索引和元素，不存在则索引返回 -1
func FindInComparableSlice[S ~[]V, V comparable](slice S, v V) (i int, t V) {
	if len(slice) == 0 {
		return -1, t
	}
	for i, value := range slice {
		if value == v {
			return i, value
		}
	}
	return -1, t
}

// FindIndexInComparableSlice 判断切片中是否存在某个元素，返回第一个匹配的索引，不存在则索引返回 -1
func FindIndexInComparableSlice[S ~[]V, V comparable](slice S, v V) int {
	if len(slice) == 0 {
		return -1
	}
	for i, value := range slice {
		if value == v {
			return i
		}
	}
	return -1
}

// FindMinimumInComparableSlice 获取切片中的最小值
func FindMinimumInComparableSlice[S ~[]V, V constraints.Ordered](slice S) (result V) {
	if len(slice) == 0 {
		return
	}
	result = slice[0]
	for i := 1; i < len(slice); i++ {
		if result > slice[i] {
			result = slice[i]
		}
	}
	return
}

// FindMinimumInSlice 获取切片中的最小值
func FindMinimumInSlice[S ~[]V, V any, N constraints.Ordered](slice S, handler OrderedValueGetter[V, N]) (result V) {
	if len(slice) == 0 {
		return
	}
	result = slice[0]
	for i := 1; i < len(slice); i++ {
		if handler(result) > handler(slice[i]) {
			result = slice[i]
		}
	}
	return
}

// FindMaximumInComparableSlice 获取切片中的最大值
func FindMaximumInComparableSlice[S ~[]V, V constraints.Ordered](slice S) (result V) {
	if len(slice) == 0 {
		return
	}
	result = slice[0]
	for i := 1; i < len(slice); i++ {
		if result < slice[i] {
			result = slice[i]
		}
	}
	return
}

// FindMaximumInSlice 获取切片中的最大值
func FindMaximumInSlice[S ~[]V, V any, N constraints.Ordered](slice S, handler OrderedValueGetter[V, N]) (result V) {
	if len(slice) == 0 {
		return
	}
	result = slice[0]
	for i := 1; i < len(slice); i++ {
		if handler(result) < handler(slice[i]) {
			result = slice[i]
		}
	}
	return
}

// FindMin2MaxInComparableSlice 获取切片中的最小值和最大值
func FindMin2MaxInComparableSlice[S ~[]V, V constraints.Ordered](slice S) (min, max V) {
	if len(slice) == 0 {
		return
	}
	min = slice[0]
	max = slice[0]
	for i := 1; i < len(slice); i++ {
		if min > slice[i] {
			min = slice[i]
		}
		if max < slice[i] {
			max = slice[i]
		}
	}
	return
}

// FindMin2MaxInSlice 获取切片中的最小值和最大值
func FindMin2MaxInSlice[S ~[]V, V any, N constraints.Ordered](slice S, handler OrderedValueGetter[V, N]) (min, max V) {
	if len(slice) == 0 {
		return
	}
	min = slice[0]
	max = slice[0]
	for i := 1; i < len(slice); i++ {
		if handler(min) > handler(slice[i]) {
			min = slice[i]
		}
		if handler(max) < handler(slice[i]) {
			max = slice[i]
		}
	}
	return
}

// FindMinFromComparableMap 获取 map 中的最小值
func FindMinFromComparableMap[M ~map[K]V, K comparable, V constraints.Ordered](m M) (result V) {
	if m == nil {
		return
	}
	var first bool
	for _, v := range m {
		if !first {
			result = v
			first = true
			continue
		}
		if result > v {
			result = v
		}
	}
	return
}

// FindMinFromMap 获取 map 中的最小值
func FindMinFromMap[M ~map[K]V, K comparable, V any, N constraints.Ordered](m M, handler OrderedValueGetter[V, N]) (result V) {
	if m == nil {
		return
	}
	var first bool
	for _, v := range m {
		if !first {
			result = v
			first = true
			continue
		}
		if handler(result) > handler(v) {
			result = v
		}
	}
	return
}

// FindMaxFromComparableMap 获取 map 中的最大值
func FindMaxFromComparableMap[M ~map[K]V, K comparable, V constraints.Ordered](m M) (result V) {
	if m == nil {
		return
	}
	for _, v := range m {
		if result < v {
			result = v
		}
	}
	return
}

// FindMaxFromMap 获取 map 中的最大值
func FindMaxFromMap[M ~map[K]V, K comparable, V any, N constraints.Ordered](m M, handler OrderedValueGetter[V, N]) (result V) {
	if m == nil {
		return
	}
	for _, v := range m {
		if handler(result) < handler(v) {
			result = v
		}
	}
	return
}

// FindMin2MaxFromComparableMap 获取 map 中的最小值和最大值
func FindMin2MaxFromComparableMap[M ~map[K]V, K comparable, V constraints.Ordered](m M) (min, max V) {
	if m == nil {
		return
	}
	var first bool
	for _, v := range m {
		if !first {
			min = v
			max = v
			first = true
			continue
		}
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return
}

// FindMin2MaxFromMap 获取 map 中的最小值和最大值
func FindMin2MaxFromMap[M ~map[K]V, K comparable, V constraints.Ordered](m M) (min, max V) {
	if m == nil {
		return
	}
	var first bool
	for _, v := range m {
		if !first {
			min = v
			max = v
			first = true
			continue
		}
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return
}

// IsFirst 判断是否是第一个元素，当 slice 长度为 0 时，返回 false
func IsFirst[S ~[]V, V comparable](slice S, v V) bool {
	if len(slice) == 0 {
		return false
	}
	return slice[0] == v
}
