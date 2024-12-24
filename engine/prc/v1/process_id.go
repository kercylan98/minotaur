package prcv1

import (
	"net/url"
	"strings"
)

func NewProcessId(physicalAddress PhysicalAddress, logicalAddress LogicalAddress) *ProcessId {
	return &ProcessId{
		LogicalAddress:  logicalAddress,
		PhysicalAddress: physicalAddress,
	}
}

// Derivation 衍生一个新的进程 Id
func (pid *ProcessId) Derivation(name string) *ProcessId {
	ld := pid.GetLogicalAddress()
	if ld != "/" && !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return NewProcessId(pid.GetPhysicalAddress(), ld+name)
}

// URL 获取进程 Id 的 URL
func (pid *ProcessId) URL() *url.URL {
	if pid == nil {
		return zeroUrl
	}
	u := &url.URL{
		Scheme: "minotaur",
		Host:   pid.GetPhysicalAddress(),
		Path:   pid.GetLogicalAddress(),
	}
	return u
}

// Equal 比较两个进程 ID 是否相同
func (pid *ProcessId) Equal(id *ProcessId) bool {
	if pid == nil || id == nil {
		return false
	}
	if pid.GetPhysicalAddress() != id.GetPhysicalAddress() {
		return false
	}
	if pid.GetLogicalAddress() != id.GetLogicalAddress() {
		return false
	}
	return true
}

// Clone 克隆进程 ID
func (pid *ProcessId) Clone() *ProcessId {
	return &ProcessId{
		LogicalAddress:  pid.GetLogicalAddress(),
		PhysicalAddress: pid.GetPhysicalAddress(),
	}
}

// GetPhysicalAddress 加载进程 ID 的物理地址，在任何时候都应该通过该函数获取物理地址
func (pid *ProcessId) GetPhysicalAddress() PhysicalAddress {
	return pid.PhysicalAddress
}

// GetLogicalAddress 加载进程 ID 的逻辑地址，在任何时候都应该通过该函数获取逻辑地址
func (pid *ProcessId) GetLogicalAddress() LogicalAddress {
	return pid.LogicalAddress
}

var zeroUrl = &url.URL{}

func ClearProcessIdCache(pid *ProcessId) {
	pid.cache.Store(nil)
}

func LoadProcessIdCache(pid *ProcessId) *any {
	return pid.cache.Load()
}

func StoreProcessIdCache(pid *ProcessId, cache *any) {
	pid.cache.Store(cache)
}

// SetProcessIdProxy 设置进程 ID 的代理
//   - 该函数无法保证并发安全，仅支持在创建进程 ID 完成还未被使用之前设置
func SetProcessIdProxy(pid *ProcessId, proxy func(source *ProcessId) (redirect *ProcessId)) {
	pid.proxy = proxy
}

func GetProcessIdProxy(pid *ProcessId) func(source *ProcessId) (redirect *ProcessId) {
	return pid.proxy
}
