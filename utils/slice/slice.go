package slice

import (
	"math/rand"
	"reflect"
)

// Del 删除特定索引的元素
func Del[V any](slice *[]V, index int) {
	s := *slice
	*slice = append(s[:index], s[index+1:]...)
}

// Copy 复制特定切片
func Copy[V any](slice []V) []V {
	var s = make([]V, len(slice), len(slice))
	for i := 0; i < len(slice); i++ {
		s[i] = slice[i]
	}
	return s
}

// CopyMatrix 复制二维数组
func CopyMatrix[V any](slice [][]V) [][]V {
	var s = make([][]V, len(slice), len(slice))
	for i := 0; i < len(slice); i++ {
		is := make([]V, len(slice[0]))
		for j := 0; j < len(slice[0]); j++ {
			is[j] = slice[i][j]
		}
		s[i] = is
	}
	return s
}

// Insert 在特定索引插入元素
func Insert[V any](slice *[]V, index int, value V) {
	s := *slice
	if index <= 0 {
		*slice = append([]V{value}, s...)
	} else if index >= len(s) {
		*slice = append(s, value)
	} else {
		*slice = append(s[:index], append([]V{value}, s[index:]...)...)
	}
}

// Move 移动特定索引
func Move[V any](slice *[]V, index, to int) {
	s := *slice
	v := s[index]
	if index == to {
		return
	} else if to < index {
		Del[V](slice, index)
		Insert(slice, to, v)
	} else {
		Insert(slice, to, v)
		Del[V](slice, index)
	}
}

// NextLoop 返回 i 的下一个数组成员，当 i 达到数组长度时从 0 开始
//   - 当 i 为 -1 时将返回第一个元素
func NextLoop[V any](slice []V, i int) (next int, value V) {
	if i == -1 {
		return 0, slice[0]
	}
	next = i + 1
	if next == len(slice) {
		next = 0
	}
	return next, slice[next]
}

// PrevLoop 返回 i 的上一个数组成员，当 i 为 0 时从数组末尾开始
//   - 当 i 为 -1 时将返回最后一个元素
func PrevLoop[V any](slice []V, i int) (prev int, value V) {
	if i == -1 {
		return len(slice) - 1, slice[len(slice)-1]
	}
	prev = i - 1
	if prev == -1 {
		prev = len(slice) - 1
	}
	return prev, slice[prev]
}

// Reverse 反转数组
func Reverse[V any](slice []V) {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
}

// Shuffle 随机打乱数组
func Shuffle[V any](slice []V) {
	for i := 0; i < len(slice); i++ {
		j := rand.Intn(len(slice))
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Distinct 去重
func Distinct[V any](slice []V) []V {
	var result []V
	for i := range slice {
		flag := true
		for j := range result {
			if reflect.DeepEqual(slice[i], result[j]) {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slice[i])
		}
	}
	return result
}

// Swap 交换数组中的两个元素
func Swap[V any](slice []V, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// ToMap 将数组转换为 map
func ToMap[K comparable, V any](slice []V, key func(V) K) map[K]V {
	m := make(map[K]V)
	for _, v := range slice {
		m[key(v)] = v
	}
	return m
}

// ToSet 将数组转换为 set
func ToSet[V comparable](slice []V) map[V]struct{} {
	m := make(map[V]struct{})
	for _, v := range slice {
		m[v] = struct{}{}
	}
	return m
}

// Merge 合并多个数组
func Merge[V any](slices ...[]V) []V {
	var slice []V
	for _, s := range slices {
		slice = append(slice, s...)
	}
	return slice
}

// GetStartPart 获取数组的前 n 个元素
func GetStartPart[V any](slice []V, n int) []V {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:n]
}

// GetEndPart 获取数组的后 n 个元素
func GetEndPart[V any](slice []V, n int) []V {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[len(slice)-n:]
}

// GetPart 获取指定区间的元素
func GetPart[V any](slice []V, start, end int) []V {
	if start < 0 {
		start = 0
	}
	if end > len(slice) {
		end = len(slice)
	}
	return slice[start:end]
}

// Contains 判断数组是否包含某个元素
func Contains[V comparable](slice []V, value V) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsAny 判断数组是否包含某个元素
func ContainsAny[V any](slice []V, values V) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, values) {
			return true
		}
	}
	return false
}

// GetIndex 判断数组是否包含某个元素，如果包含则返回索引
func GetIndex[V comparable](slice []V, value V) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// GetIndexAny 判断数组是否包含某个元素，如果包含则返回索引
func GetIndexAny[V any](slice []V, values V) int {
	for i, v := range slice {
		if reflect.DeepEqual(v, values) {
			return i
		}
	}
	return -1
}

// Combinations 获取给定数组的所有组合，包括重复元素的组合
func Combinations[T any](a []T) [][]T {
	n := len(a)

	// 去除重复元素，保留唯一元素
	uniqueSet := make(map[uintptr]bool)
	uniqueSlice := make([]T, 0, n)
	for _, val := range a {
		ptr := reflect.ValueOf(val).Pointer()
		if !uniqueSet[ptr] {
			uniqueSet[ptr] = true
			uniqueSlice = append(uniqueSlice, val)
		}
	}

	n = len(uniqueSlice)        // 去重后的数组长度
	totalCombinations := 1 << n // 2的n次方
	var result [][]T
	for i := 0; i < totalCombinations; i++ {
		var currentCombination []T
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				currentCombination = append(currentCombination, uniqueSlice[j])
			}
		}
		result = append(result, currentCombination)
	}
	return result
}

// LimitedCombinations 获取给定数组的所有组合，且每个组合的成员数量限制在指定范围内
func LimitedCombinations[T any](a []T, minSize, maxSize int) [][]T {
	n := len(a)
	if n == 0 || minSize <= 0 || maxSize <= 0 || minSize > maxSize {
		return nil
	}

	var result [][]T
	var currentCombination []T

	var backtrack func(startIndex int, currentSize int)
	backtrack = func(startIndex int, currentSize int) {
		if currentSize >= minSize && currentSize <= maxSize {
			combination := make([]T, len(currentCombination))
			copy(combination, currentCombination)
			result = append(result, combination)
		}

		for i := startIndex; i < n; i++ {
			currentCombination = append(currentCombination, a[i])
			backtrack(i+1, currentSize+1)
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	backtrack(0, 0)
	return result
}

// IsIntersectWithCheck 判断两个切片是否有交集
func IsIntersectWithCheck[T any](a, b []T, checkHandle func(a, b T) bool) bool {
	for _, a := range a {
		for _, b := range b {
			if checkHandle(a, b) {
				return true
			}
		}
	}
	return false
}

// IsIntersect 判断两个切片是否有交集
func IsIntersect[T any](a, b []T) bool {
	for _, a := range a {
		for _, b := range b {
			if reflect.DeepEqual(a, b) {
				return true
			}
		}
	}
	return false
}

// SubWithCheck 获取移除指定元素后的切片
//   - checkHandle 返回 true 表示需要移除
func SubWithCheck[T any](a, b []T, checkHandle func(a, b T) bool) []T {
	var result []T
	for _, a := range a {
		flag := false
		for _, b := range b {
			if checkHandle(a, b) {
				flag = true
				break
			}
		}
		if !flag {
			result = append(result, a)
		}
	}
	return result
}
