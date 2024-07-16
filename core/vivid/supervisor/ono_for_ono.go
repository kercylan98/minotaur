package supervisor

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"time"
)

// OneForOne 创建一个 OneForOne 策略的监督者
//   - 当一个 Actor 发生故障时，只有该 Actor 会被处理
func OneForOne(restartCount int, baseDelay, maxDelay time.Duration, decide Decide) vivid.SupervisorStrategy {
	return &oneForOne{
		decide:       decide,
		restartCount: restartCount,
		baseDelay:    baseDelay,
		maxDelay:     maxDelay,
	}
}

type oneForOne struct {
	decide             Decide
	restartCount       int
	exponentialBackoff int
	baseDelay          time.Duration
	maxDelay           time.Duration
}

func (o *oneForOne) OnAccident(system *vivid.ActorSystem, accident vivid.Accident) {
	switch o.decide(accident.Reason(), accident.Message()) {
	case DirectiveRestart:
		next := chrono.StandardExponentialBackoff(accident.RestartCount(), o.restartCount, o.baseDelay, o.maxDelay)
		if next != -1 {
			time.AfterFunc(next, func() {
				accident.Responsible().Restart(accident.AccidentActor())
			})
		} else {
			accident.Responsible().Stop(accident.AccidentActor())
		}
	case DirectiveStop:
		accident.Responsible().Stop(accident.AccidentActor())
	case DirectiveEscalate:
		accident.Responsible().Escalate(accident)
	case DirectiveResume:
		accident.Responsible().Resume(accident.AccidentActor())
	default:
		panic("not support directive")
	}
}
