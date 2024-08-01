package prc

import (
	"github.com/kercylan98/minotaur/toolkit"
	"net/url"
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

func (ref *ProcessRef) GetId() *ProcessId {
	return ref.id
}

// PhysicalAddress 获取进程引用的物理地址
func (ref *ProcessRef) PhysicalAddress() PhysicalAddress {
	return ref.GetId().PhysicalAddress
}

// LogicalAddress 获取进程引用的逻辑地址
func (ref *ProcessRef) LogicalAddress() LogicalAddress {
	return ref.GetId().LogicalAddress
}

func (ref *ProcessRef) UnmarshalJSON(bytes []byte) error {
	if ref.GetId() == nil {
		ref.id = new(ProcessId)
	}
	err := toolkit.UnmarshalJSONE(bytes, ref.GetId())
	if err == nil {
		ref.cache.Store(nil)
	}
	return nil
}

func (ref *ProcessRef) MarshalJSON() ([]byte, error) {
	return toolkit.MarshalJSONE(ref)
}

// DerivationProcessId 衍生一个新的进程 Id
func (ref *ProcessRef) DerivationProcessId(name string) *ProcessId {
	return ref.GetId().Derivation(name)
}

// URL 获取进程引用的 URL
func (ref *ProcessRef) URL() *url.URL {
	return ref.GetId().URL()
}

// Clone 克隆一个进程引用
func (ref *ProcessRef) Clone() *ProcessRef {
	return &ProcessRef{
		id: ref.GetId(),
	}
}

func (ref *ProcessRef) Equal(r *ProcessRef) bool {
	return ref.GetId().Equal(r.GetId())
}
