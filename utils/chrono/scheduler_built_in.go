package chrono

import "time"

const (
	BuiltInSchedulerWheelSize = 50
)

var (
	buildInSchedulerPool *SchedulerPool
	builtInScheduler     *Scheduler
)

func init() {
	buildInSchedulerPool = NewDefaultSchedulerPool()
	builtInScheduler = NewScheduler(DefaultSchedulerTick, BuiltInSchedulerWheelSize)
}

// BuiltInSchedulerPool 获取内置的由 NewDefaultSchedulerPool 函数创建的时间调度器对象池
func BuiltInSchedulerPool() *SchedulerPool {
	return buildInSchedulerPool
}

// BuiltInScheduler 获取内置的由 NewScheduler(DefaultSchedulerTick, BuiltInSchedulerWheelSize) 创建的时间调度器
func BuiltInScheduler() *Scheduler {
	return builtInScheduler
}

// UnregisterTask 调用内置时间调度器 BuiltInScheduler 的 Scheduler.UnregisterTask 函数
func UnregisterTask(name string) {
	BuiltInScheduler().UnregisterTask(name)
}

// RegisterCronTask 调用内置时间调度器 BuiltInScheduler 的 Scheduler.RegisterCronTask 函数
func RegisterCronTask(name, expression string, function interface{}, args ...interface{}) error {
	return BuiltInScheduler().RegisterCronTask(name, expression, function, args...)
}

// RegisterImmediateCronTask 调用内置时间调度器 BuiltInScheduler 的 Scheduler.RegisterImmediateCronTask 函数
func RegisterImmediateCronTask(name, expression string, function interface{}, args ...interface{}) error {
	return BuiltInScheduler().RegisterImmediateCronTask(name, expression, function, args...)
}

// RegisterAfterTask 调用内置时间调度器 BuiltInScheduler 的 Scheduler.RegisterAfterTask 函数
func RegisterAfterTask(name string, after time.Duration, function interface{}, args ...interface{}) {
	BuiltInScheduler().RegisterAfterTask(name, after, function, args...)
}

// RegisterRepeatedTask 调用内置时间调度器 BuiltInScheduler 的 Scheduler.RegisterRepeatedTask 函数
func RegisterRepeatedTask(name string, after, interval time.Duration, times int, function interface{}, args ...interface{}) {
	BuiltInScheduler().RegisterRepeatedTask(name, after, interval, times, function, args...)
}

// RegisterDayMomentTask 调用内置时间调度器 BuiltInScheduler 的 Scheduler.RegisterDayMomentTask 函数
func RegisterDayMomentTask(name string, lastExecuted time.Time, offset time.Duration, hour, min, sec int, function interface{}, args ...interface{}) {
	BuiltInScheduler().RegisterDayMomentTask(name, lastExecuted, offset, hour, min, sec, function, args...)
}
