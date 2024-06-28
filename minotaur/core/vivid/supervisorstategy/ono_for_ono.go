package supervisorstategy

import "github.com/kercylan98/minotaur/minotaur/core/vivid"

// OneForOne 创建一个 OneForOne 策略的监督者
//   - 当一个 Actor 发生故障时，只有该 Actor 会被处理
func OneForOne(decide vivid.SupervisorDecide, restartLimit int) vivid.SupervisorStrategy {
	return &oneForOne{decide: decide, restartLimit: restartLimit}
}

type oneForOne struct {
	decide       vivid.SupervisorDecide
	restartLimit int
}

func (o *oneForOne) OnAccident(system *vivid.ActorSystem, accident vivid.Accident) {
	switch o.decide(accident.Reason(), accident.Message()) {
	case vivid.DirectiveRestart:
		if accident.RestartCount() < o.restartLimit {
			accident.Responsible().Restart(accident.AccidentActor())
		} else {
			accident.Responsible().Stop(accident.AccidentActor())
		}
	case vivid.DirectiveStop:
		accident.Responsible().Stop(accident.AccidentActor())
	case vivid.DirectiveEscalate:
		accident.Responsible().Escalate(accident)
	case vivid.DirectiveResume:
		accident.Responsible().Resume(accident.AccidentActor())
	default:
		panic("not support directive")
	}
}
