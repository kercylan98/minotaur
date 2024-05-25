package vivid

type Message = any

// OnOptionApply 该消息将在 Actor 可选项应用前发送，可用于 Actor 对可选项的检查要求
//   - 该阶段 Actor 尚未初始化，不要期待在该阶段能够处理任何 ActorOptions 以外的内容
type OnOptionApply[T Actor] struct {
	Options *ActorOptions[T]
}

// OnPreStart 是 Actor 生命周期的开始阶段，通常用于初始化 Actor 的状态
type OnPreStart struct {
}

// OnDestroy 是 Actor 生命周期的销毁阶段，通常用于释放 Actor 的资源
type OnDestroy struct {
}
