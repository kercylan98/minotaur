package module

// Module 是所有模块的核心接口，必须实现
type Module[App Application] interface {
	// OnInitialize 在模块自身初始化时调用，用于进行基本的内部设置
	OnInitialize(app App) error
}

// Dependency 处理依赖模块的初始化逻辑，模块可以选择实现
type Dependency[App Application] interface {
	Module[App]
	// OnDependencyInitialize 在所有依赖模块初始化完成后调用
	OnDependencyInitialize(loader *Loader[App]) error
}

// DependencySetup 处理依赖内容的初始化逻辑，模块可以选择实现
type DependencySetup[App Application] interface {
	Module[App]
	// OnDependencySetup 在所有依赖模块的基础上完成当前模块的额外设置
	// （例如：读取配置、获取依赖服务实例等）
	OnDependencySetup() error
}

// Runnable 提供运行时逻辑，模块可以选择实现
type Runnable[App Application] interface {
	Module[App]
	// OnRun 启动模块的主要功能，例如启动服务、监听事件等
	OnRun()
}

// Stoppable 提供停止逻辑，模块可以选择实现
type Stoppable[App Application] interface {
	Module[App]
	// OnStop 在应用关闭或模块被卸载时调用，用于释放资源
	OnStop() error
}
