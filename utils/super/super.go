package super

import (
	"os"
	"time"
)

var launchTime = time.Now()

// LaunchTime 获取程序启动时间
func LaunchTime() time.Time {
	return launchTime
}

// Hostname 获取主机名
func Hostname() string {
	return os.Getenv("HOSTNAME")
}
