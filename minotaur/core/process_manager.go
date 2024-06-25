package core

func NewProcessManager(address Address, bucketSize int, defaultProcess ...Process) *ProcessManager {
	mgr := &ProcessManager{
		processRegisterTable: newProcessRegisterTable(address, bucketSize, defaultProcess...),
	}

	return mgr
}

// ProcessManager 进程管理器
type ProcessManager struct {
	*processRegisterTable
}

func (mgr *ProcessManager) GetProcess(ref *ProcessRef) Process {
	process := ref.cache.Load()
	if process != nil {
		p := *process
		if status, ok := p.(ProcessStatus); ok && status.Deaden() {
			ref.cache.Store(nil)
		} else {
			return p
		}
	}

	cache, exists := mgr.processRegisterTable.GetProcess(ref.address)
	if exists {
		ref.cache.Store(&cache)
	}
	return cache
}
