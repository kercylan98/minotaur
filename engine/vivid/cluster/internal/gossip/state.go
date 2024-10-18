package gossip

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/log"
	"sort"
)

func newState(ctx vivid.ActorContext, actor *GossiperActor) *State {
	vc := newVectorClock()
	vc.Increment(ctx.Ref().PhysicalAddress) // 版本初始 1

	node := &Node{
		Id:     newNodeId(ctx.Ref()),
		Status: GossipNodeStatus_GNS_Joining,
		Vc:     vc,
	}
	state := &State{
		ctx:   ctx,
		actor: actor,
		gossip: &Gossip{
			Members: []*Node{node},
			Seen:    []*NodeId{node.Id},
		},
		node: node,
	}

	return state
}

type State struct {
	ctx    vivid.ActorContext
	actor  *GossiperActor
	gossip *Gossip
	node   *Node
}

// Increment VectorClock 自增 1
func (s *State) Increment() {
	s.node.Vc.Increment(s.ctx.Ref().PhysicalAddress)
}

// AddMember 向集群中添加新节点，并相应更新 Gossip 状态和 VectorClock。
func (s *State) AddMember(node *Node) {
	// 检查新成员是否已存在于成员列表中
	for _, member := range s.gossip.Members {
		if member.Id.Ref.PhysicalAddress == node.Id.Ref.PhysicalAddress {
			return // 如果已存在，不做任何操作
		}
	}

	// 添加新节点到成员列表
	s.gossip.Members = append(s.gossip.Members, node)

	// 状态更新，重置 Seen 列表为仅含自身
	s.gossip.Seen = []*NodeId{s.node.Id}

	// 自增节点的 VectorClock，以反映新节点的加入
	s.Increment()
}

// MarkSeen 标记某个节点已经看到当前 Gossip 信息。
func (s *State) MarkSeen(nodeId *NodeId) {
	// 检查该节点是否已经在 Seen 列表中
	for _, seenNode := range s.gossip.Seen {
		if seenNode.Ref.PhysicalAddress == nodeId.Ref.PhysicalAddress {
			return // 如果已在 Seen 列表中，不做重复操作
		}
	}

	// 将节点加入 Seen 列表
	s.gossip.Seen = append(s.gossip.Seen, nodeId)
}

// MergeGossip 合并收到的 Gossip 数据，并更新当前节点的状态。
func (s *State) MergeGossip(gossiped *Gossiped) {
	// 如果接收到的 Gossip 版本更新，则进行合并
	s.gossip = gossiped.Gossip
	if s.node.Vc.CompareTo(gossiped.GossiperVersion) != VectorClockOrdering_VCO_Same {
		s.node.Vc.Merge(gossiped.GossiperVersion)
		for _, member := range s.gossip.Members {
			if member.Id.PhysicalAddressEqual(s.node.Id) {
				member.Vc = s.node.Vc
				s.node = member // 确保指针一致
			}
		}

		// 状态更新，重置 Seen 列表为仅含自身
		// 这里不改变版本，改变版本将导致永远无终止
		s.gossip.Seen = []*NodeId{}
	} else {
		for _, member := range s.gossip.Members {
			if member.Id.PhysicalAddressEqual(s.node.Id) {
				member.Vc = s.node.Vc
				s.node = member // 确保指针一致
			}
		}
	}
}

// CalcLeaderNode 计算当前集群的领导者节点。
func (s *State) CalcLeaderNode() *Node {
	// 仅有一个节点，那么就是它
	if len(s.gossip.Members) == 1 {
		return s.gossip.Members[0]
	}

	// 可用的第一个节点为领导节点
	sort.Slice(s.gossip.Members, func(i, j int) bool {
		a, b := s.gossip.Members[i], s.gossip.Members[j]

		// 先比较状态，Alive 靠前
		if a.Status != b.Status {
			return a.Status == GossipNodeStatus_GNS_Alive
		}

		// 如果状态相同，按 PhysicalAddress 升序排序
		return a.Id.Ref.PhysicalAddress < b.Id.Ref.PhysicalAddress
	})

	return s.gossip.Members[0]
}

// GossipUpdate 用于在集群中传播 Gossip 数据，并确保所有节点状态的一致性。
func (s *State) GossipUpdate() {
	// 建立 seen 映射
	seenMap := make(map[string]struct{})
	for _, id := range s.gossip.Seen {
		seenMap[id.Ref.PhysicalAddress] = struct{}{}
	}

	// 保留非 seen 中状态正常的节点
	var targets []*Node
	for _, member := range s.gossip.Members {
		if member.Id.Ref.PhysicalAddress == s.node.Id.Ref.PhysicalAddress {
			continue
		}

		if _, ok := seenMap[member.Id.Ref.PhysicalAddress]; ok {
			continue
		}

		switch member.Status {
		case GossipNodeStatus_GNS_Joining, GossipNodeStatus_GNS_Alive, GossipNodeStatus_GNS_Leaving, GossipNodeStatus_GNS_Exit:
			targets = append(targets, member)
		}
	}

	// 选择最多一定数量的节点进行传播
	var num = 5
	if len(targets) < num {
		num = len(targets)
	}

	if num == 0 {
		return // 没有节点可发送 gossip，直接返回
	}

	for _, node := range collection.ChooseRandomSliceElementN(targets, num) {
		s.ctx.Ask(node.Id.Ref, &Gossiped{
			Gossip:          s.gossip,
			GossiperVersion: s.node.Vc,
		})
		s.actor.logger.Debug("cluster", log.String("gossip", fmt.Sprintf("%s -> %s", s.node.Id.Ref.URL().String(), node.Id.Ref.URL().String())))
	}
}
