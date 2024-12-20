package cluster

import (
	"context"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip"
	clusterv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/v1"
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
	// 订阅集群收敛
	ctx.Subscribe(gossip.TopicNodeConverged)
	// 启动 gossip actor
	a.gossipRef = ctx.ActorOfF(func() vivid.Actor {
		return gossip.NewGossiperActor(a.seedNodes)
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
	defer a.system.nodeRWLock.Unlock()

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
