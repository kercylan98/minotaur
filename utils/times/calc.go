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

// CalcNextTimeWithRefer 根据参考时间计算下一个整点时间
//   - 假设当 now 为 14:15:16 ， 参考时间为 10 分钟， 则返回 14:20:00
//   - 假设当 now 为 14:15:16 ， 参考时间为 20 分钟， 则返回 14:20:00
//
// 当 refer 小于 1 分钟时，将会返回当前时间
func CalcNextTimeWithRefer(now time.Time, refer time.Duration) time.Time {
	referInSeconds := int(refer.Minutes()) * 60
	if referInSeconds <= 0 {
		return now
	}

	minutes := now.Minute()
	seconds := now.Second()

	remainder := referInSeconds - (minutes*60+seconds)%referInSeconds
	nextTime := now.Add(time.Duration(remainder) * time.Second)
	nextTime = time.Date(nextTime.Year(), nextTime.Month(), nextTime.Day(), nextTime.Hour(), nextTime.Minute(), 0, 0, nextTime.Location())
	return nextTime
}
