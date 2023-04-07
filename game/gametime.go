package game

import "time"

// Time 带有偏移的游戏时间
type Time struct {
	Offset time.Duration
}

// Now 获取包含偏移的当前时间
func (slf *Time) Now() time.Time {
	t := time.Now()
	if slf.Offset == 0 {
		return t
	}
	return t.Add(slf.Offset)
}

// Since 获取两个时间之间的差
func (slf *Time) Since(t time.Time) time.Duration {
	return slf.Now().Sub(t)
}
