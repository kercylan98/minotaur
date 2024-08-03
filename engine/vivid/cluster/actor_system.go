package cluster

import (
	"fmt"
	"github.com/hashicorp/memberlist"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/cm"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/log"
	"google.golang.org/protobuf/proto"
	"net"
	"time"
)

func init() {
	vivid.RegisterFutureAskType(func(ctx any) *vivid.ActorSystem {
		if as, ok := ctx.(*ActorSystem); ok {
			return as.ActorSystem
		}
		return nil
	})
}

func NewActorSystem(sharedAddress, bindAddress prc.PhysicalAddress, configurator ...ActorSystemConfigurator) *ActorSystem {
	config := newActorSystemConfiguration()
	for _, c := range configurator {
		c.Configure(config)
	}

	system := &ActorSystem{
		config:      config,
		bindAddress: bindAddress,
		metadata:    new(cm.Metadata),
		state:       newActorSystemState(),
	}

	config.WithShared(sharedAddress)
	config.WithShutdownAfterHooks(system.onShutdown)

	system.ActorSystem = vivid.NewActorSystemWithConfiguration(config.ActorSystemConfiguration)
	system.metadata.LaunchAt = time.Now().UnixMilli()
	system.metadata.Abilities = collection.ConvertMapValuesToBoolMap(config.abilities)

	system.start()

	return system
}

type ActorSystem struct {
	*vivid.ActorSystem // 如果单独使用，那么一切行为将越过集群
	config             *ActorSystemConfiguration
	metadata           *cm.Metadata
	state              *actorSystemState
	bindAddress        prc.PhysicalAddress
	memberlist         *memberlist.Memberlist
}

// ClusterRef 返回该集群的引用
func (sys *ActorSystem) ClusterRef() vivid.ActorRef {
	return sys.metadata.ProcessId
}

// findClusterNode 寻找符合条件的集群节点
func (sys *ActorSystem) findClusterNode(ability string) vivid.ActorRef {
	var mds []*cm.Metadata
	for _, node := range sys.memberlist.Members() {
		var md = new(cm.Metadata)

		// false: not match cluster generate condition
		if err := proto.Unmarshal(node.Meta, md); err != nil {
			sys.Logger().Error("ActorSystemCluster", log.String("metadata error", node.Name), log.String("address", node.Addr.String()), log.Err(err))
			continue
		} else if !md.Abilities[ability] {
			continue
		}

		mds = append(mds, md)
	}

	if len(mds) == 0 {
		return nil
	}

	collection.Shuffle(&mds)
	return mds[0].ProcessId
}

// ActorOfC 以特定身份获取集群中的对应能力 Actor 的引用
func (sys *ActorSystem) ActorOfC(identity, ability string) vivid.ActorRef {
	nodeRef := sys.findClusterNode(ability)
	if nodeRef == nil {
		return sys.ActorSystem.Abyss()
	}

	msg := &cm.ActorOf{
		Identity: identity,
		Ability:  ability,
	}
	actorRef, err := vivid.FutureAsk[vivid.ActorRef](sys, nodeRef, msg).Result()
	if err != nil {
		sys.Logger().Error("ActorSystemCluster", log.Err(err))
		return sys.ActorSystem.Abyss()
	}

	return actorRef
}

func (sys *ActorSystem) start() {
	bindAddr, bindPort, err := net.SplitHostPort(sys.bindAddress)
	if err != nil {
		panic(err)
	}

	sys.metadata.ProcessId = sys.ActorOfF(func() vivid.Actor {
		return newDrillmasterActor(sys)
	}, func(descriptor *vivid.ActorDescriptor) {
		descriptor.WithName("cluster")
	})

	memberlistConfig := memberlist.DefaultLocalConfig()
	if sys.config.name != charproc.None {
		memberlistConfig.Name = sys.config.name
	}
	memberlistConfig.BindAddr, memberlistConfig.BindPort = bindAddr, convert.StringToInt(bindPort)
	memberlistConfig.AdvertisePort = memberlistConfig.BindPort
	if sys.config.advertiseAddr != charproc.None {
		advertiseAddr, advertisePort, err := net.SplitHostPort(sys.config.advertiseAddr)
		if err != nil {
			panic(err)
		}
		memberlistConfig.AdvertiseAddr, memberlistConfig.AdvertisePort = advertiseAddr, convert.StringToInt(advertisePort)
	}
	memberlistConfig.Delegate = newActorSystemDelegate(sys)

	list, err := memberlist.Create(memberlistConfig)
	if err != nil {
		panic(err)
	}

	sys.memberlist = list
	if len(sys.config.seedNodes) == 0 {
		_, err = sys.memberlist.Join([]string{fmt.Sprintf("%s:%s", bindAddr, bindPort)})
	} else {
		_, err = sys.memberlist.Join(sys.config.seedNodes)
	}

	if err != nil {
		panic(err)
	}
}

func (sys *ActorSystem) onShutdown() {
	var err error
	if err = sys.memberlist.Leave(15 * time.Second); err != nil {
		panic(err)
	}
	if err = sys.memberlist.Shutdown(); err != nil {
		panic(err)
	}
}
