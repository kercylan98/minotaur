package cluster

import (
	"context"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip"
	clusterv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/v1"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"sort"
)

type (
	actorSystemActorExitMessage struct {
		cancel context.CancelFunc
	}
)

func newActorSystemActor(system *ActorSystem, seedNodes []prc.PhysicalAddress) *actorSystemActor {
	return &actorSystemActor{
		system:    system,
		seedNodes: seedNodes,
	}
}

type actorSystemActor struct {
	system    *ActorSystem
	seedNodes []prc.PhysicalAddress // 种子节点
	gossipRef vivid.ActorRef        // gossip actor ref
	exited    context.CancelFunc    // 退出信号
	leaderRef vivid.ActorRef        // 集群当前的领导者引用
	nodeState *gossip.NodeState     // 当前节点状态
}

func (a *actorSystemActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		a.onLaunch(ctx)
	case vivid.ActorRef:
		a.onLeaderChanged(ctx, m)
	case *actorSystemActorExitMessage:
		a.onActorSystemActorExitMessage(ctx, m)
	case *gossip.ActorClusterExitingMessage:
		a.onGossipActorClusterExitingMessage(ctx, m)
	case *gossip.ActorClusterExitedMessage:
		a.onGossipActorClusterExitedMessage(ctx, m)
	case gossip.ClusterConvergedEvent: // 集群收敛消息
		a.onGossipClusterConvergedEvent(ctx, m)
	case *clusterv1.SpawnFixedActor:
		a.onActorOf(ctx, m)
	}
}

func (a *actorSystemActor) onLaunch(ctx vivid.ActorContext) {
	// 初始化节点状态
	a.nodeState = &gossip.NodeState{
		FixedActorProviders: collection.ConvertMapValuesToBoolMap(vivid.GetFixedActorProviders(a.system.ActorSystem)),
		OnlyActorProviders:  collection.ConvertMapValuesToBoolMap(a.system.config.onlyActorProviders),
	}
	// 订阅集群收敛
	ctx.Subscribe(gossip.TopicNodeConverged)
	// 启动 gossip actor
	a.gossipRef = ctx.ActorOfF(func() vivid.Actor {
		return gossip.NewGossiperActor(a.nodeState, a.seedNodes)
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("gossip")
	})
}

func (a *actorSystemActor) onActorSystemActorExitMessage(ctx vivid.ActorContext, m *actorSystemActorExitMessage) {
	a.exited = m.cancel
	ctx.Tell(a.gossipRef, &gossip.ActorLeaveClusterMessage{})
}

func (a *actorSystemActor) onGossipActorClusterExitingMessage(ctx vivid.ActorContext, m *gossip.ActorClusterExitingMessage) {
	// 善后工作
}

func (a *actorSystemActor) onGossipActorClusterExitedMessage(ctx vivid.ActorContext, m *gossip.ActorClusterExitedMessage) {
	a.exited()
}

func (a *actorSystemActor) onLeaderChanged(ctx vivid.ActorContext, m vivid.ActorRef) {
	a.leaderRef = m
}

func (a *actorSystemActor) onGossipClusterConvergedEvent(ctx vivid.ActorContext, m gossip.ClusterConvergedEvent) {
	// 保留可达节点
	var activeList = make(map[prc.PhysicalAddress]*gossip.Node)
	for _, node := range m {
		switch node.Status {
		case gossip.NodeStatusAlive:
			activeList[node.Id.Ref.GetPhysicalAddress()] = node
		}
	}

	a.system.nodeRWLock.Lock()
	// 移除陈旧
	for key, node := range a.system.nodes {
		_, exist := activeList[node.GetId()]
		if exist {
			continue
		}
		delete(a.system.nodes, key)
	}

	// 补充新增
	for _, node := range activeList {
		_, exist := a.system.nodes[node.Id.Ref.GetPhysicalAddress()]
		if exist {
			continue
		}
		create := newNode(a.system, ctx, node)
		a.system.nodes[node.Id.Ref.GetPhysicalAddress()] = create
	}
	a.system.nodeRWLock.Unlock()

	// 整理数据
	var aliveOnlyActors = make(map[string]int64)
	for _, node := range m {
		if node.Status == gossip.NodeStatusAlive {
			for name, aliveTime := range node.UserState.AliveOnlyActors {
				aliveOnlyActors[name] = aliveTime
			}
		}
	}

	var onlyNodes = make(map[string][]*gossip.Node)
	for _, node := range m {
		for name := range node.UserState.OnlyActorProviders {
			onlyNodes[name] = append(onlyNodes[name], node)
		}
	}

	for name, nodes := range onlyNodes {
		if aliveOnlyActors[name] > 0 {
			continue
		}

		// 在选择目标中生成唯一 Actor
		selected := CalcFirstNode(nodes)
	}
}

func (a *actorSystemActor) onActorOf(ctx vivid.ActorContext, m *clusterv1.SpawnFixedActor) {
	ref, err := vivid.SpawnActorFromFixedProvider(ctx.System(), ctx, m.Name)
	if err != nil {
		ctx.Reply(ref)
		return
	}
	ctx.Reply(&clusterv1.SpawnFixedActorResult{
		Ref: ref,
	})
}

func CalcFirstNode(nodes []*gossip.Node) *gossip.Node {
	if len(nodes) == 0 {
		return nil
	}

	// 仅有一个节点，那么就是它
	if len(nodes) == 1 {
		return nodes[0]
	}

	// 可用的第一个节点为领导节点
	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]

		// 先比较状态，Alive 靠前
		if a.Status != b.Status {
			return a.Status == gossip.NodeStatusAlive
		}

		// 如果状态相同，按 PhysicalAddress 升序排序
		return a.Id.Ref.PhysicalAddress < b.Id.Ref.PhysicalAddress
	})

	return nodes[0]
}
