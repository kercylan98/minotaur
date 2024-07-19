package prc

import (
	"os"
	"strings"
)

func NewProcessId(address LogicalAddress, physicalAddress PhysicalAddress) *ProcessId {
	return &ProcessId{
		PhysicalPid:     uint32(os.Getpid()),
		LogicalAddress:  address,
		PhysicalAddress: physicalAddress,
		NetworkProtocol: "",
		ClusterName:     "",
	}
}

// Derivation 衍生一个新的进程 Id
func (pid *ProcessId) Derivation(name string) *ProcessId {
	if pid.LogicalAddress != "/" && !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return &ProcessId{
		PhysicalPid:     pid.PhysicalPid,
		LogicalAddress:  pid.LogicalAddress + name,
		PhysicalAddress: pid.PhysicalAddress,
		NetworkProtocol: pid.NetworkProtocol,
		ClusterName:     pid.ClusterName,
	}
}
