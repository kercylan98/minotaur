package core

func NewProcessManager(host string, port uint64, bucketSize int) *ProcessManager {
	mgr := &ProcessManager{
		processRegisterTable: newProcessRegisterTable(bucketSize),
		host:                 host,
		port:                 port,
	}

	return mgr
}

// ProcessManager 进程管理器
type ProcessManager struct {
	*processRegisterTable
	host string
	port uint64
}

func (mgr *ProcessManager) GetProcess(ref *ProcessRef) Process {
	process := ref.cache.Load()
	if process != nil {
		if p := *process; p.Deaden() {
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
