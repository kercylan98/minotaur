package toolkit

import (
	"os"
	"time"
)

var launchTime = time.Now()
var pid int
var hostname string

func init() {
	pid = os.Getpid()
	hostname, _ = os.Hostname()
}

// LaunchTime 获取程序启动时间
func LaunchTime() time.Time {
	return launchTime
}

// Hostname 获取主机名
func Hostname() string {
	return hostname
}

// PID 获取进程 PID
func PID() int {
	return pid
}
