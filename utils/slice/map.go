package slice

// Map 将切片中的每一个元素转换为另一种类型的元素
//   - slice: 待转换的切片
//   - converter: 转换函数
//
// 这不会改变原有的切片，而是返回一个新的切片
func Map[V any, T any](slice []V, converter func(index int, item V) T) []T {
	var s = make([]T, len(slice))
	for i := 0; i < len(slice); i++ {
		s[i] = converter(i, slice[i])
	}
	return s
}
