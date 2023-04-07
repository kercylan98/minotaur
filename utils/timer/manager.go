package timer

import (
	"reflect"
	"sync"
	"time"

	"github.com/RussellLuo/timingwheel"
)

// Manager 管理器
type Manager struct {
	timer  *Timer
	wheel  *timingwheel.TimingWheel
	timers map[string]*Scheduler
	lock   sync.RWMutex
}

// Release 释放管理器，并将管理器重新放回 Timer 池中
func (slf *Manager) Release() {
	slf.timer.lock.Lock()
	defer slf.timer.lock.Unlock()

	slf.lock.Lock()

	for name, scheduler := range slf.timers {
		scheduler.close()
		delete(slf.timers, name)
	}

	slf.lock.Unlock()

	slf.timer.Managers = append(slf.timer.Managers, slf)
}

// StopTimer 停止特定名称的调度器
func (slf *Manager) StopTimer(name string) {
	slf.lock.Lock()
	defer slf.lock.Unlock()

	if s, ok := slf.timers[name]; ok {
		s.close()
		delete(slf.timers, name)
	}
}

// IsStopped 特定名称的调度器是否已停止
func (slf *Manager) IsStopped(name string) bool {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	if s, ok := slf.timers[name]; ok {
		return s.isClosed()
	}
	return true
}

// GetSchedulers 获取所有调度器名称
func (slf *Manager) GetSchedulers() []string {
	slf.lock.RLock()
	defer slf.lock.RUnlock()
	names := make([]string, 0, len(slf.timers))
	for name := range slf.timers {
		names = append(names, name)
	}
	return names
}

// After 设置一个在特定时间后运行一次的调度器
func (slf *Manager) After(name string, after time.Duration, handleFunc interface{}, args ...interface{}) {
	slf.Loop(name, after, timingWheelTick, 1, handleFunc, args...)
}

// Loop 设置一个在特定时间后仿佛运行的调度器
func (slf *Manager) Loop(name string, after, interval time.Duration, times int, handleFunc interface{}, args ...interface{}) {
	slf.StopTimer(name)

	if after < timingWheelTick {
		after = timingWheelTick
	}
	if interval < timingWheelTick {
		interval = timingWheelTick
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
		manager:  slf,
	}

	slf.lock.Lock()
	slf.timers[name] = scheduler
	scheduler.timer = slf.wheel.ScheduleFunc(scheduler, scheduler.caller)
	slf.lock.Unlock()
}

