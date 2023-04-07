package timer

import "time"

var (
	// 时间轮一刻度长度
	timingWheelTick = time.Millisecond * 10
)
