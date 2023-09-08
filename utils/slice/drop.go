package slice

// Drop 从切片中删除指定数量的元素
//   - start: 起始位置
//   - n: 删除的元素数量
//   - slice: 待删除元素的切片
//
// 关于 start 的取值：
//   - 当 start < 0 时，start 将会从右至左计数，即 -1 表示最后一个元素，-2 表示倒数第二个元素，以此类推
func Drop[V any](start, n int, slice []V) []V {
	var s = make([]V, len(slice))
	copy(s, slice)
	if start < 0 {
		start = len(s) + start - n + 1
		if start < 0 {
			start = 0
		}
	}

	end := start + n
	if end > len(s) {
		end = len(s)
	}

	return append(s[:start], s[end:]...)
}

// DropBy 从切片中删除指定的元素
//   - slice: 待删除元素的切片
func DropBy[V any](slice []V, fn func(index int, value V) bool) []V {
	var s = make([]V, len(slice))
	copy(s, slice)
	for i := 0; i < len(s); i++ {
		if fn(i, s[i]) {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}
