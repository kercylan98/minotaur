package cluster

import "github.com/hashicorp/memberlist"

type NodeInfo struct {
	*memberlist.Node
	metadata *Metadata
}

func (n *NodeInfo) GetId() string {
	return n.Address()
}

func (n *NodeInfo) GetWeight() int {
	return int(n.metadata.Weight)
}
