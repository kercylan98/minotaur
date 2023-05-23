package times

import (
	"time"
)

// CalcNextSec 计算下一个N秒在多少秒之后
func CalcNextSec(sec int) int {
	now := time.Now().Unix()
	next := now + int64(sec) - now%int64(sec)
	return int(next - now)
}

// CalcNextSecWithTime 计算下一个N秒在多少秒之后
func CalcNextSecWithTime(t time.Time, sec int) int {
	now := t.Unix()
	next := now + int64(sec) - now%int64(sec)
	return int(next - now)
}
