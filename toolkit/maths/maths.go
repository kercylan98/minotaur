package maths

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"math"
)

// Sqrt 传入任意数值类型，返回其 float64 类型的平方根
func Sqrt[T constraints.Number](x T) float64 {
	return math.Sqrt(float64(x))
}

// MinMax 传入任意数值类型，返回其最小值和最大值
func MinMax[T constraints.Number](a, b T) (min, max T) {
	if a < b {
		return a, b
	}
	return b, a
}

// MaxMin 传入任意数值类型，返回其最大值和最小值
func MaxMin[T constraints.Number](a, b T) (max, min T) {
	if a < b {
		return b, a
	}
	return a, b
}

// Clamp 将给定值限制在最小值和最大值之间
func Clamp[T constraints.Number](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
