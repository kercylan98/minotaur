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
	"google.golang.org/protobuf/proto"
	"time"
)

type (
	actorSystemActorExitMessage struct {
		cancel context.CancelFunc
	}
)

func newActorSystemActor(system *ActorSystem, seedNodes []prc.PhysicalAddress, gossipRefHandler func(ref vivid.ActorRef)) *actorSystemActor {
	return &actorSystemActor{
		system:           system,
		seedNodes:        seedNodes,
		gossipRefHandler: gossipRefHandler,
	}
}

type actorSystemActor struct {
	system           *ActorSystem
	seedNodes        []prc.PhysicalAddress    // 种子节点
	gossipRef        vivid.ActorRef           // gossip actor ref
	exited           context.CancelFunc       // 退出信号
	leaderRef        vivid.ActorRef           // 集群当前的领导者引用
	gossipRefHandler func(ref vivid.ActorRef) // 集群 Gossip 引用处理器
	nodeId           *gossip.NodeId           // 集群自身节点 ID
}

func (a *actorSystemActor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *gossip.NodeId:
		a.nodeId = m
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
	}
}

func (a *actorSystemActor) onLaunch(ctx vivid.ActorContext) {
	// 订阅集群收敛
	ctx.Subscribe(gossip.TopicNodeConverged)
	// 启动 gossip actor
	a.gossipRef = ctx.ActorOfF(func() vivid.Actor {
		return gossip.NewGossiperActor(a.seedNodes,
			gossip.FunctionalUserDataProvider(func() (key string, value proto.Message) {
				return clusterv1.UserDataKey_USER_DATA_KEY_ONLY_ACTOR_PROVIDERS.String(), &clusterv1.OnlyActorProviders{
					OnlyActorProviders: collection.ConvertMapValuesToBoolMap(a.system.config.onlyActorProviders),
				}
			}),
		)
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("gossip")
	})
	a.gossipRefHandler(a.gossipRef)
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

	// 构建集群内唯一 Actor
	a.onGossipClusterConvergedProcessOnlyActors(ctx)
}

func (a *actorSystemActor) onGossipClusterConvergedProcessOnlyActors(ctx vivid.ActorContext) {
	var aliveActors = make(map[string]*clusterv1.AliveOnlyActorInfos_AliveOnlyActorInfo) // 集群内存活的 Actor 名称及其信息
	var expiredActors []*clusterv1.AliveOnlyActorInfos_AliveOnlyActorInfo                // 自己过期的 Actor
	var requiredActors = make(map[string]struct{})                                       // 集群内所需的唯一 Actor 名称

	aliveNodes, err := vivid.FutureAsk[*gossip.GetAvailableNodesResponseMessage](ctx, a.gossipRef, gossip.GetAvailableNodesRequest).Result()
	if err != nil {
		ctx.System().Logger().Error("cluster", log.String("event", "get available nodes failed"), log.String("error", err.Error()))
		return
	}

	// 获取可用节点存活的 Actor 信息
	aliveActorNodeUserDataMap, err := gossip.GetUserData[*clusterv1.AliveOnlyActorInfos](ctx, a.gossipRef, clusterv1.UserDataKey_USER_DATA_KEY_ALIVE_ONLY_ACTORS.String(), aliveNodes.NodeIds...)
	if err != nil {
		ctx.System().Logger().Error("cluster", log.String("event", "get-user-data"),
			log.String("key", clusterv1.UserDataKey_USER_DATA_KEY_ALIVE_ONLY_ACTORS.String()),
			log.Err(err))
		return
	}

	// 标记是否有变化
	var changed bool

	// 分析自身节点所需的 Actor
	for name := range a.system.config.onlyActorProviders {
		requiredActors[name] = struct{}{}
	}

	// 分析过期或存活的 Actor
	for _, infos := range aliveActorNodeUserDataMap {
		for actorName, info := range infos.AliveOnlyActors {
			currAliveActor, exist := aliveActors[actorName]
			if exist {
				// 创建时间小于当前时间，自身为过期的 Actor
				if info.GenerateTime < currAliveActor.GenerateTime {
					// 如果是自身的 Actor，那么标记
					if info.NodeId.Equal(a.nodeId) {
						expiredActors = append(expiredActors, info)
					}
					continue
				} else {
					// 创建时间大于当前时间，自身为存活的 Actor，当前存活的标记过期
					aliveActors[actorName] = info
					if currAliveActor.NodeId.Equal(a.nodeId) {
						expiredActors = append(expiredActors, currAliveActor)
					}
				}
			} else {
				// 不存在，标记自身为存活的 Actor
				aliveActors[actorName] = info
			}
		}
	}

	// 生成集群内缺乏的 Actor
	for actorName := range requiredActors {
		if _, exist := aliveActors[actorName]; exist {
			continue
		}

		// 计算合适的节点，如果节点是自身，那么创建
		var minNodeIdKey gossipv1.NodeKey
		for _, id := range aliveNodes.NodeIds {
			key := id.Key()
			if minNodeIdKey == "" || key < minNodeIdKey {
				minNodeIdKey = key
			}
		}
		nodeIdKey := a.nodeId.Key()
		if minNodeIdKey != nodeIdKey {
			continue
		}

		provider := a.system.config.onlyActorProviders[actorName]
		// 创建 Actor
		ref := ctx.ActorOf(vivid.FunctionalActorProvider(func() vivid.Actor {
			return provider.ProvideActor()
		}), provider.ProvideConfigurator())

		infos, exist := aliveActorNodeUserDataMap[nodeIdKey]
		if !exist {
			infos = &clusterv1.AliveOnlyActorInfos{}
			aliveActorNodeUserDataMap[nodeIdKey] = infos
		}
		if infos.AliveOnlyActors == nil {
			infos.AliveOnlyActors = make(map[string]*clusterv1.AliveOnlyActorInfos_AliveOnlyActorInfo)
		}
		info := &clusterv1.AliveOnlyActorInfos_AliveOnlyActorInfo{
			Name:         actorName,
			Ref:          ref,
			GenerateTime: time.Now().UnixMilli(),
		}
		infos.AliveOnlyActors[actorName] = info
		changed = true

		ctx.System().Logger().Debug("cluster", log.String("generate", "only_actor"), log.String("name", actorName))
	}

	// 移除自己过期的 Actor
	for _, info := range expiredActors {
		ctx.Terminate(info.Ref, true)
		delete(aliveActorNodeUserDataMap[info.NodeId.Key()].AliveOnlyActors, info.Name)

		ctx.System().Logger().Debug("cluster", log.String("expire", "only_actor"), log.String("name", info.Name), log.String("ref", info.Ref.URL().String()))
		changed = true
	}

	if changed {
		if err = gossip.SetUserData(ctx, a.gossipRef, clusterv1.UserDataKey_USER_DATA_KEY_ALIVE_ONLY_ACTORS.String(), aliveActorNodeUserDataMap[a.nodeId.Key()]); err != nil {
			ctx.System().Logger().Error("cluster", log.String("event", "set-user-data"),
				log.String("key", clusterv1.UserDataKey_USER_DATA_KEY_ALIVE_ONLY_ACTORS.String()),
				log.Err(err))
		}
	}
}
