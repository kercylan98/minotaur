package cluster

import (
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip"
	"sort"
)

func GetAliveNodeWithASCLLAsc(nodes []*gossip.Node) *gossip.Node {
	if len(nodes) == 0 {
		return nil
	}

	if len(nodes) == 1 {
		return nodes[0]
	}

	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]

		if a.Status != b.Status {
			return a.Status == gossip.NodeStatusAlive
		}

		return a.Id.Ref.PhysicalAddress < b.Id.Ref.PhysicalAddress
	})

	return nodes[0]
}

func GetAliveNodeWithLaunchTimeAsc(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}

	if len(nodes) == 1 {
		return nodes[0]
	}
	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]

		if a.gossipNode.Status != b.gossipNode.Status {
			return a.gossipNode.Status == gossip.NodeStatusAlive
		}

		return a.gossipNode.LaunchTimestampMillis < b.gossipNode.LaunchTimestampMillis
	})

	return nodes[0]
}
