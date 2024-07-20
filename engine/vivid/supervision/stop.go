package supervision

import "github.com/kercylan98/minotaur/toolkit"

var stopStrategy = toolkit.NewInertiaSingleton[Strategy](func() Strategy {
	return FunctionalStrategy(func(record *AccidentRecord) {
		record.Supervisor.Stop(record.Victim)
	})
})

// StopStrategy 是一个监管策略，它将在发生事故时停止进程
func StopStrategy() Strategy {
	return stopStrategy.Get()
}
