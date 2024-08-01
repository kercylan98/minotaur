package prc

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
	if pid.LogicalAddress != "/" && !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return &ProcessId{
		LogicalAddress:  pid.LogicalAddress + name,
		PhysicalAddress: pid.PhysicalAddress,
	}
}

// URL 获取进程 Id 的 URL
func (pid *ProcessId) URL() *url.URL {
	u := &url.URL{
		Scheme: "minotaur",
		Host:   pid.PhysicalAddress,
		Path:   pid.LogicalAddress,
	}
	return u
}

// Equal 比较两个进程 ID 是否相同
func (pid *ProcessId) Equal(id *ProcessId) bool {
	if pid.PhysicalAddress != id.PhysicalAddress {
		return false
	}
	if pid.LogicalAddress != id.LogicalAddress {
		return false
	}
	return true
}

func (pid *ProcessId) Clone() *ProcessId {
	return &ProcessId{
		LogicalAddress:  pid.LogicalAddress,
		PhysicalAddress: pid.PhysicalAddress,
	}
}
