package chrono

import (
	"fmt"
	"sync"
	"time"
)

const (
	SchedulerPoolDefaultSize      = 96
	SchedulerPoolDefaultTick      = time.Millisecond * 10
	SchedulerPoolDefaultWheelSize = 10
)

// NewDefaultSchedulerPool 创建一个默认参数的并发安全的时间调度器对象池
//   - size: SchedulerPoolDefaultSize
//   - tick: SchedulerPoolDefaultTick
//   - wheelSize: SchedulerPoolDefaultWheelSize
func NewDefaultSchedulerPool() *SchedulerPool {
	scheduler, err := NewSchedulerPool(SchedulerPoolDefaultSize, SchedulerPoolDefaultTick, SchedulerPoolDefaultWheelSize)
	if err != nil {
		panic(err) // 该错误不应该发生，用于在参数或实现变更后的提示
	}
	return scheduler
}

// NewSchedulerPool 创建一个并发安全的时间调度器对象池
func NewSchedulerPool(size int, tick time.Duration, wheelSize int64) (*SchedulerPool, error) {
	if size <= 0 {
		return nil, fmt.Errorf("scheduler pool size must greater than 0, got: %d", size)
	}
	if wheelSize <= 0 {
		return nil, fmt.Errorf("scheduler pool wheelSize must greater than 0, got: %d", size)
	}
	return &SchedulerPool{
		size:       size,
		tick:       tick,
		wheelSize:  wheelSize,
		generation: 1,
	}, nil
}

// SchedulerPool 并发安全的时间调度器对象池
type SchedulerPool struct {
	schedulers []*Scheduler                     // 池中维护的时间调度器
	lock       sync.RWMutex                     // 用于并发安全的锁
	tick       time.Duration                    // 时间周期
	wheelSize  int64                            // 时间轮尺寸
	size       int                              // 池大小，控制了池中时间调度器的数量
	generation int64                            // 池的代数
	executor   func(name string, caller func()) // 任务执行器
}

// SetExecutor 设置该事件调度器对象池中整体的任务执行器
func (p *SchedulerPool) SetExecutor(executor func(name string, caller func())) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.executor = executor
}

// SetSize 改变时间调度器对象池的大小，当传入的大小小于或等于 0 时，将会返回错误，并且不会发生任何改变
//   - 设置时间调度器对象池的大小可以在运行时动态调整，但是需要注意的是，调整后的大小不会影响已经产生的 Scheduler
//   - 已经产生的 Scheduler 在被释放后将不会回到 SchedulerPool 中
func (p *SchedulerPool) SetSize(size int) error {
	if size <= 0 {
		return fmt.Errorf("scheduler pool size must greater than 0, got: %d", size)
	}
	p.lock.Lock()
	defer p.lock.Unlock()
	p.size = size
	return nil
}

// Get 获取一个特定时间周期及时间轮尺寸的时间调度器，当池中存在可用的时间调度器时，将会直接返回，否则将会创建一个新的时间调度器
func (p *SchedulerPool) Get() *Scheduler {
	p.lock.Lock()
	defer p.lock.Unlock()
	var scheduler *Scheduler
	if len(p.schedulers) > 0 {
		scheduler = p.schedulers[0]
		p.schedulers = p.schedulers[1:]
		return scheduler
	}
	return newScheduler(p, p.tick, p.wheelSize)
}

// Recycle 释放定时器池的资源并将其重置为全新的状态
//   - 执行该函数后，已有的时间调度器将会被停止，且不会重新加入到池中，一切都是新的开始
func (p *SchedulerPool) Recycle() {
	p.lock.Lock()
	defer p.lock.Unlock()
	for _, scheduler := range p.schedulers {
		scheduler.wheel.Stop()
	}
	p.schedulers = nil
	p.generation++
	return
}

func (p *SchedulerPool) getGeneration() int64 {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.generation
}

func (p *SchedulerPool) getExecutor() func(name string, caller func()) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.executor
}
