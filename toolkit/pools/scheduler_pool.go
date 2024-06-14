package pools

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"sync"
	"time"
)

// NewSchedulerPool 创建一个 SchedulerPool 对象
//   - capacity: 调度器池容量
//   - tick: 时间轮刻度
//   - wheelSize: 时间轮大小
//   - hungry: 是否饥饿模式，饥饿模式下，调度器池默认将创建 capacity 指定数量的调度器
func NewSchedulerPool(capacity int, tick time.Duration, wheelSize int64, hungry ...bool) *SchedulerPool {
	if capacity <= 0 {
		panic(fmt.Errorf("capacity must be greater than 0, but got %d", capacity))
	}
	pool := &SchedulerPool{
		schedulers: make([]*chrono.Scheduler, 0, capacity),
		tick:       tick,
		wheelSize:  wheelSize,
	}

	if len(hungry) > 0 && hungry[0] {
		for i := 0; i < capacity; i++ {
			pool.schedulers = append(pool.schedulers, chrono.NewScheduler(tick, wheelSize))
		}
	}

	return pool
}

// SchedulerPool 是一个线程安全的 chrono.Scheduler 对象池
type SchedulerPool struct {
	schedulers []*chrono.Scheduler
	rw         sync.RWMutex
	tick       time.Duration
	wheelSize  int64
}

// Get 获取一个调度器
func (p *SchedulerPool) Get() *chrono.Scheduler {
	p.rw.Lock()

	var scheduler *chrono.Scheduler
	if len(p.schedulers) == 0 {
		scheduler = chrono.NewScheduler(p.tick, p.wheelSize)
	} else {
		scheduler = p.schedulers[0]
		p.schedulers = p.schedulers[1:]
	}

	p.rw.Unlock()
	return scheduler
}

// Put 将使用完成的调度器放回缓冲区，如果缓冲区已满，则关闭调度器并释放资源
func (p *SchedulerPool) Put(scheduler *chrono.Scheduler) {
	scheduler.Clear()

	p.rw.Lock()
	defer p.rw.Unlock()

	if len(p.schedulers) < cap(p.schedulers) {
		p.schedulers = append(p.schedulers, scheduler)
	} else {
		scheduler.Close()
	}
}
