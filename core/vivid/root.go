package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"time"
)

type root struct {
	strategy SupervisorStrategy
}

func (r *root) OnReceive(ctx ActorContext) {
}

func (r *root) OnAccident(system *ActorSystem, accident Accident) {
	if r.strategy != nil {
		r.strategy.OnAccident(system, accident)
	} else {
		next := chrono.StandardExponentialBackoff(accident.RestartCount(), 10, time.Millisecond*300, time.Second*3)
		if next != -1 {
			time.AfterFunc(next, func() {
				accident.Responsible().Restart(accident.AccidentActor())
			})
		} else {
			accident.Responsible().Stop(accident.AccidentActor())
		}
	}
}
