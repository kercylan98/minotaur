package prc

import "sync/atomic"

// ProcessRef 进程引用是包含缓存的 ProcessId，它能够将 Process 信息进行缓存而不必频繁地向资源控制器搜索
type ProcessRef struct {
	id    *ProcessId
	cache atomic.Pointer[Process]
}
