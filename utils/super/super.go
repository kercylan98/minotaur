package super

import "time"

var launchTime = time.Now()

// LaunchTime 获取程序启动时间
func LaunchTime() time.Time {
	return launchTime
}
