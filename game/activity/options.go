package activity

import (
	"github.com/kercylan98/minotaur/utils/times"
	"time"
)

// NewOptions 创建活动选项
func NewOptions() *Options {
	return new(Options)
}

// initOptions 初始化活动选项
func initOptions(opts ...*Options) *Options {
	var opt *Options
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt == nil {
		opt = NewOptions()
	}
	return opt
}

// Options 活动选项
type Options struct {
	Tl   *times.StateLine[byte] // 活动时间线
	Loop time.Duration          // 活动循环，时间间隔小于等于 0 表示不循环
}

// WithUpcomingTime 设置活动预告时间
func (slf *Options) WithUpcomingTime(t time.Time) *Options {
	if slf.Tl == nil {
		slf.Tl = times.NewStateLine[byte](stateClosed)
	}
	slf.Tl.AddState(stateUpcoming, t)
	return slf
}

// WithStartTime 设置活动开始时间
func (slf *Options) WithStartTime(t time.Time) *Options {
	if slf.Tl == nil {
		slf.Tl = times.NewStateLine[byte](stateClosed)
	}
	slf.Tl.AddState(stateStarted, t)
	return slf
}

// WithEndTime 设置活动结束时间
func (slf *Options) WithEndTime(t time.Time) *Options {
	if slf.Tl == nil {
		slf.Tl = times.NewStateLine[byte](stateClosed)
	}
	slf.Tl.AddState(stateEnded, t)
	return slf
}

// WithExtendedShowTime 设置延长展示时间
func (slf *Options) WithExtendedShowTime(t time.Time) *Options {
	if slf.Tl == nil {
		slf.Tl = times.NewStateLine[byte](stateClosed)
	}
	slf.Tl.AddState(stateExtendedShowEnded, t)
	return slf
}

// WithLoop 设置活动循环，时间间隔小于等于 0 表示不循环
//   - 当活动状态展示结束后，会根据该选项设置的时间间隔重新开始
func (slf *Options) WithLoop(interval time.Duration) *Options {
	slf.Loop = interval
	return slf
}
