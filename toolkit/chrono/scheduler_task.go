package chrono

import (
	"github.com/RussellLuo/timingwheel"
	"github.com/gorhill/cronexpr"
	"reflect"
	"sync"
	"time"
)

// schedulerTask 调度器
type schedulerTask struct {
	lock      sync.RWMutex
	scheduler *Scheduler           // 任务所属的调度器
	timer     *timingwheel.Timer   // 任务执行定时器
	name      string               // 任务名称
	after     time.Duration        // 任务首次执行延迟
	interval  time.Duration        // 任务执行间隔
	function  reflect.Value        // 任务执行函数
	args      []reflect.Value      // 任务执行参数
	expr      *cronexpr.Expression // 任务执行时间表达式

	total   int  // 任务执行次数
	trigger int  // 任务已执行次数
	kill    bool // 任务是否已关闭
}

// Name 获取任务名称
func (t *schedulerTask) Name() string {
	return t.name
}

// Next 获取任务下一次执行的时间
func (t *schedulerTask) Next(prev time.Time) time.Time {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.kill || (t.expr != nil && t.total > 0 && t.trigger > t.total) {
		return time.Time{}
	}
	if t.expr != nil {
		next := t.expr.Next(prev)
		return next
	}
	if t.trigger == 0 {
		t.trigger++
		return prev.Add(t.after)
	}
	t.trigger++
	return prev.Add(t.interval)
}

func (t *schedulerTask) caller() {
	t.lock.Lock()

	if t.kill {
		t.lock.Unlock()
		return
	}

	if t.total > 0 && t.trigger > t.total {
		t.lock.Unlock()
		t.scheduler.UnregisterTask(t.name)
	} else {
		t.lock.Unlock()
	}
	t.function.Call(t.args)
}

func (t *schedulerTask) close() {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.kill {
		return
	}
	t.kill = true
	if t.total <= 0 || t.trigger < t.total {
		t.timer.Stop()
	}
}
