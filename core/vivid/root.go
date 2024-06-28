package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/log"
)

type root struct {
}

func (r *root) OnReceive(ctx ActorContext) {

}

func (r *root) OnAccident(system *ActorSystem, accident Accident) {
	log.Error("Accident", log.String("actor", accident.AccidentActor().Address().String()), log.Any("reason", accident.Reason()))
	if accident.RestartCount() < 3 {
		accident.Responsible().Restart(accident.AccidentActor())
	} else {
		accident.Responsible().Stop(accident.AccidentActor())
	}
}
