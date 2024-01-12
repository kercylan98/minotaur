package listings

// NewMatrix 创建一个新的 Matrix 实例。
func NewMatrix[V any](dimensions ...int) *Matrix[V] {
	total := 1
	for _, dim := range dimensions {
		total *= dim
	}
	return &Matrix[V]{
		dimensions: dimensions,
		data:       make([]V, total),
	}
}

type Matrix[V any] struct {
	dimensions []int // 维度大小的切片
	data       []V   // 存储矩阵数据的一维切片
}

// Get 获取矩阵中给定索引的元素。
func (slf *Matrix[V]) Get(index ...int) *V {
	if len(index) != len(slf.dimensions) {
		return nil
	}

	var offset = 0
	for i, dim := range slf.dimensions {
		if index[i] < 0 || index[i] >= dim {
			return nil
		}
		offset = offset*dim + index[i]
	}
	return &slf.data[offset]
}

// Set 设置矩阵中给定索引的元素。
func (slf *Matrix[V]) Set(index []int, value V) {
	if len(index) != len(slf.dimensions) {
		return
	}

	var offset = 0
	for i, dim := range slf.dimensions {
		if index[i] < 0 || index[i] >= dim {
			return
		}
		offset = offset*dim + index[i]
	}
	slf.data[offset] = value
}

// Dimensions 返回矩阵的维度大小。
func (slf *Matrix[V]) Dimensions() []int {
	return slf.dimensions
}

// Clear 清空矩阵。
func (slf *Matrix[V]) Clear() {
	slf.data = make([]V, len(slf.data))
}
