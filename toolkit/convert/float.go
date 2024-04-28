package convert

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"strconv"
)

// FloatToBoolean 将 float32/float64 转换为 bool 类型
func FloatToBoolean[F constraints.Float](f F) bool {
	return f != 0
}

// FloatToInt 将 float32/float64 转换为 int 类型
func FloatToInt[F constraints.Float, I constraints.Int](f F) I {
	return I(f)
}

// Float32ToString 将 float32 转换为 string 类型
func Float32ToString(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', -1, 32)
}

// Float64ToString 将 float64 转换为 string 类型
func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

// FloatToIntRound 将 float32/float64 转换为 int 类型，四舍五入
func FloatToIntRound[F constraints.Float, I constraints.Int](f F) I {
	return I(f + F(0.5))
}

// FloatToIntCeil 将 float32/float64 转换为 int 类型，向上取整
func FloatToIntCeil[F constraints.Float, I constraints.Int](f F) I {
	return I(f + F(1))
}

// FloatToIntFloor 将 float32/float64 转换为 int 类型，向下取整
func FloatToIntFloor[F constraints.Float, I constraints.Int](f F) I {
	return FloatToInt[F, I](f)
}
