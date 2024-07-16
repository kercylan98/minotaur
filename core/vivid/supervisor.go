package vivid

import "time"

type Supervisor interface {
	// Children 获取该监管者下的所有子 Actor
	Children() []ActorRef

	// Restart 重启子级 Actor，如果子级 Actor 还存在更深的子级 Actor，那么它们将会被停止
	//  - 当 children 为空时，重启所有子级 Actor
	Restart(children ...ActorRef)

	// Stop 停止子级 Actor 及其所有子级 Actor
	//  - 当 children 为空时，停止所有子级 Actor
	Stop(children ...ActorRef)

	// Resume 恢复子级 Actor
	Resume(children ...ActorRef)

	// Escalate 升级
	Escalate(accident Accident)
}

// Accident 事故现场，该接口为内部实现，提供可自定义监督策略的能力
type Accident interface {
	// AccidentActor 事故发生的 Actor
	AccidentActor() ActorRef

	// Responsible 事故应该由哪个监管者负责
	Responsible() Supervisor

	// Reason 事故原因
	Reason() Message

	// Message 导致事故的消息
	Message() Message

	// RestartCount 获取重启次数
	RestartCount() int

	// LastRestartTime 获取最后一次重启时间
	LastRestartTime() time.Time

	trySetResponsible(Supervisor)
	trySupervisorStrategy(*ActorSystem) bool
}

// 事故状态
type accidentState struct {
	restartTimes []time.Time
}

type accident struct {
	responsible        Supervisor // 理应负责的监管者
	accidentActor      ActorRef
	supervisorStrategy SupervisorStrategy
	reason             Message
	message            Message
	state              *accidentState
}

func (a *accident) RestartCount() int {
	return len(a.state.restartTimes)
}

func (a *accident) LastRestartTime() time.Time {
	if len(a.state.restartTimes) > 0 {
		return a.state.restartTimes[len(a.state.restartTimes)-1]
	}
	return time.Time{}
}

func (a *accident) AccidentActor() ActorRef {
	return a.accidentActor
}

func (a *accident) Responsible() Supervisor {
	return a.responsible
}

func (a *accident) Reason() Message {
	return a.reason
}

func (a *accident) Message() Message {
	return a.message
}

func (a *accident) trySetResponsible(s Supervisor) {
	if a.responsible == nil {
		a.responsible = s
	}
}

func (a *accident) trySupervisorStrategy(system *ActorSystem) bool {
	if a.supervisorStrategy != nil {
		a.supervisorStrategy.OnAccident(system, a)
		return true
	}
	return false
}
