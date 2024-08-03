package cluster

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/prc/codec"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/random"
)

func newActorSystemConfiguration() *ActorSystemConfiguration {
	return &ActorSystemConfiguration{
		name:                     random.HostName(),
		ActorSystemConfiguration: vivid.NewActorSystemConfiguration(),
		abilities:                make(map[string]*ability),
	}
}

type ActorSystemConfiguration struct {
	*vivid.ActorSystemConfiguration

	name          string              // 集群内唯一的节点名称
	seedNodes     []string            // 默认散播的种子节点地址
	abilities     map[string]*ability // 集群提供的能力
	advertiseAddr prc.PhysicalAddress // 广告地址
}

// WithAdvertiseAddress 设置广告地址
func (c *ActorSystemConfiguration) WithAdvertiseAddress(address prc.PhysicalAddress) *ActorSystemConfiguration {
	c.advertiseAddr = address
	return c
}

// WithShared 设置是否开启网络共享，开启后 ActorSystem 将允许通过网络与其他 ActorSystem 交互。
//   - 默认的网络序列化是采用的 ProtoBuffer，如果需要调整，可指定编解码器
//
// 在集群模式下，address 的值将被 advertiseAddress 所取代，但 codec 仍然有效
func (c *ActorSystemConfiguration) WithShared(address prc.PhysicalAddress, codec ...codec.Codec) *ActorSystemConfiguration {
	c.ActorSystemConfiguration.WithShared(address, codec...)
	return c
}

// WithSeedNodes 设置默认散播的种子节点地址
func (c *ActorSystemConfiguration) WithSeedNodes(nodes ...string) *ActorSystemConfiguration {
	c.seedNodes = nodes
	return c
}

// WithName 设置集群内节点名称
func (c *ActorSystemConfiguration) WithName(name string) *ActorSystemConfiguration {
	c.name = name
	return c
}

// WithAbility 声明集群能提供的能力
func (c *ActorSystemConfiguration) WithAbility(name string, provider ActorProvider, configurator ...vivid.ActorDescriptorConfigurator) *ActorSystemConfiguration {
	_, exist := c.abilities[name]
	if exist {
		panic(fmt.Errorf("ability %s already exists", name))
	}

	c.abilities[name] = newAbility(provider, configurator...)
	return c
}
