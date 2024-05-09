package geometry

import "github.com/kercylan98/minotaur/toolkit/constraints"

// CalcAngleDifference 计算两个极角之间的最小角度差。
// 结果在 -180 到 180 度之间，适用于极角、方位角或其他类似场景。
func CalcAngleDifference[T constraints.Number](a, b T) float64 {
	diff := float64(a) - float64(b)
	if diff > 180 {
		diff -= 360
	} else if diff < -180 {
		diff += 360
	}
	return diff
}
