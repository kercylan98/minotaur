package vivid

type SupervisorDecide func(reason, message Message) Directive

type Supervisor interface {
	// Children 获取该监管者下的所有子 Actor
	Children() []ActorRef

	// Restart 重启子级 Actor，如果子级 Actor 还存在更深的子级 Actor，那么它们将会被停止
	//  - 当 children 为空时，重启所有子级 Actor
	Restart(children ...ActorRef)

	// Stop 停止子级 Actor 及其所有子级 Actor
	//  - 当 children 为空时，停止所有子级 Actor
	Stop(children ...ActorRef)

	// Resume 恢复Actor
	Resume()

	// Escalate 升级
	Escalate(accidents *_OnAccidents)
}
