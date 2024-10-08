package supervision

import (
	"github.com/kercylan98/minotaur/engine/prc"
)

func NewAccidentRecord(primeCulprit, victim *prc.ProcessId, supervisor Supervisor, message, reason prc.Message, strategy Strategy, state *AccidentState, stack []byte) *AccidentRecord {
	return &AccidentRecord{
		PrimeCulprit: primeCulprit,
		Victim:       victim,
		Supervisor:   supervisor,
		Message:      message,
		Reason:       reason,
		Strategy:     strategy,
		State:        state,
		Stack:        stack,
	}
}

// AccidentRecord 事故记录
type AccidentRecord struct {
	PrimeCulprit *prc.ProcessId // 事故元凶（通常是消息发送人，可能不存在或逃逸）
	Victim       *prc.ProcessId // 事故受害者
	Supervisor   Supervisor     // 事故监管者，将由其进行决策
	Message      prc.Message    // 造成事故发生的消息
	Reason       prc.Message    // 事故原因
	Strategy     Strategy       // 受害人携带的监督策略，应由责任人执行
	State        *AccidentState // 事故状态
	Stack        []byte         // 事件堆栈
}
