package slice

// Zoom 将切片的长度缩放到指定的大小，如果 newSize 小于 slice 的长度，则会截断 slice，如果 newSize 大于 slice 的长度，则会在 slice 的末尾添加零值数据
func Zoom[V any](newSize int, slice []V) []V {
	if newSize < 0 {
		newSize = 0
	}
	var s = make([]V, newSize)
	copy(s, slice)
	return s
}
