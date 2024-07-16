package vivid

// SupervisorLogger 监管日志记录器
type SupervisorLogger func(reason, message Message)

// SupervisorStrategy 监管策略
type SupervisorStrategy interface {
	OnAccident(system *ActorSystem, accident Accident)
}

// FunctionalSupervisorStrategy 函数式监管策略
type FunctionalSupervisorStrategy func(system *ActorSystem, accident Accident)

// OnAccident 当发生意外时
func (f FunctionalSupervisorStrategy) OnAccident(system *ActorSystem, accident Accident) {
	f(system, accident)
}
