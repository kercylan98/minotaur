package random

import (
	"math/rand"
	"time"
)

// Int64 返回一个介于min和max之间的int64类型的随机数。
func Int64(min int64, max int64) int64 {
	return min + rand.Int63n(max+1-min)
}

// Int 返回一个介于min和max之间的的int类型的随机数。
func Int(min int, max int) int {
	return int(Int64(int64(min), int64(max)))
}

// Duration 返回一个介于min和max之间的的Duration类型的随机数。
func Duration(min int64, max int64) time.Duration {
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
