package chrono

import (
	"github.com/RussellLuo/timingwheel"
	"github.com/gorhill/cronexpr"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"reflect"
	"sync"
	"time"
)

const (
	DefaultSchedulerTick      = time.Millisecond * 10
	DefaultSchedulerWheelSize = 10
)

const (
	SchedulerForever   = -1 //  无限循环
	SchedulerOnce      = 1  // 一次
	SchedulerInstantly = 0  // 立刻
)

type SchedulerExecutor func(name string, caller func(), separate func())

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
	scheduler := &Scheduler{
		tick:  tick,
		wheel: timingwheel.NewTimingWheel(tick, wheelSize),
		tasks: make(map[string]*schedulerTask),
	}
	scheduler.wheel.Start()
	return scheduler
}

// Scheduler 并发安全的时间调度器
type Scheduler struct {
	wheel    *timingwheel.TimingWheel  // 时间轮
	tasks    map[string]*schedulerTask // 所有任务
	lock     sync.RWMutex              // 用于确保并发安全的锁
	tick     time.Duration             // 时间周期
	executor SchedulerExecutor         // 任务执行器
}

func (s *Scheduler) unlockUnregisterTask(name string) {
	if task, exist := s.tasks[name]; exist {
		task.close()
		delete(s.tasks, name)
	}
}

// SetExecutor 设置任务执行器
func (s *Scheduler) SetExecutor(executor SchedulerExecutor) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.executor = executor
}

// Clear 清空所有任务但不关闭时间调度器
func (s *Scheduler) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.executor = nil
	for name, task := range s.tasks {
		task.close()
		delete(s.tasks, name)
	}
}

// Close 关闭时间调度器并停止所有任务，此后不再可用
func (s *Scheduler) Close() {
	s.lock.Lock()
	defer s.lock.Unlock()
	for name, task := range s.tasks {
		task.close()
		delete(s.tasks, name)
	}
	s.wheel.Stop()
}

// UnregisterTask 取消特定任务的执行计划的注册
//   - 如果任务不存在，则不执行任何操作
func (s *Scheduler) UnregisterTask(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.unlockUnregisterTask(name)
}

// GetRegisteredTasks 获取所有未执行完成的任务名称
func (s *Scheduler) GetRegisteredTasks() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return collection.ConvertMapKeysToSlice(s.tasks)
}

// RegisterCronTask 通过 cron 表达式注册一个任务。
//   - 当 cron 表达式错误时，将会返回错误信息
func (s *Scheduler) RegisterCronTask(name, expression string, function any, args ...any) error {
	expr, err := cronexpr.Parse(expression)
	if err != nil {
		return err
	}
	s.task(name, 0, 0, expr, 0, function, args...)
	return nil
}

// RegisterImmediateCronTask 与 RegisterCronE 相同，但是会立即执行一次
func (s *Scheduler) RegisterImmediateCronTask(name, expression string, function any, args ...any) error {
	if err := s.RegisterCronTask(name, expression, function, args...); err != nil {
		return err
	}
	s.call(name, function, args...)
	return nil
}

// RegisterAfterTask 注册一个在特定时间后执行一次的任务
func (s *Scheduler) RegisterAfterTask(name string, after time.Duration, function any, args ...any) {
	s.task(name, after, s.tick, nil, 1, function, args...)
}

// RegisterRepeatedTask 注册一个在特定时间后反复执行的任务
func (s *Scheduler) RegisterRepeatedTask(name string, after, interval time.Duration, times int, function any, args ...any) {
	s.task(name, after, interval, nil, times, function, args...)
}

// RegisterDayMomentTask 注册一个在每天特定时刻执行的任务
//   - 其中 lastExecuted 为上次执行时间，adjust 为时间偏移量，hour、min、sec 为时、分、秒
//   - 当上次执行时间被错过时，将会立即执行一次
func (s *Scheduler) RegisterDayMomentTask(name string, lastExecuted time.Time, offset time.Duration, hour, min, sec int, function any, args ...any) {
	now := time.Now().Add(offset)
	if lastExecuted.Before(now) && now.Sub(lastExecuted) > Day {
		s.call(name, function, args...)
	}

	moment := GetNextMoment(now, hour, min, sec)
	s.RegisterRepeatedTask(name, moment.Sub(now), time.Hour*24, SchedulerForever, function, args...)
}

func (s *Scheduler) task(name string, after, interval time.Duration, expr *cronexpr.Expression, times int, function any, args ...any) {
	// init task
	if expr == nil {
		if after < s.tick {
			after = s.tick
		}
		if interval < s.tick {
			interval = s.tick
		}
	}

	var argVof = make([]reflect.Value, len(args))
	for i, v := range args {
		argVof[i] = reflect.ValueOf(v)
	}

	task := &schedulerTask{
		name:      name,
		after:     after,
		interval:  interval,
		total:     times,
		function:  reflect.ValueOf(function),
		args:      argVof,
		scheduler: s,
		expr:      expr,
	}

	// register task
	s.lock.Lock()
	defer s.lock.Unlock()

	s.unlockUnregisterTask(name)

	s.tasks[name] = task
	var caller = task.caller
	if s.executor != nil {
		caller = func() {
			s.executor(task.Name(), task.caller, func() {
				task.separate = true
			})
		}
	}
	s.wheel.ScheduleFunc(task, caller)
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
		}, func() {})
	} else {
		f.Call(values)
	}
}
