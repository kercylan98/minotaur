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
func (s *PagedSlice[T]) Get(index int) *T {
	return &s.pages[index/s.pageSize][index%s.pageSize]
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
