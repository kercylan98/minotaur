package supervision

import (
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"time"
)

// OneForOne 创建一个 OneForOne 策略的监督者
//   - 当一个 Actor 发生故障时，只有该 Actor 会被处理
//   - 当 restartCount 为负数时，将会无限重启
func OneForOne(restartCount int, baseDelay, maxDelay time.Duration, decide Decide) Strategy {
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

func (o *oneForOne) OnPolicyDecision(record *AccidentRecord) {
	switch o.decide.Decide(record) {
	case DirectiveRestart:
		next := chrono.StandardExponentialBackoff(record.State.AccidentCount(), o.restartCount, o.baseDelay, o.maxDelay)
		if next != -1 {
			time.AfterFunc(next, func() {
				record.Supervisor.Restart(record.Victim)
			})
		} else {
			record.Supervisor.Stop(record.Victim)
		}
	case DirectiveStop:
		record.Supervisor.Stop(record.Victim)
	case DirectiveEscalate:
		record.Supervisor.Escalate(record)
	case DirectiveResume:
		record.Supervisor.Resume(record.Victim)
	default:
		panic("not support directive")
	}
}
