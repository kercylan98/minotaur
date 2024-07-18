package supervision

import "github.com/kercylan98/minotaur/toolkit"

var resumeStrategy = toolkit.NewInertiaSingleton[Strategy](func() Strategy {
	return FunctionalStrategy(func(record *AccidentRecord) {
		record.Supervisor.Resume(record.Victim)
	})
})

// ResumeStrategy 是一个监管策略，它将在发生事故时恢复进程
func ResumeStrategy() Strategy {
	return resumeStrategy.Get()
}
