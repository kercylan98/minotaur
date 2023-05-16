package timer

import (
	"reflect"
	"sync"
	"time"

	"github.com/RussellLuo/timingwheel"
)

// Scheduler 调度器
type Scheduler struct {
	name     string
	after    time.Duration
	interval time.Duration

	total   int
	trigger int
	kill    bool

	cbFunc reflect.Value
	cbArgs []reflect.Value

	timer *timingwheel.Timer

	ticker *Ticker

	lock sync.RWMutex
}

// Name 获取调度器名称
func (slf *Scheduler) Name() string {
	return slf.name
}

// Next 获取下一次执行的时间
func (slf *Scheduler) Next(prev time.Time) time.Time {
	slf.lock.RLock()
	defer slf.lock.RUnlock()

	if slf.kill || (slf.total > 0 && slf.trigger >= slf.total) {
		return time.Time{}
	}
	if slf.trigger == 0 {
		return prev.Add(slf.after)
	}
	return prev.Add(slf.interval)
}

// Caller 可由外部发起调用的执行函数
func (slf *Scheduler) Caller() {
	slf.lock.Lock()

	if slf.kill {
		slf.lock.Unlock()
		return
	}

	slf.trigger++
	if slf.total > 0 && slf.trigger >= slf.total {
		slf.lock.Unlock()
		slf.ticker.StopTimer(slf.name)
	} else {
		slf.lock.Unlock()
	}
	slf.cbFunc.Call(slf.cbArgs)
}

// isClosed 检查调度器是否已关闭
func (slf *Scheduler) isClosed() bool {
	slf.lock.RLock()
	defer slf.lock.RUnlock()

	return slf.kill
}

// close 关闭调度器
func (slf *Scheduler) close() {
	slf.lock.Lock()
	defer slf.lock.Unlock()

	if slf.kill {
		return
	}
	slf.kill = true
	if slf.total <= 0 || slf.trigger < slf.total {
		slf.timer.Stop()
	}
}
