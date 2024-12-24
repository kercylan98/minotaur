package cluster

import (
	"context"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip"
	gossipv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip/v1"
	clusterv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/v1"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"sort"
	"time"
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
	var changed bool
	var setChanged = func() {
		changed = true
	}

	// 筛选存活节点更新节点列表
	a.onGossipClusterConvergedFilterAliveNodes(ctx, m)

	// 构建集群内唯一 Actor
	a.onGossipClusterConvergedProcessOnlyActors(ctx, m, setChanged)

	// 状态变更，继续收敛
	if changed {
		ctx.Tell(a.gossipRef, gossip.OnStateChanged)
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

func (a *actorSystemActor) onGossipClusterConvergedFilterAliveNodes(ctx vivid.ActorContext, m gossip.ClusterConvergedEvent) {
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

func (a *actorSystemActor) onGossipClusterConvergedProcessOnlyActors(ctx vivid.ActorContext, m gossip.ClusterConvergedEvent, changed func()) {
	// 过滤存活 Actor 并且记录信息
	var aliveOnlyActors = make(map[string][]*gossip.AliveOnlyActorInfo)
	for _, node := range m {
		if node.Status == gossip.NodeStatusAlive {
			for name, info := range node.UserState.AliveOnlyActors {
				aliveOnlyActors[name] = append(aliveOnlyActors[name], info)
			}
		}
	}

	// 分析集群内所需的唯一 Actor
	var onlyNodes = make(map[string]struct{})
	for _, node := range m {
		for name := range node.UserState.OnlyActorProviders {
			onlyNodes[name] = struct{}{}
		}
	}

	// 如果集群内缺乏唯一 Actor，且自身节点满足要求，那么创建
	for name := range onlyNodes {
		if len(aliveOnlyActors[name]) > 0 {
			continue
		}

		provider, exist := a.system.config.onlyActorProviders[name]
		if !exist {
			continue
		}

		selected := GetAliveNodeWithLaunchTimeAsc(m)
		if selected == nil || selected.Id.Ref.GetPhysicalAddress() != ctx.PhysicalAddress() {
			continue
		}

		// 创建 Actor
		ref := ctx.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return provider.ProvideActor()
		}), provider.ProvideConfigurator())

		if a.nodeState.AliveOnlyActors == nil {
			a.nodeState.AliveOnlyActors = make(map[string]*gossipv1.AliveOnlyActorInfo)
		}
		info := &gossip.AliveOnlyActorInfo{
			Ref:          ref,
			GenerateTime: time.Now().UnixMilli(),
		}
		a.nodeState.AliveOnlyActors[name] = info
		aliveOnlyActors[name] = append(aliveOnlyActors[name], info)
		changed()
	}

	// 仅保留最新的
	for key, infos := range aliveOnlyActors {
		if len(infos) > 1 {
			sort.Slice(infos, func(i, j int) bool {
				return infos[i].GenerateTime > infos[j].GenerateTime
			})
			for i := 1; i < len(infos); i++ {
				ctx.Terminate(infos[i].Ref, true)
			}
			aliveOnlyActors[key] = infos
			changed()
		}
	}

}
