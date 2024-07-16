package supervisor

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"github.com/kercylan98/minotaur/toolkit"
)

var stopSingleton = toolkit.NewInertiaSingleton(func() vivid.SupervisorStrategy {
	return vivid.FunctionalSupervisorStrategy(func(system *vivid.ActorSystem, accident vivid.Accident) {
		accident.Responsible().Stop(accident.AccidentActor())
	})
})

// Stop 任何意外都将执行停止
func Stop() vivid.SupervisorStrategy {
	return stopSingleton.Get()
}
