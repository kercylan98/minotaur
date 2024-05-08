package maths

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"math"
)

// Sqrt 传入任意数值类型，返回其 float64 类型的平方根
func Sqrt[T constraints.Number](x T) float64 {
	return math.Sqrt(float64(x))
}
