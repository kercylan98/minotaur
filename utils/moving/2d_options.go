package moving

import (
	"errors"
	"github.com/kercylan98/minotaur/utils/generic"
	"time"
)

type TwoDimensionalOption[EID generic.Basic, PosType generic.SignedNumber] func(moving *TwoDimensional[EID, PosType])

// WithTwoDimensionalTimeUnit 通过特定时间单位创建
//   - 默认单位为1毫秒，最小单位也为1毫秒
func WithTwoDimensionalTimeUnit[EID generic.Basic, PosType generic.SignedNumber](duration time.Duration) TwoDimensionalOption[EID, PosType] {
	return func(moving *TwoDimensional[EID, PosType]) {
		if duration < time.Millisecond {
			panic(errors.New("time unit milliseconds minimum"))
		}
		moving.timeUnit = float64(duration)
	}
}

// WithTwoDimensionalIdleWaitTime 通过特定的空闲等待时间创建
//   - 默认情况下在没有新的移动计划时将限制 100毫秒 + 移动间隔事件(默认100毫秒)
func WithTwoDimensionalIdleWaitTime[EID generic.Basic, PosType generic.SignedNumber](duration time.Duration) TwoDimensionalOption[EID, PosType] {
	return func(moving *TwoDimensional[EID, PosType]) {
		moving.idle = duration
	}
}

// WithTwoDimensionalInterval 通过特定的移动间隔时间创建
func WithTwoDimensionalInterval[EID generic.Basic, PosType generic.SignedNumber](duration time.Duration) TwoDimensionalOption[EID, PosType] {
	return func(moving *TwoDimensional[EID, PosType]) {
		moving.interval = duration
	}
}
