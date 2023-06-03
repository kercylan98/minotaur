package components

import (
	"errors"
	"time"
)

type Moving2DOption func(moving *Moving2D)

// WithMoving2DTimeUnit 通过特定时间单位创建
//   - 默认单位为1毫秒，最小单位也为1毫秒
func WithMoving2DTimeUnit(duration time.Duration) Moving2DOption {
	return func(moving *Moving2D) {
		if duration < time.Millisecond {
			panic(errors.New("time unit milliseconds minimum"))
		}
		moving.timeUnit = float64(duration)
	}
}

// WithMoving2DIdleWaitTime 通过特定的空闲等待时间创建
//   - 默认情况下在没有新的移动计划时将限制 100毫秒 + 移动间隔事件(默认100毫秒)
func WithMoving2DIdleWaitTime(duration time.Duration) Moving2DOption {
	return func(moving *Moving2D) {
		moving.idle = duration
	}
}

// WithMoving2DInterval 通过特定的移动间隔时间创建
func WithMoving2DInterval(duration time.Duration) Moving2DOption {
	return func(moving *Moving2D) {
		moving.interval = duration
	}
}
