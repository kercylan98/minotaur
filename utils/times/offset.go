package times

import (
	"sync"
	"time"
)

var sourceLocation = time.Local
var offsetLock sync.Mutex
var currentOffsetDuration time.Duration

// SetGlobalTimeOffset 设置全局时间偏移量
func SetGlobalTimeOffset(offset time.Duration) {
	offsetLock.Lock()
	defer offsetLock.Unlock()
	time.Local = sourceLocation
	_, currentOffset := time.Now().Zone()
	newOffset := currentOffset + int(offset.Seconds())
	location := time.FixedZone("OFFSET", newOffset)
	time.Local = location
	currentOffsetDuration = offset
}

// NowByNotOffset 获取未偏移的当前时间
func NowByNotOffset() time.Time {
	offsetLock.Lock()
	defer offsetLock.Unlock()
	offset := time.Local
	time.Local = sourceLocation
	now := time.Now()
	time.Local = offset
	return now
}

// GetGlobalTimeOffset 获取全局时间偏移量
func GetGlobalTimeOffset() time.Duration {
	offsetLock.Lock()
	defer offsetLock.Unlock()
	return currentOffsetDuration
}

// ResetGlobalTimeOffset 重置全局时间偏移量
func ResetGlobalTimeOffset() {
	offsetLock.Lock()
	defer offsetLock.Unlock()
	time.Local = sourceLocation
	currentOffsetDuration = 0
}
