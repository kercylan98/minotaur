package cluster

import (
	"context"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip"
	gossipv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip/v1"
	clusterv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/v1"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/log"
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

	a.system.nodeRWLock.Lock()
	defer a.system.nodeRWLock.Unlock()

	// 筛选存活节点更新节点列表
	a.onGossipClusterConvergedFilterAliveNodes(ctx, m)

	// 构建集群内唯一 Actor
	a.onGossipClusterConvergedProcessOnlyActors(ctx, setChanged)

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
	var activeList = make(map[prc.PhysicalAddress]*Node)
	for _, node := range m {
		switch node.Status {
		case gossip.NodeStatusAlive:
			activeList[node.Id.Ref.GetPhysicalAddress()] = newNode(a.system, ctx, node)
		}
	}
	a.system.nodes = activeList
}

func (a *actorSystemActor) onGossipClusterConvergedProcessOnlyActors(ctx vivid.ActorContext, changed func()) {
	var aliveActors = make(map[string]*gossip.AliveOnlyActorInfo)    // 集群内存活的 Actor 名称及其信息
	var expiredActors = make(map[*Node][]*gossip.AliveOnlyActorInfo) // 集群内过期的 Actor
	var requiredActors = make(map[string][]*Node)                    // 集群内所需的唯一 Actor 名称及能提供的节点
	for _, node := range a.system.nodes {
		// 分析集群内所需的 Actor
		for name := range node.gossipNode.UserState.OnlyActorProviders {
			requiredActors[name] = append(requiredActors[name], node)
		}

		//ctx.System().Logger().Debug("cluster", log.String("node", node.nodeRef.URL().String()),
		//	log.Any("required", collection.ConvertMapKeysToSlice(node.gossipNode.UserState.OnlyActorProviders)),
		//	log.Any("alive-only", node.gossipNode.UserState.AliveOnlyActors))

		// 获取存活的 Actor
		for actorName, info := range node.gossipNode.UserState.AliveOnlyActors {
			currAliveActor, exist := aliveActors[actorName]
			if exist {
				if info.GenerateTime < currAliveActor.GenerateTime {
					// 创建时间小于当前时间，自身为过期的 Actor
					expiredActors[node] = append(expiredActors[node], info)
					continue
				} else {
					// 创建时间大于当前时间，自身为存活的 Actor，当前存活的标记过期
					aliveActors[actorName] = info
					expiredActors[node] = append(expiredActors[node], currAliveActor)
				}
			} else {
				// 不存在，标记自身为存活的 Actor
				aliveActors[actorName] = info
			}
		}
	}

	// 生成集群内缺乏的 Actor
	for actorName, nodes := range requiredActors {
		if _, exist := aliveActors[actorName]; exist {
			continue
		}

		// 计算合适的节点
		selected := GetAliveNodeWithLaunchTimeAsc(nodes)
		if selected == nil || selected.gossipNode.Id.Ref.GetPhysicalAddress() != ctx.PhysicalAddress() {
			if selected == nil {
				ctx.System().Logger().Debug("cluster", log.String("generate_only_actor", actorName), log.String("selected", "no available node"))
			} else {
				ctx.System().Logger().Debug("cluster", log.String("generate_only_actor", actorName), log.String("selected", selected.gossipNode.Id.Ref.URL().String()))
				//ctx.System().Logger().Debug("cluster", log.String("generate_only_actor", actorName), log.String("selected", selected.gossipNode.Id.Ref.URL().String()),
				//	log.Any("contestants", collection.MappingFromSlice(nodes, func(value *Node) *gossip.Node {
				//		return value.gossipNode
				//	})))
			}
			continue
		}

		provider := a.system.config.onlyActorProviders[actorName]
		// 创建 Actor
		ref := ctx.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return provider.ProvideActor()
		}), provider.ProvideConfigurator())

		if selected.gossipNode.UserState.AliveOnlyActors == nil {
			selected.gossipNode.UserState.AliveOnlyActors = make(map[string]*gossipv1.AliveOnlyActorInfo)
		}
		info := &gossip.AliveOnlyActorInfo{
			Name:         actorName,
			Ref:          ref,
			GenerateTime: time.Now().UnixMilli(),
		}
		selected.gossipNode.UserState.AliveOnlyActors[actorName] = info
		aliveActors[actorName] = info
		changed()

		ctx.System().Logger().Debug("cluster", log.String("generate", "only_actor"), log.String("name", actorName))
	}

	// 移除集群内过期的 Actor
	for node, infos := range expiredActors {
		for _, info := range infos {
			ctx.Terminate(info.Ref, true)
			delete(node.gossipNode.UserState.AliveOnlyActors, info.Name)

			ctx.System().Logger().Debug("cluster", log.String("expire", "only_actor"), log.String("name", info.Name), log.String("ref", info.Ref.URL().String()))
		}
		changed()
	}

}
