package prc

import prcv1 "github.com/kercylan98/minotaur/engine/prc/v1"

func NewProcessId(physicalAddress PhysicalAddress, logicalAddress LogicalAddress) *ProcessId {
	return prcv1.NewProcessId(physicalAddress, logicalAddress)
}
