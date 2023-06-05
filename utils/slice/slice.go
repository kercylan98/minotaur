package slice

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
