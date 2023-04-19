package g2d

import (
	"minotaur/utils/g2d/matrix"
)

// NewMatrix 生成特定宽高的二维矩阵
func NewMatrix[T any](width, height int) *matrix.Matrix[T] {
	return matrix.NewMatrix[T](width, height)
}
