package maths

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"math"
	"sort"
)

const (
	DefaultTolerance = 0.0001 // 默认误差范围
	Zero             = 0      // 零
)

// GetDefaultTolerance 获取默认误差范围
func GetDefaultTolerance() float64 {
	return DefaultTolerance
}

// Pow 整数幂运算
func Pow(a, n int) int {
	if a == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	var result = 1
	factor := a
	for n != 0 {
		if n&1 != 0 {
			// 当前位是 1，需要乘进去
			result *= factor
		}
		factor *= factor
		n = n >> 1
	}
	return result
}

// PowInt64 整数幂运算
func PowInt64(a, n int64) int64 {
	if a == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	var result int64 = 1
	factor := a
	for n != 0 {
		if n&1 != 0 {
			// 当前位是 1，需要乘进去
			result *= factor
		}
		factor *= factor
		n = n >> 1
	}
	return result
}

// Min 返回两个数之中较小的值
func Min[V generic.Number](a, b V) V {
	if a < b {
		return a
	}
	return b
}

// Max 返回两个数之中较大的值
func Max[V generic.Number](a, b V) V {
	if a > b {
		return a
	}
	return b
}

// MinMax 将两个数按照较小的和较大的顺序进行返回
func MinMax[V generic.Number](a, b V) (min, max V) {
	if a < b {
		return a, b
	}
	return b, a
}

// MaxMin 将两个数按照较大的和较小的顺序进行返回
func MaxMin[V generic.Number](a, b V) (max, min V) {
	if a > b {
		return a, b
	}
	return b, a
}

// Clamp 将给定值限制在最小值和最大值之间
func Clamp[V generic.Number](value, min, max V) V {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Tolerance 检查两个值是否在一个误差范围内
func Tolerance[V generic.Number](value1, value2, tolerance V) bool {
	return V(math.Abs(float64(value1-value2))) <= tolerance
}

// Merge 通过一个参考值合并两个数字
func Merge[V generic.SignedNumber](refer, a, b V) V {
	return b*refer + a
}

// UnMerge 通过一个参考值取消合并的两个数字
func UnMerge[V generic.SignedNumber](refer, num V) (a, b V) {
	a = V(math.Mod(float64(num), float64(refer)))
	b = num / refer
	return a, b
}

// MergeToInt64 将两个数字合并为一个 int64 数字
func MergeToInt64[V generic.SignedNumber](v1 V, v2 V) int64 {
	a, b := int64(v1), int64(v2)
	mergedNumber := (a << 32) | (b & 0xFFFFFFFF)
	return mergedNumber
}

// UnMergeInt64 将一个 int64 数字拆分为两个数字
func UnMergeInt64[V generic.SignedNumber](n int64) (V, V) {
	a := n >> 32
	b := n & 0xFFFFFFFF
	return V(a), V(b)
}

// ToContinuous 将一组非连续的数字转换为从1开始的连续数字
//   - 返回值是一个 map，key 是从 1 开始的连续数字，value 是原始数字
func ToContinuous[V generic.Integer](nums []V) map[V]V {
	if len(nums) == 0 {
		return nil
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var result = make(map[V]V)
	for i, num := range nums {
		result[V(i+1)] = num
	}
	return result
}

// CountDigits 接收一个整数 num 作为输入，并返回该数字的位数
func CountDigits[V generic.Number](num V) int {
	// 处理0的特殊情况
	if num == 0 {
		return 1
	}

	// 取绝对值
	absNum := V(math.Abs(float64(num)))

	// 计算位数
	count := 0
	for absNum > 0 {
		absNum /= 10
		count++
	}

	return count
}

// GetDigitValue 接收一个整数 num 和一个表示目标位数的整数 digit 作为输入，并返
// 回数字 num 在指定位数上的数值。我们使用 math.Abs() 函数获取 num 的绝对值，并通
// 过除以10的操作将 num 移动到目标位数上。然后，通过取余运算得到位数上的数值
func GetDigitValue(num int64, digit int) int64 {
	absNum := int64(math.Abs(float64(num)))
	for i := 0; i < digit; i++ {
		absNum /= 10
	}
	return absNum % 10
}

// JoinNumbers 将一组数字连接起来
func JoinNumbers[V generic.Number](num1 V, n ...V) V {
	var v = num1
	for _, v2 := range n {
		base := V(math.Pow10(int(math.Log10(float64(v2))) + 1))
		v = v*base + v2
	}
	return v
}
