package timer

import (
	"github.com/gorhill/cronexpr"
	"reflect"
	"sync"
	"time"

	"github.com/RussellLuo/timingwheel"
)

// Ticker 定时器
type Ticker struct {
	timer  *Pool
	wheel  *timingwheel.TimingWheel
	timers map[string]*Scheduler
	lock   sync.RWMutex

	handler func(name string, caller func())
	mark    string
}

// Mark 获取定时器的标记
//   - 通常用于鉴别定时器来源
func (slf *Ticker) Mark() string {
	return slf.mark
}

// Release 释放定时器，并将定时器重新放回 Pool 池中
func (slf *Ticker) Release() {
	slf.timer.lock.Lock()
	defer slf.timer.lock.Unlock()

	slf.lock.Lock()
	slf.mark = ""
	slf.handler = nil
	for name, scheduler := range slf.timers {
		scheduler.close()
		delete(slf.timers, name)
	}
	slf.lock.Unlock()

	if len(slf.timer.tickers) < tickerPoolSize && !slf.timer.closed {
		slf.timer.tickers = append(slf.timer.tickers, slf)
	} else {
		slf.wheel.Stop()
	}
}

// StopTimer 停止特定名称的调度器
func (slf *Ticker) StopTimer(name string) {
	slf.lock.Lock()
	defer slf.lock.Unlock()

	if s, ok := slf.timers[name]; ok {
		s.close()
		delete(slf.timers, name)
	}
}

// IsStopped 特定名称的调度器是否已停止
func (slf *Ticker) IsStopped(name string) bool {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	if s, ok := slf.timers[name]; ok {
		return s.isClosed()
	}
	return true
}

// GetSchedulers 获取所有调度器名称
func (slf *Ticker) GetSchedulers() []string {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	names := make([]string, 0, len(slf.timers))
	for name := range slf.timers {
		names = append(names, name)
	}
	return names
}

// Cron 通过 cron 表达式设置一个调度器，当 cron 表达式错误时，将会引发 panic
func (slf *Ticker) Cron(name, expression string, handleFunc interface{}, args ...interface{}) {
	expr := cronexpr.MustParse(expression)
	slf.loop(name, 0, 0, expr, 0, handleFunc, args...)
}

// CronByInstantly 与 Cron 相同，但是会立即执行一次
func (slf *Ticker) CronByInstantly(name, expression string, handleFunc interface{}, args ...interface{}) {
	func(name, expression string, handleFunc interface{}, args ...interface{}) {
		var values = make([]reflect.Value, len(args))
		for i, v := range args {
			values[i] = reflect.ValueOf(v)
		}
		f := reflect.ValueOf(handleFunc)
		slf.lock.RLock()
		defer slf.lock.RUnlock()
		if slf.handler != nil {
			slf.handler(name, func() {
				f.Call(values)
			})
		} else {
			f.Call(values)
		}
	}(name, expression, handleFunc, args...)

	slf.Cron(name, expression, handleFunc, args...)
}

// After 设置一个在特定时间后运行一次的调度器
func (slf *Ticker) After(name string, after time.Duration, handleFunc interface{}, args ...interface{}) {
	slf.loop(name, after, timingWheelTick, nil, 1, handleFunc, args...)
}

// Loop 设置一个在特定时间后反复运行的调度器
func (slf *Ticker) Loop(name string, after, interval time.Duration, times int, handleFunc interface{}, args ...interface{}) {
	slf.loop(name, after, interval, nil, times, handleFunc, args...)
}

// Loop 设置一个在特定时间后反复运行的调度器
func (slf *Ticker) loop(name string, after, interval time.Duration, expr *cronexpr.Expression, times int, handleFunc interface{}, args ...interface{}) {
	slf.StopTimer(name)

	if expr == nil {
		if after < timingWheelTick {
			after = timingWheelTick
		}
		if interval < timingWheelTick {
			interval = timingWheelTick
		}
	}

	var values = make([]reflect.Value, len(args))
	for i, v := range args {
		values[i] = reflect.ValueOf(v)
	}

	scheduler := &Scheduler{
		name:     name,
		after:    after,
		interval: interval,
		total:    times,
		cbFunc:   reflect.ValueOf(handleFunc),
		cbArgs:   values,
		ticker:   slf,
		expr:     expr,
	}

	slf.lock.Lock()
	slf.timers[name] = scheduler
	if slf.handler != nil {
		scheduler.timer = slf.wheel.ScheduleFunc(scheduler, func() {
			slf.lock.RLock()
			handler := slf.handler
			slf.lock.RUnlock()
			if handler != nil {
				slf.handler(scheduler.Name(), scheduler.Caller)
			}
		})
	} else {
		scheduler.timer = slf.wheel.ScheduleFunc(scheduler, scheduler.Caller)
	}
	slf.lock.Unlock()
}
