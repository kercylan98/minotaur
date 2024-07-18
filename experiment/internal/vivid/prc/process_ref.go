package prc

import (
	"sync/atomic"
)

func NewProcessRef(id *ProcessId) *ProcessRef {
	return &ProcessRef{
		id: id,
	}
}

// ProcessRef 进程引用是包含缓存的 ProcessId，它能够将 Process 信息进行缓存而不必频繁地向资源控制器搜索
type ProcessRef struct {
	id    *ProcessId
	cache atomic.Pointer[Process]
}

// PhysicalAddress 获取进程引用的物理地址
func (ref *ProcessRef) PhysicalAddress() PhysicalAddress {
	return ref.id.PhysicalAddress
}

// LogicalAddress 获取进程引用的逻辑地址
func (ref *ProcessRef) LogicalAddress() LogicalAddress {
	return ref.id.LogicalAddress
}

// DerivationProcessId 衍生一个新的进程 Id
func (ref *ProcessRef) DerivationProcessId(name string) *ProcessId {
	return ref.id.Derivation(name)
}
