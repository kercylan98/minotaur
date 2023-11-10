package times

import (
	"sync"
	"time"
)

var sourceLocation = time.Local
var offsetLock sync.Mutex

// SetGlobalTimeOffset 设置全局时间偏移量
func SetGlobalTimeOffset(offset time.Duration) {
	offsetLock.Lock()
	defer offsetLock.Unlock()
	time.Local = sourceLocation
	_, currentOffset := time.Now().Zone()
	newOffset := currentOffset + int(offset.Seconds())
	location := time.FixedZone("OFFSET", newOffset)
	time.Local = location
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
