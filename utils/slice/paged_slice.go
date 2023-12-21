package slice

// NewPagedSlice 创建一个新的 PagedSlice 实例。
func NewPagedSlice[T any](pageSize int) *PagedSlice[T] {
	return &PagedSlice[T]{
		pages:    make([][]T, 0, pageSize),
		pageSize: pageSize,
	}
}

// PagedSlice 是一个高效的动态数组，它通过分页管理内存并减少频繁的内存分配来提高性能。
type PagedSlice[T any] struct {
	pages    [][]T
	pageSize int
	len      int
	lenLast  int
}

// Add 添加一个元素到 PagedSlice 中。
func (slf *PagedSlice[T]) Add(value T) {
	if slf.lenLast == len(slf.pages[len(slf.pages)-1]) {
		slf.pages = append(slf.pages, make([]T, slf.pageSize))
		slf.lenLast = 0
	}

	slf.pages[len(slf.pages)-1][slf.lenLast] = value
	slf.len++
	slf.lenLast++
}

// Get 获取 PagedSlice 中给定索引的元素。
func (slf *PagedSlice[T]) Get(index int) *T {
	if index < 0 || index >= slf.len {
		return nil
	}

	return &slf.pages[index/slf.pageSize][index%slf.pageSize]
}

// Set 设置 PagedSlice 中给定索引的元素。
func (slf *PagedSlice[T]) Set(index int, value T) {
	if index < 0 || index >= slf.len {
		return
	}

	slf.pages[index/slf.pageSize][index%slf.pageSize] = value
}

// Len 返回 PagedSlice 中元素的数量。
func (slf *PagedSlice[T]) Len() int {
	return slf.len
}

// Clear 清空 PagedSlice。
func (slf *PagedSlice[T]) Clear() {
	slf.pages = make([][]T, 0)
	slf.len = 0
	slf.lenLast = 0
}

// Range 迭代 PagedSlice 中的所有元素。
func (slf *PagedSlice[T]) Range(f func(index int, value T) bool) {
	for i := 0; i < slf.len; i++ {
		if !f(i, slf.pages[i/slf.pageSize][i%slf.pageSize]) {
			return
		}
	}
}
