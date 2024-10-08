package listings

// NewPagedSlice 创建一个新的 PagedSlice 实例。
func NewPagedSlice[T any](pageSize int) *PagedSlice[T] {
	ps := &PagedSlice[T]{
		pageSize: pageSize,
	}
	return ps
}

// PagedSlice 是一个高效的动态数组，它通过分页管理内存并减少频繁的内存分配来提高性能。
type PagedSlice[T any] struct {
	pages    [][]T
	pageSize int
	len      int
	lenLast  int
}

// Grow 扩容 PagedSlice 的长度但不设置值
func (s *PagedSlice[T]) Grow(indexes []int) {
	if len(indexes) == 0 {
		return
	}

	// 找到最大索引
	maxIndex := indexes[0]
	for _, index := range indexes {
		if index > maxIndex {
			maxIndex = index
		}
	}

	// 扩展 PagedSlice 的长度，以便容纳所有索引
	if maxIndex >= s.len {
		totalPages := (maxIndex / s.pageSize) + 1
		for len(s.pages) < totalPages {
			s.pages = append(s.pages, make([]T, s.pageSize))
		}
		s.len = maxIndex + 1
		s.lenLast = (maxIndex % s.pageSize) + 1
	}
}

// GrowSet 扩展 PagedSlice 的长度并设置给定索引的元素。
func (s *PagedSlice[T]) GrowSet(index int, value T) {
	// 如果索引超出了当前长度，则扩展 PagedSlice 的长度。
	if index >= s.len {
		// 计算需要的总页数
		totalPages := (index / s.pageSize) + 1

		// 扩展页数
		for len(s.pages) < totalPages {
			s.pages = append(s.pages, make([]T, s.pageSize))
		}

		// 更新长度
		s.len = index + 1
		s.lenLast = (index % s.pageSize) + 1
	}

	// 设置给定索引的元素
	pageIndex := index / s.pageSize
	elementIndex := index % s.pageSize
	s.pages[pageIndex][elementIndex] = value
}

// BatchGrowSet 批量设置给定索引的元素。
func (s *PagedSlice[T]) BatchGrowSet(indexes []int, values []T) {
	if len(indexes) != len(values) {
		panic("indexes and values slices must have the same length")
	}

	if len(indexes) == 0 {
		return
	}

	// 找到最大索引
	maxIndex := indexes[0]
	for _, index := range indexes {
		if index > maxIndex {
			maxIndex = index
		}
	}

	// 扩展 PagedSlice 的长度，以便容纳所有索引
	if maxIndex >= s.len {
		totalPages := (maxIndex / s.pageSize) + 1
		for len(s.pages) < totalPages {
			s.pages = append(s.pages, make([]T, s.pageSize))
		}
		s.len = maxIndex + 1
		s.lenLast = (maxIndex % s.pageSize) + 1
	}

	// 批量设置元素
	for i, index := range indexes {
		pageIndex := index / s.pageSize
		elementIndex := index % s.pageSize
		s.pages[pageIndex][elementIndex] = values[i]
	}
}

// BatchSet 批量设置给定索引的元素。
func (s *PagedSlice[T]) BatchSet(indexes []int, values []T) {
	if len(indexes) != len(values) {
		panic("indexes and values slices must have the same length")
	}

	if len(indexes) == 0 {
		return
	}

	// 批量设置元素
	for i, index := range indexes {
		pageIndex := index / s.pageSize
		elementIndex := index % s.pageSize
		s.pages[pageIndex][elementIndex] = values[i]
	}
}

// Add 添加一个元素到 PagedSlice 中。
func (s *PagedSlice[T]) Add(value T) {
	if len(s.pages) == 0 || s.lenLast == len(s.pages[len(s.pages)-1]) {
		s.pages = append(s.pages, make([]T, s.pageSize))
		s.lenLast = 0
	}

	s.pages[len(s.pages)-1][s.lenLast] = value
	s.len++
	s.lenLast++
}

// Del 删除 PagedSlice 中给定索引的元素。
func (s *PagedSlice[T]) Del(index int) {
	if index < 0 || index >= s.len {
		return
	}

	lastIndex := s.len - 1
	s.pages[index/s.pageSize][index%s.pageSize] = s.pages[lastIndex/s.pageSize][lastIndex%s.pageSize]

	s.len--
	if s.len%s.pageSize == 0 && len(s.pages) > 1 {
		s.pages = s.pages[:len(s.pages)-1]
		s.lenLast = s.pageSize
	} else {
		s.lenLast = s.len % s.pageSize
	}
}

// Get 获取 PagedSlice 中给定索引的元素。
func (s *PagedSlice[T]) Get(index int) (v *T) {
	if s == nil {
		return
	}
	res := s.pages[index/s.pageSize][index%s.pageSize]
	return &res
}

// Set 设置 PagedSlice 中给定索引的元素。
func (s *PagedSlice[T]) Set(index int, value T) {
	if index < 0 || index >= s.len {
		return
	}

	s.pages[index/s.pageSize][index%s.pageSize] = value
}

// Len 返回 PagedSlice 中元素的数量。
func (s *PagedSlice[T]) Len() int {
	return s.len
}
