package chrono

import (
	"github.com/RussellLuo/timingwheel"
	"github.com/gorhill/cronexpr"
	"github.com/kercylan98/minotaur/utils/collection"
	"reflect"
	"sync"
	"time"
)

const (
	DefaultSchedulerTick      = SchedulerPoolDefaultTick
	DefaultSchedulerWheelSize = SchedulerPoolDefaultWheelSize
)

const (
	SchedulerForever   = -1 //  无限循环
	SchedulerOnce      = 1  // 一次
	SchedulerInstantly = 0  // 立刻
)

// NewDefaultScheduler 创建一个默认的时间调度器
//   - tick: DefaultSchedulerTick
//   - wheelSize: DefaultSchedulerWheelSize
func NewDefaultScheduler() *Scheduler {
	return NewScheduler(DefaultSchedulerTick, DefaultSchedulerWheelSize)
}

// NewScheduler 创建一个并发安全的时间调度器
//   - tick: 时间轮的刻度间隔。
//   - wheelSize: 时间轮的大小。
func NewScheduler(tick time.Duration, wheelSize int64) *Scheduler {
	return newScheduler(nil, tick, wheelSize)
}

func newScheduler(pool *SchedulerPool, tick time.Duration, wheelSize int64) *Scheduler {
	scheduler := &Scheduler{
		pool:  pool,
		wheel: timingwheel.NewTimingWheel(tick, wheelSize),
		tasks: make(map[string]*schedulerTask),
	}
	if pool != nil {
		scheduler.generation = pool.getGeneration()
	}
	scheduler.wheel.Start()
	return scheduler
}

// Scheduler 并发安全的时间调度器
type Scheduler struct {
	pool       *SchedulerPool            // 时间调度器所属的池，当该值为 nil 时，该时间调度器不属于任何池
	wheel      *timingwheel.TimingWheel  // 时间轮
	tasks      map[string]*schedulerTask // 所有任务
	lock       sync.RWMutex              // 用于确保并发安全的锁
	generation int64                     // 时间调度器的代数
	tick       time.Duration             // 时间周期

	executor func(name string, caller func()) // 任务执行器
}

// SetExecutor 设置任务执行器
//   - 如果该任务执行器来自于时间调度器对象池，那么默认情况下将会使用时间调度器对象池的任务执行器，主动设置将会覆盖默认的任务执行器
func (s *Scheduler) SetExecutor(executor func(name string, caller func())) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.executor = executor
}

// Release 释放时间调度器，时间调度器被释放后将不再可用，如果时间调度器属于某个池且池未满，则会重新加入到池中
//   - 释放后所有已注册的任务将会被取消
func (s *Scheduler) Release() {
	s.lock.Lock()
	defer s.lock.Unlock()
	for name, task := range s.tasks {
		task.close()
		delete(s.tasks, name)
	}

	if s.pool == nil || s.pool.getGeneration() != s.generation {
		s.wheel.Stop()
		return
	}

	s.pool.lock.Lock()
	if len(s.pool.schedulers) < s.pool.size {
		s.pool.schedulers = append(s.pool.schedulers, s)
		s.pool.lock.Unlock()
		return
	}
	s.pool.lock.Unlock()
	s.wheel.Stop()
}

// UnregisterTask 取消特定任务的执行计划的注册
//   - 如果任务不存在，则不执行任何操作
func (s *Scheduler) UnregisterTask(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if task, exist := s.tasks[name]; exist {
		task.close()
		delete(s.tasks, name)
	}
}

// GetRegisteredTasks 获取所有未执行完成的任务名称
func (s *Scheduler) GetRegisteredTasks() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return collection.ConvertMapKeysToSlice(s.tasks)
}

// RegisterCronTask 通过 cron 表达式注册一个任务。
//   - 当 cron 表达式错误时，将会返回错误信息
func (s *Scheduler) RegisterCronTask(name, expression string, function interface{}, args ...interface{}) error {
	expr, err := cronexpr.Parse(expression)
	if err != nil {
		return err
	}
	s.task(name, 0, 0, expr, 0, function, args...)
	return nil
}

// RegisterImmediateCronTask 与 RegisterCronE 相同，但是会立即执行一次
func (s *Scheduler) RegisterImmediateCronTask(name, expression string, function interface{}, args ...interface{}) error {
	if err := s.RegisterCronTask(name, expression, function, args...); err != nil {
		return err
	}
	s.call(name, function, args...)
	return nil
}

// RegisterAfterTask 注册一个在特定时间后执行一次的任务
func (s *Scheduler) RegisterAfterTask(name string, after time.Duration, function interface{}, args ...interface{}) {
	s.task(name, after, s.pool.tick, nil, 1, function, args...)
}

// RegisterRepeatedTask 注册一个在特定时间后反复执行的任务
func (s *Scheduler) RegisterRepeatedTask(name string, after, interval time.Duration, times int, function interface{}, args ...interface{}) {
	s.task(name, after, interval, nil, times, function, args...)
}

// RegisterDayMomentTask 注册一个在每天特定时刻执行的任务
//   - 其中 lastExecuted 为上次执行时间，adjust 为时间偏移量，hour、min、sec 为时、分、秒
//   - 当上次执行时间被错过时，将会立即执行一次
func (s *Scheduler) RegisterDayMomentTask(name string, lastExecuted time.Time, offset time.Duration, hour, min, sec int, function interface{}, args ...interface{}) {
	now := time.Now().Add(offset)
	if IsMomentReached(now, lastExecuted, hour, min, sec) {
		s.call(name, function, args...)
	}

	moment := GetNextMoment(now, hour, min, sec)
	s.RegisterRepeatedTask(name, moment.Sub(now), time.Hour*24, SchedulerForever, function, args...)
}

func (s *Scheduler) task(name string, after, interval time.Duration, expr *cronexpr.Expression, times int, function interface{}, args ...interface{}) {
	s.UnregisterTask(name)

	if expr == nil {
		if after < s.tick {
			after = s.tick
		}
		if interval < s.tick {
			interval = s.tick
		}
	}

	var values = make([]reflect.Value, len(args))
	for i, v := range args {
		values[i] = reflect.ValueOf(v)
	}

	task := &schedulerTask{
		name:      name,
		after:     after,
		interval:  interval,
		total:     times,
		function:  reflect.ValueOf(function),
		args:      values,
		scheduler: s,
		expr:      expr,
	}
	var executor func(name string, caller func())
	if s.pool != nil {
		executor = s.pool.getExecutor()
	}
	s.lock.Lock()
	if s.executor != nil {
		executor = s.pool.getExecutor()
	}

	s.tasks[name] = task
	if executor != nil {
		task.timer = s.wheel.ScheduleFunc(task, func() {
			executor(task.Name(), task.caller)
		})
	} else {
		task.timer = s.wheel.ScheduleFunc(task, task.caller)
	}
	s.lock.Unlock()
}

func (s *Scheduler) call(name string, function any, args ...any) {
	var values = make([]reflect.Value, len(args))
	for i, v := range args {
		values[i] = reflect.ValueOf(v)
	}
	f := reflect.ValueOf(function)
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.executor != nil {
		s.executor(name, func() {
			f.Call(values)
		})
	} else {
		f.Call(values)
	}
}
