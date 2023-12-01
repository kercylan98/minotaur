package sole

import (
	"strconv"
	"sync/atomic"
)

var autoIncrementUint32 uint32 = 0
var autoIncrementUint64 uint64 = 0
var autoIncrementInt32 int32 = 0
var autoIncrementInt64 int64 = 0
var autoIncrementInt uint64 = 0
var autoIncrementString uint64 = 0

// AutoIncrementUint32 获取一个自增的 uint32 值
func AutoIncrementUint32() uint32 {
	return atomic.AddUint32(&autoIncrementUint32, 1) - 1
}

// AutoIncrementUint64 获取一个自增的 uint64 值
func AutoIncrementUint64() uint64 {
	return atomic.AddUint64(&autoIncrementUint64, 1) - 1
}

// AutoIncrementInt32 获取一个自增的 int32 值
func AutoIncrementInt32() int32 {
	return atomic.AddInt32(&autoIncrementInt32, 1) - 1
}

// AutoIncrementInt64 获取一个自增的 int64 值
func AutoIncrementInt64() int64 {
	return atomic.AddInt64(&autoIncrementInt64, 1) - 1
}

// AutoIncrementInt 获取一个自增的 int 值
func AutoIncrementInt() int {
	num := atomic.AddUint64(&autoIncrementInt, 1)
	result := num % (1 << (strconv.IntSize - 1))
	return int(result)
}

// AutoIncrementString 获取一个自增的字符串
func AutoIncrementString() string {
	return strconv.FormatUint(atomic.AddUint64(&autoIncrementString, 1)-1, 10)
}
