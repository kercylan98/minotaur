package random

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"math/rand"
	"time"
)

// Int 返回一个包含 min 和 max 及之间的 int 类型的随机数
func Int[I constraints.Int](min, max I) I {
	if min > max {
		panic("min > max")
	}
	if min == max {
		return min
	}
	return min + I(rand.Intn(int(max-min+1)))
}

// IntN 返回一个介于 0 和 n 之间的 int 类型的随机数
func IntN[I constraints.Int](n I) I {
	return I(rand.Intn(int(n)))
}

// Float 返回一个介于 min 和 max 之间的 float 类型的随机数
func Float[F constraints.Float](min, max F) F {
	if min == max {
		return min
	}
	return min + F(rand.Float64())*(max-min)
}

// FloatN 返回一个介于 0 和 n 之间的 float 类型的随机数
func FloatN[F constraints.Float](n F) F {
	return F(rand.Float64()) * n
}

// Duration 返回一个介于 min 和 max 之间的 time.Duration 类型的随机数
func Duration(min time.Duration, max time.Duration) time.Duration {
	return time.Duration(Int(int64(min), int64(max)))
}

// Bool 返回一个随机的布尔值
func Bool() bool {
	return Int(0, 1) == 1
}
