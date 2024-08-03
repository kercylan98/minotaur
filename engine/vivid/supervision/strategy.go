package supervision

// Strategy 是用于对事故记录进行决策处理的监督策略
type Strategy interface {
	// OnPolicyDecision 当发生事故时
	OnPolicyDecision(record *AccidentRecord)
}

// FunctionalStrategy 是一个函数类型的监督策略
type FunctionalStrategy func(record *AccidentRecord)

// OnPolicyDecision 当发生事故时
func (f FunctionalStrategy) OnPolicyDecision(record *AccidentRecord) {
	f(record)
}
