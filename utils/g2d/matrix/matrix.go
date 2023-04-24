package matrix

// NewMatrix 生成特定宽高的二维矩阵
func NewMatrix[T any](width, height int) *Matrix[T] {
	matrix := &Matrix[T]{
		w: width, h: height,
	}
	matrix.m = make([][]T, width)
	for x := 0; x < len(matrix.m); x++ {
		matrix.m[x] = make([]T, height)
	}
	return matrix
}

// Matrix 二维矩阵
type Matrix[T any] struct {
	w, h int
	m    [][]T
}

// GetWidth 获取二维矩阵宽度
func (slf *Matrix[T]) GetWidth() int {
	return slf.w
}

// GetHeight 获取二维矩阵高度
func (slf *Matrix[T]) GetHeight() int {
	return slf.h
}

// GetMatrix 获取二维矩阵
func (slf *Matrix[T]) GetMatrix() [][]T {
	return slf.m
}

// Get 获取特定坐标的内容
func (slf *Matrix[T]) Get(x, y int) T {
	return slf.m[x][y]
}

// Set 设置特定坐标的内容
func (slf *Matrix[T]) Set(x, y int, data T) {
	slf.m[x][y] = data
}

// Swap 交换两个位置的内容
func (slf *Matrix[T]) Swap(x1, y1, x2, y2 int) {
	a, b := slf.Get(x1, y1), slf.Get(x2, y2)
	slf.m[x1][y1], slf.m[x2][y2] = b, a
}

// TrySwap 尝试交换两个位置的内容，交换后不满足表达式时进行撤销
func (slf *Matrix[T]) TrySwap(x1, y1, x2, y2 int, expressionHandle func(matrix *Matrix[T]) bool) {
	a, b := slf.Get(x1, y1), slf.Get(x2, y2)
	slf.m[x1][y1], slf.m[x2][y2] = b, a
	if !expressionHandle(slf) {
		slf.m[x1][y1], slf.m[x2][y2] = a, b
	}
}
