package vivid

import "github.com/kercylan98/minotaur/engine/prc"

type ActorRef = *prc.ProcessId

// NewActorRef 通过特定的物理地址和逻辑地址，创建一个指向特定 ActorSystem 的 ActorRef
func NewActorRef(physicalAddress prc.PhysicalAddress, logicAddress prc.LogicalAddress) ActorRef {
	return prc.NewProcessId(physicalAddress, logicAddress)
}
