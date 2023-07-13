package slice

import "math/rand"

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

// GetPart 获取数组的部分元素
func GetPart[V any](slice []V, start, end int) []V {
	if start < 0 {
		start = 0
	}
	if end > len(slice) {
		end = len(slice)
	}
	return slice[start:end]
}
