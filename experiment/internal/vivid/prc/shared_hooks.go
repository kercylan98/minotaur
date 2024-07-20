package prc

// SharedStartHook 共享启动时的钩子
type SharedStartHook interface {
	// OnSharedStart 共享启动
	OnSharedStart()
}

// FunctionalSharedStartHook 是一个函数式的共享启动时的钩子
type FunctionalSharedStartHook func()

// OnSharedStart 共享启动
func (f FunctionalSharedStartHook) OnSharedStart() {
	f()
}
