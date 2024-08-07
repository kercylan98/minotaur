package prc

import (
	"net/url"
	"strings"
)

var zeroUrl = &url.URL{}

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
