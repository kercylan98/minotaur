package gossip

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"time"
)

func newNodeId(ref vivid.ActorRef) *NodeId {
	return &NodeId{
		Ref:  ref,
		Guid: time.Now().UnixMicro(),
	}
}

func (ni *NodeId) Equal(other *NodeId) bool {
	return ni.Guid == other.Guid && ni.Ref.PhysicalAddress == other.Ref.PhysicalAddress
}

func (ni *NodeId) PhysicalAddressEqual(other *NodeId) bool {
	return ni.Ref.PhysicalAddress == other.Ref.PhysicalAddress
}
