package cluster

import (
	prcv1 "github.com/kercylan98/minotaur/engine/prc/v1"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip"
)

func newNode(system *ActorSystem, ctx vivid.ActorContext, gossipNode *gossip.Node) *Node {
	return &Node{
		system:     system,
		ctx:        ctx,
		gossipNode: gossipNode,
		nodeRef:    vivid.NewActorRef(gossipNode.Id.Ref.PhysicalAddress, "/user/cluster"),
	}
}

type Node struct {
	system     *ActorSystem
	ctx        vivid.ActorContext
	nodeRef    vivid.ActorRef
	gossipNode *gossip.Node
}

func (n *Node) GetId() string {
	return n.gossipNode.Id.Ref.PhysicalAddress
}

func (n *Node) GetWeight() int {
	return 10
}

func (n *Node) update(gossipNode *gossip.Node) {
	n.gossipNode = gossipNode
	if n.nodeRef.PhysicalAddress != gossipNode.Id.Ref.PhysicalAddress {
		n.nodeRef.PhysicalAddress = gossipNode.Id.Ref.PhysicalAddress
		prcv1.ClearProcessIdCache(n.nodeRef)
	}
}
