package vivid

const (
	DirectiveStop     Directive = iota + 1 // 停止子 Actor。_OnTerminated 生命周期钩子方法会被调用，子 Actor 的状态会被清理
	DirectiveRestart                       // 停止并重新启动子 Actor。子 Actor 的状态会被重置
	DirectiveResume                        // 忽略异常，让子 Actor 继续处理后续的消息。子 Actor 的状态保持不变
	DirectiveEscalate                      // 升级 Actor（向上级 Actor 报告异常）
)

// Directive 监管指令
type Directive uint8

// SupervisorStrategy 监管策略
type SupervisorStrategy interface {
	OnAccident(system *ActorSystem, supervisor Supervisor, accident ActorRef, reason, message Message)
}
