package supervision

const (
	DirectiveStop     Directive = iota + 1 // 停止子 Actor。OnTerminated 生命周期钩子方法会被调用，子 Actor 的状态会被清理
	DirectiveRestart                       // 停止并重新启动子 Actor。子 Actor 的状态会被重置
	DirectiveResume                        // 忽略异常，让子 Actor 继续处理后续的消息。子 Actor 的状态保持不变
	DirectiveEscalate                      // 升级 Actor（向上级 Actor 报告异常）
)

// Directive 监管指令
type Directive uint8

// Decide 监管决策是一个用于监督者策略的函数接口
type Decide interface {
	Decide(record *AccidentRecord) Directive
}

// FunctionalDecide 是一个用于监督者策略的函数类型
type FunctionalDecide func(record *AccidentRecord) Directive

// Decide 实现了 Decide 接口
func (f FunctionalDecide) Decide(record *AccidentRecord) Directive {
	return f(record)
}
