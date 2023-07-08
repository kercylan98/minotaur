package task

type Option func(task *Task)

// WithInitCount 通过初始化计数的方式创建任务
func WithInitCount(count int) Option {
	return func(task *Task) {
		task.count = count
	}
}

// WithDone 通过指定任务完成计数的方式创建任务
func WithDone(done int) Option {
	return func(task *Task) {
		task.done = done
	}
}
