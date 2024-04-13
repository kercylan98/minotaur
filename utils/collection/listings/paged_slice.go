package listings

// NewPagedSlice 创建一个新的 PagedSlice 实例。
func NewPagedSlice[T any](pageSize int) *PagedSlice[T] {
	return &PagedSlice[T]{
		pageSize: pageSize,
	}
}

// PagedSlice 是一个高效的动态数组，它通过分页管理内存并减少频繁的内存分配来提高性能。
type PagedSlice[T any] struct {
	pages    [][]T
	len      int
	lenLast  int
	pageSize int
}

// Add 添加一个元素到 PagedSlice 中。
func (p *PagedSlice[T]) Add(value T) {
	if p.len == 0 || p.lenLast == p.pageSize {
		p.pages = append(p.pages, make([]T, p.pageSize))
		p.lenLast = 0
	}
	p.pages[len(p.pages)-1][p.lenLast] = value
	p.len++
	p.lenLast++
}

// Get 获取 PagedSlice 中给定索引的元素。
func (p *PagedSlice[T]) Get(index int) *T {
	return &p.pages[index/p.pageSize][index%p.pageSize]
}

// Set 设置 PagedSlice 中给定索引的元素。
func (p *PagedSlice[T]) Set(index int, value T) {
	p.pages[index/p.pageSize][index%p.pageSize] = value
}

// Len 返回 PagedSlice 中元素的数量。
func (p *PagedSlice[T]) Len() int {
	return p.len
}

// Clear 清空 PagedSlice。
func (p *PagedSlice[T]) Clear() {
	p.pages = make([][]T, 0)
	p.len = 0
	p.lenLast = 0
}

// Range 迭代 PagedSlice 中的所有元素。
func (p *PagedSlice[T]) Range(f func(index int, value T) bool) {
	for i := 0; i < p.len; i++ {
		if !f(i, p.pages[i/p.pageSize][i%p.pageSize]) {
			return
		}
	}
}
