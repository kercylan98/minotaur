package cluster

import (
	"context"
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/engine/prc"
	prcv1 "github.com/kercylan98/minotaur/engine/prc/v1"
	"github.com/kercylan98/minotaur/engine/vivid"
	clusterv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/v1"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/log"
	"sync"
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
		nodes:  make(map[prc.PhysicalAddress]*Node),
	}

	config.WithShared(address)
	config.WithComponents(&component{cluster: system})

	system.ActorSystem = vivid.NewActorSystemWithConfiguration(config.ActorSystemConfiguration)

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

	// 去重种子节点
	seedNodes = collection.DeduplicateSlice(seedNodes)
	system.systemRef = system.ActorOfF(func() vivid.Actor {
		return newActorSystemActor(system, seedNodes)
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("cluster")
	})

	return system
}

type ActorSystem struct {
	*vivid.ActorSystem                           // 如果单独使用，那么一切行为将越过集群
	config             *ActorSystemConfiguration // 集群配置
	systemRef          vivid.ActorRef            // 集群 ActorSystem 的 Actor 引用
	nodeRWLock         sync.RWMutex
	nodes              map[prc.PhysicalAddress]*Node
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

func (sys *ActorSystem) getAvailableNodeWithFixedProvider(name string) *Node {
	sys.nodeRWLock.RLock()
	defer sys.nodeRWLock.RUnlock()

	var targets []*Node
	for _, node := range sys.nodes {
		if node.gossipNode.UserState.FixedActorProviders[name] {
			targets = append(targets, node)
		}
	}

	return sys.config.nodeBalancer.Select(targets)
}

// GetOnlyActor 获取一个集群内唯一的 Actor 引用
func (sys *ActorSystem) GetOnlyActor(name string) vivid.ActorRef {
	proxyRef := vivid.NewActorRef("", "")
	prcv1.SetProcessIdProxy(proxyRef, func(source *prcv1.ProcessId) (redirect *prcv1.ProcessId) {
		sys.nodeRWLock.RLock()
		nodes := collection.CloneMap(sys.nodes)
		sys.nodeRWLock.RUnlock()

		for _, node := range nodes {
			alive, exist := node.gossipNode.UserState.AliveOnlyActors[name]
			if !exist {
				continue
			}
			return alive.Ref
		}
		return nil
	})

	return proxyRef
}

func (sys *ActorSystem) SpawnFixedActor(name string, timeout ...time.Duration) (ref vivid.ActorRef, err error) {
	node := sys.getAvailableNodeWithFixedProvider(name)
	if node == nil {
		return nil, errors.New("no available node")
	}

	var result any
	if result, err = sys.ActorSystem.FutureAsk(node.nodeRef, &clusterv1.SpawnFixedActor{Name: name}, timeout...).Result(); err != nil {
		return
	}
	actorOfResult, ok := result.(*clusterv1.SpawnFixedActorResult)
	if !ok {
		return nil, fmt.Errorf("actor of result type error, expect %T, got %T, please check the cluster version", &clusterv1.SpawnFixedActorResult{}, result)
	}
	ref = actorOfResult.Ref

	return
}
