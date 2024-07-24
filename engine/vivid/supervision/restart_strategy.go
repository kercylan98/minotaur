package supervision

import "github.com/kercylan98/minotaur/toolkit"

var restartStrategy = toolkit.NewInertiaSingleton[Strategy](func() Strategy {
	return FunctionalStrategy(func(record *AccidentRecord) {
		record.Supervisor.Stop(record.Victim)
	})
})

// RestartStrategy 是一个监管策略，它将在发生事故时立即重启
func RestartStrategy() Strategy {
	return restartStrategy.Get()
}
