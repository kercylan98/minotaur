package prc

import (
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"net/url"
	"os"
	"strings"
)

func NewProcessId(physicalAddress PhysicalAddress, logicalAddress LogicalAddress) *ProcessId {
	return NewClusterProcessId(charproc.None, physicalAddress, logicalAddress)
}

func NewClusterProcessId(clusterName ClusterName, physicalAddress PhysicalAddress, logicalAddress LogicalAddress) *ProcessId {
	return &ProcessId{
		PhysicalPid:     uint32(os.Getpid()),
		LogicalAddress:  logicalAddress,
		PhysicalAddress: physicalAddress,
		NetworkProtocol: "",
		ClusterName:     clusterName,
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

// URL 获取进程 Id 的 URL
func (pid *ProcessId) URL() *url.URL {
	u := &url.URL{
		Scheme: "minotaur",
		Host:   pid.PhysicalAddress,
		Path:   pid.LogicalAddress,
	}
	if pid.ClusterName != charproc.None {
		u.User = url.User(pid.ClusterName)
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
	if pid.ClusterName != id.ClusterName {
		return false
	}
	return true
}
