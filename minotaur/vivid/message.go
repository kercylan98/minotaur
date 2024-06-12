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
}
