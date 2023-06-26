package offset

import "time"

var global *Time

func init() {
	global = NewTime(0)
}

// NewTime 新建一个包含偏移的时间
func NewTime(offset time.Duration) *Time {
	return &Time{offset: offset}
}

// Time 带有偏移量的时间
type Time struct {
	offset time.Duration
}

// SetOffset 设置时间偏移
func (slf *Time) SetOffset(offset time.Duration) {
	slf.offset = offset
}

// Now 获取当前时间偏移后的时间
func (slf *Time) Now() time.Time {
	return time.Now().Add(slf.offset)
}

// Since 获取当前时间偏移后的时间自从 t 以来经过的时间
func (slf *Time) Since(t time.Time) time.Duration {
	return slf.Now().Sub(t)
}

// SetGlobal 设置全局偏移时间
func SetGlobal(offset time.Duration) {
	global.SetOffset(offset)
}

// GetGlobal 获取全局偏移时间
func GetGlobal() *Time {
	return global
}

// Now 获取当前时间偏移后的时间
func Now() time.Time {
	return global.Now()
}

// Since 获取当前时间偏移后的时间自从 t 以来经过的时间
func Since(t time.Time) time.Duration {
	return global.Since(t)
}
