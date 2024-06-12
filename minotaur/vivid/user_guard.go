package vivid

type userGuardActor struct {
}

func (u *userGuardActor) OnReceive(ctx MessageContext) {
	switch msg := ctx.GetMessage().(type) {
	case OnInit[*userGuardActor]:
		msg.Options.WithSupervisor(func(message, reason Message) Directive {
			switch message.(type) {
			case OnBoot:
				// OnBoot 阶段发生错误可能存在包含网络 IO 的初始化，因此需要重试
				return DirectiveRestart
			case OnTerminate:
				// OnTerminate 阶段发生错误通常是用户持久化等行为发生错误，后续会停止，应忽略
				return DirectiveResume
			default:
				// 其他情况下，尝试重启 Actor
				return DirectiveRestart
			}
		})
	}
}
