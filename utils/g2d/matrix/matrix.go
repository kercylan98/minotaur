package matrix

import "github.com/kercylan98/minotaur/utils/g2d"

// NewMatrix 生成特定宽高的二维矩阵
//   - 虽然提供了通过x、y坐标的操作函数，但是建议无论如何使用pos进行处理
//   - 该矩阵为XY，而非YX
func NewMatrix[T any](width, height int) *Matrix[T] {
	matrix := &Matrix[T]{
		w: width, h: height,
	}
	matrix.m = make([]T, width*height)
	return matrix
}

// Matrix 二维矩阵
type Matrix[T any] struct {
	w, h int
	m    []T
}

// GetWidth 获取二维矩阵宽度
func (slf *Matrix[T]) GetWidth() int {
	return slf.w
}

// GetHeight 获取二维矩阵高度
func (slf *Matrix[T]) GetHeight() int {
	return slf.h
}

// GetWidth2Height 获取二维矩阵的宽度和高度
func (slf *Matrix[T]) GetWidth2Height() (width, height int) {
	return slf.w, slf.h
}

// GetMatrix 获取二维矩阵
//   - 通常建议使用 GetMatrixWithPos 进行处理这样将拥有更高的效率
func (slf *Matrix[T]) GetMatrix() [][]T {
	var result = make([][]T, slf.w)
	for x := 0; x < slf.w; x++ {
		ys := make([]T, slf.h)
		for y := 0; y < slf.h; y++ {
			ys[y] = slf.m[g2d.PositionToInt(slf.w, x, y)]
		}
		result[x] = ys
	}
	return result
}

// GetMatrixWithPos 获取顺序的矩阵
func (slf *Matrix[T]) GetMatrixWithPos() []T {
	return slf.m
}

// Get 获取特定坐标的内容
func (slf *Matrix[T]) Get(x, y int) (value T) {
	return slf.m[g2d.PositionToInt(slf.w, x, y)]
}

// GetWithPos 获取特定坐标的内容
func (slf *Matrix[T]) GetWithPos(pos int) (value T) {
	return slf.m[pos]
}

// Set 设置特定坐标的内容
func (slf *Matrix[T]) Set(x, y int, data T) {
	slf.m[g2d.PositionToInt(slf.w, x, y)] = data
}

// SetWithPos 设置特定坐标的内容
func (slf *Matrix[T]) SetWithPos(pos int, data T) {
	slf.m[pos] = data
}

// Swap 交换两个位置的内容
func (slf *Matrix[T]) Swap(x1, y1, x2, y2 int) {
	a, b := slf.Get(x1, y1), slf.Get(x2, y2)
	slf.m[g2d.PositionToInt(slf.w, x1, y1)], slf.m[g2d.PositionToInt(slf.w, x2, y2)] = b, a
}

// SwapWithPos 交换两个位置的内容
func (slf *Matrix[T]) SwapWithPos(pos1, pos2 int) {
	a, b := slf.m[pos1], slf.m[pos2]
	slf.m[pos1], slf.m[pos2] = b, a
}

// TrySwap 尝试交换两个位置的内容，交换后不满足表达式时进行撤销
func (slf *Matrix[T]) TrySwap(x1, y1, x2, y2 int, expressionHandle func(matrix *Matrix[T]) bool) {
	pos1 := g2d.PositionToInt(slf.w, x1, y1)
	pos2 := g2d.PositionToInt(slf.w, x2, y2)
	a, b := slf.Get(x1, y1), slf.Get(x2, y2)
	slf.m[pos1], slf.m[pos2] = b, a
	if !expressionHandle(slf) {
		slf.m[pos1], slf.m[pos2] = a, b
	}
}

// TrySwapWithPos 尝试交换两个位置的内容，交换后不满足表达式时进行撤销
func (slf *Matrix[T]) TrySwapWithPos(pos1, pos2 int, expressionHandle func(matrix *Matrix[T]) bool) {
	a, b := slf.m[pos1], slf.m[pos2]
	slf.m[pos1], slf.m[pos2] = b, a
	if !expressionHandle(slf) {
		slf.m[pos1], slf.m[pos2] = a, b
	}
}

// FillFull 根据提供的生成器填充整个矩阵
func (slf *Matrix[T]) FillFull(generateHandle func(x, y int) T) {
	for x := 0; x < slf.w; x++ {
		for y := 0; y < slf.h; y++ {
			slf.m[g2d.PositionToInt(slf.w, x, y)] = generateHandle(x, y)
		}
	}
}

// FillFullWithPos 根据提供的生成器填充整个矩阵
func (slf *Matrix[T]) FillFullWithPos(generateHandle func(pos int) T) {
	for pos := 0; pos < len(slf.m); pos++ {
		slf.m[pos] = generateHandle(pos)
	}
}
