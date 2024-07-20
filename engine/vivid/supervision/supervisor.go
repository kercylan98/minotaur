package supervision

import "github.com/kercylan98/minotaur/engine/prc"

type Supervisor interface {
	// Ref 获取监管者的进程引用
	Ref() *prc.ProcessRef

	// Children 获取该监管者下的所有子进程引用
	Children() []*prc.ProcessRef

	// Restart 重启指定进程，会导致该进程的所有子进程被停止
	Restart(refs ...*prc.ProcessRef)

	// Stop 停止指定进程及其所有子进程
	Stop(refs ...*prc.ProcessRef)

	// Resume 恢复指定进程
	Resume(refs ...*prc.ProcessRef)

	// Escalate 升级事故
	Escalate(record *AccidentRecord)
}
