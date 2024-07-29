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

func (ref *ProcessRef) UnmarshalJSON(bytes []byte) error {
	if ref.id == nil {
		ref.id = new(ProcessId)
	}
	err := toolkit.UnmarshalJSONE(bytes, ref.id)
	if err == nil {
		ref.cache.Store(nil)
	}
	return nil
}

func (ref *ProcessRef) MarshalJSON() ([]byte, error) {
	return toolkit.MarshalJSONE(ref)
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

// ChangePhysicalAddress 改变进程引用的物理地址并返回新的进程 Id
func (ref *ProcessRef) ChangePhysicalAddress(address PhysicalAddress) *ProcessId {
	nid := ref.GetId()
	nid.PhysicalAddress = address
	return nid
}

// URL 获取进程引用的 URL
func (ref *ProcessRef) URL() *url.URL {
	return ref.id.URL()
}

// Clone 克隆一个进程引用
func (ref *ProcessRef) Clone() *ProcessRef {
	return &ProcessRef{
		id: ref.id,
	}
}

func (ref *ProcessRef) GetId() *ProcessId {
	return &ProcessId{
		LogicalAddress:  ref.id.LogicalAddress,
		PhysicalAddress: ref.id.PhysicalAddress,
	}
}

func (ref *ProcessRef) Equal(r *ProcessRef) bool {
	return ref.id.Equal(r.id)
}
