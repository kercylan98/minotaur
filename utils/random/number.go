package random

import (
	"math/rand"
	"time"
)

// Int64 返回一个介于min和max之间的int64类型的随机数。
func Int64(min int64, max int64) int64 {
	if min == max {
		return min
	}
	return min + rand.Int63n(max+1-min)
}

// Int32 返回一个介于min和max之间的的int32类型的随机数。
func Int32(min int32, max int32) int32 {
	if min == max {
		return min
	}
	return int32(Int64(int64(min), int64(max)))
}

// Float32Range 返回一个介于min和max之间的的float32类型的随机数。
func Float32Range(min float32, max float32) float32 {
	if min == max {
		return min
	}
	return min + rand.Float32()*(max-min)
}

// Float64Range 返回一个介于min和max之间的的float64类型的随机数。
func Float64Range(min float64, max float64) float64 {
	if min == max {
		return min
	}
	return min + rand.Float64()*(max-min)
}

// Int 返回一个介于min和max之间的的int类型的随机数。
func Int(min int, max int) int {
	if min == max {
		return min
	}
	return int(Int64(int64(min), int64(max)))
}

// Duration 返回一个介于min和max之间的的Duration类型的随机数。
func Duration(min int64, max int64) time.Duration {
	if min == max {
		return time.Duration(min)
	}
	return time.Duration(Int64(min, max))
}

// Float64 返回一个0~1的浮点数
func Float64() float64 {
	return rand.Float64()
}

// Float32 返回一个0~1的浮点数
func Float32() float32 {
	return rand.Float32()
}

// IntN 返回一个0~n的整数
func IntN(n int) int {
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

// Bool 返回一个随机的布尔值
func Bool() bool {
	return rand.Intn(2) == 1
}
