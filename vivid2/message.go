package vivid

type Message = any

// OnPreStart 是 Actor 生命周期的开始阶段，通常用于初始化 Actor 的状态
type OnPreStart struct {
}

// OnDestroy 是 Actor 生命周期的销毁阶段，通常用于释放 Actor 的资源
type OnDestroy struct {
}
