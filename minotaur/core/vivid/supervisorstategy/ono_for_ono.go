package supervisorstategy

import "github.com/kercylan98/minotaur/minotaur/core/vivid"

// OneForOne 创建一个 OneForOne 策略的监督者
//   - 当一个 Actor 发生故障时，只有该 Actor 会被处理
func OneForOne(decide vivid.SupervisorDecide) vivid.SupervisorStrategy {
	return &oneForOne{decide: decide}
}

type oneForOne struct {
	decide vivid.SupervisorDecide
}

func (o *oneForOne) OnAccident(system *vivid.ActorSystem, supervisor vivid.Supervisor, accident vivid.ActorRef, reason, message vivid.Message) {
	switch o.decide(reason, message) {
	case vivid.DirectiveRestart:
		supervisor.Restart(accident)
	case vivid.DirectiveStop:
		supervisor.Stop(accident)
	case vivid.DirectiveEscalate:
		supervisor.Escalate(reason, message, nil)
	case vivid.DirectiveResume:
		panic("not impl")
	default:
		panic("not support directive")
	}
}
