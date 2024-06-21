package vivid

type Message = any

// OnInit 该消息将在 Actor 可选项应用前发送，可用于 Actor 对可选项的检查要求
//   - 该阶段 Actor 尚未初始化完成，不要期待在该阶段能够处理任何 ActorOptions 以外的内容
type OnInit[T Actor] struct {
	Options *ActorOptions[T]
}

// OnBoot 该消息在 Actor 初始化完成后发送，用于通知 Actor 初始化完成。该阶段 Actor 已经可以处理消息
type OnBoot struct {
}

// OnRestart 该消息将在 Actor 重启前发送
type OnRestart struct {
}

// OnTerminate 在 Actor 收到该消息后，将会被安全的终止
type OnTerminate struct {
	restart bool // 是否重启

	Context any // 用户自定义的上下文
}

// onActorRefTyped 该消息将在 ActorOfT 时发送，用于 ActorRef 的类型转换
type onActorRefTyped struct {
	ref ActorRef
}

// OnActorTyped 该消息在 Actor 被类型化后发送，用于在 Actor 本身的消息中获取到类型化的 ActorRef
type OnActorTyped[T ActorTyped] struct {
	Typed T
}
