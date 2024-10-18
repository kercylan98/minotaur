package gossip_test

import (
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip"
	"testing"
)

func TestHashRing_GetNeighbours(t *testing.T) {
	var ring = gossip.NewHashRing(5)
	ring.AddNode("node1")
	ring.AddNode("node2")
	ring.AddNode("node3")
	ring.AddNode("node4")
	ring.AddNode("node5")

	t.Log(ring.GetNeighbours("node1", 1))
	t.Log(ring.GetNeighbours("node1", 2))
	t.Log(ring.GetNeighbours("node1", 3))
	t.Log(ring.GetNeighbours("node1", 4))
	t.Log(ring.GetNeighbours("node1", 5))
}
