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
