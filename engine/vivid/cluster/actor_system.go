package cluster

import (
	"context"
	"errors"
	"github.com/kercylan98/minotaur/engine/prc"
	prcv1 "github.com/kercylan98/minotaur/engine/prc/v1"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip"
	clusterv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/v1"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/protobuf/proto"
	"sync/atomic"
	"time"
)

func init() {
	vivid.RegisterFutureAskType(func(ctx any) *vivid.ActorSystem {
		as, _ := ctx.(*ActorSystem)
		return as.ActorSystem
	})
}

func NewActorSystem(address prc.PhysicalAddress, configurator ...ActorSystemConfigurator) *ActorSystem {
	return NewFixedSeedNodesActorSystem(address, nil, configurator...)
}

// NewFixedSeedNodesActorSystem 创建一个固定种子节点的集群 ActorSystem
func NewFixedSeedNodesActorSystem(address prc.PhysicalAddress, seedNodes []prc.PhysicalAddress, configurator ...ActorSystemConfigurator) *ActorSystem {
	config := newActorSystemConfiguration()
	for _, c := range configurator {
		c.Configure(config)
	}
	seedNodes = append(seedNodes, config.seeds...)

	system := &ActorSystem{
		config: config,
	}

	config.WithShared(address)
	config.WithComponents(&component{cluster: system})

	system.ActorSystem = vivid.NewActorSystemWithConfiguration(config.ActorSystemConfiguration)

	if config.seedProvider != nil {
		if err := toolkit.RetryByExponentialBackoff(func() error {
			provideSeeds, err := config.seedProvider.Provide()
			if err != nil {
				system.Logger().Error("cluster", log.String("status", "backoff provide seeds failed"), log.Err(err))
			} else {
				seedNodes = append(seedNodes, provideSeeds...)
			}
			return err
		}, 64, time.Second, time.Minute, 2, 0.5); err != nil {
			panic(err)
		}
	}

	// 去重种子节点
	seedNodes = collection.DeduplicateSlice(seedNodes)
	system.systemRef = system.ActorOfF(func() vivid.Actor {
		return newActorSystemActor(system, seedNodes, func(ref vivid.ActorRef) {
			system.gossipRef = ref
		})
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("cluster")
	})

	return system
}

type ActorSystem struct {
	*vivid.ActorSystem                           // 如果单独使用，那么一切行为将越过集群
	config             *ActorSystemConfiguration // 集群配置
	systemRef          vivid.ActorRef            // 集群 ActorSystem 的 Actor 引用
	gossipRef          vivid.ActorRef            // 集群 ActorSystem 的 gossip Actor 引用
}

func (sys *ActorSystem) onShutdown() {
	sys.Logger().Info("cluster", log.String("status", "shutdown in progress"))

	ctx, cancel := context.WithTimeout(context.Background(), sys.config.shutdownTimeout)
	sys.Tell(sys.systemRef, &actorSystemActorExitMessage{cancel: cancel})
	<-ctx.Done()

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		sys.Logger().Error("cluster", log.String("status", "shutdown timeout"))
	} else {
		sys.Logger().Info("cluster", log.String("status", "shutdown success"))
	}
}

// SetUserData 设置集群内用户数据，该数据可以在集群内传播
func (sys *ActorSystem) SetUserData(key string, value proto.Message) error {
	return gossip.SetUserData(sys.Context(), sys.gossipRef, key, value)
}

// GetUserData 获取集群内用户数据，如果 nodeIds 不指定，将获取所有节点的用户数据
func (sys *ActorSystem) GetUserData(key string, nodeIds ...*NodeId) (map[NodeKey]proto.Message, error) {
	return gossip.GetUserData[proto.Message](sys.Context(), sys.gossipRef, key, nodeIds...)
}

// GetUserData 获取集群中的用户数据，当 nodeIds 不存在时，将获取所有节点
func GetUserData[M proto.Message](sys *ActorSystem, key string, nodeIds ...*NodeId) (map[NodeKey]M, error) {
	return gossip.GetUserData[M](sys.Context(), sys.gossipRef, key, nodeIds...)
}

// SetUserData 设置集群中的用户数据
func SetUserData(sys *ActorSystem, key string, value proto.Message) error {
	return gossip.SetUserData(sys.Context(), sys.gossipRef, key, value)
}

// GetOnlyActor 获取一个集群内唯一的 Actor 引用
func (sys *ActorSystem) GetOnlyActor(name string) vivid.ActorRef {
	var proxyRef = vivid.NewActorRef("", "")
	var cachePointer atomic.Pointer[prc.ProcessId]
	prcv1.SetProcessIdProxy(proxyRef, func(source *prcv1.ProcessId) (redirect *prcv1.ProcessId) {
		cache := cachePointer.Load()
		if cache != nil {
			if _, err := vivid.Ping(sys.ActorSystem, cache); err != nil {
				cache = nil
				cachePointer.Store(nil)
			}
		}
		if cache == nil {
			// 获取可用节点
			aliveActorNodeUserDataMap, err := gossip.GetUserData[*clusterv1.AliveOnlyActorInfos](sys.Context(), sys.gossipRef, clusterv1.UserDataKey_USER_DATA_KEY_ALIVE_ONLY_ACTORS.String())
			if err != nil {
				sys.Logger().Error("cluster", log.String("event", "get-user-data"),
					log.String("key", clusterv1.UserDataKey_USER_DATA_KEY_ALIVE_ONLY_ACTORS.String()),
					log.Err(err))
			}
			for _, infos := range aliveActorNodeUserDataMap {
				alive, exist := infos.AliveOnlyActors[name]
				if !exist {
					continue
				}
				cache = alive.Ref
				cachePointer.Store(cache)
			}
		}

		return cache
	})

	return proxyRef
}
