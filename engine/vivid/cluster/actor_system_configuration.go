package cluster

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"time"
)

func newActorSystemConfiguration() *ActorSystemConfiguration {
	asc := &ActorSystemConfiguration{
		ActorSystemConfiguration: vivid.NewActorSystemConfiguration(),
		shutdownTimeout:          time.Minute,
		//nodeBalancer: FunctionalBalancer(func(nodes []*Node) *Node {
		//	if len(nodes) == 0 {
		//		return nil
		//	}
		//	return collection.ChooseRandomSliceElement(nodes)
		//}),
	}
	return asc
}

type ActorSystemConfiguration struct {
	*vivid.ActorSystemConfiguration
	seeds              []prc.PhysicalAddress        // 种子节点
	seedProvider       SeedProvider                 // 种子节点提供者
	shutdownTimeout    time.Duration                // 关闭超时时间
	onlyActorProviders map[string]OnlyActorProvider // 集群内唯一的 Actor
	//nodeBalancer       Balancer                     // 节点负载均衡器
}

//
//// WithNodeBalancer 设置节点负载均衡器
//func (c *ActorSystemConfiguration) WithNodeBalancer(balancer Balancer) *ActorSystemConfiguration {
//	if balancer == nil {
//		panic(fmt.Errorf("the node balancer cannot be nil"))
//	}
//	c.nodeBalancer = balancer
//	return c
//}

// WithSeeds 设置种子节点
func (c *ActorSystemConfiguration) WithSeeds(seeds ...prc.PhysicalAddress) *ActorSystemConfiguration {
	c.seeds = append(c.seeds, seeds...)
	return c
}

// WithSeedProvider 设置种子节点提供者，当集群启动时将会额外通过种子节点提供者获取种子节点
func (c *ActorSystemConfiguration) WithSeedProvider(provider SeedProvider) *ActorSystemConfiguration {
	c.seedProvider = provider
	return c
}

// WithOnlyActorProvider 设置集群内唯一的 Actor
func (c *ActorSystemConfiguration) WithOnlyActorProvider(name string, provider OnlyActorProvider) *ActorSystemConfiguration {
	if c.onlyActorProviders == nil {
		c.onlyActorProviders = make(map[string]OnlyActorProvider)
	}
	if _, exist := c.onlyActorProviders[name]; exist {
		panic(fmt.Errorf("the only actor provider[%v] has already been registered", name))
	}
	c.onlyActorProviders[name] = provider
	return c
}

// WithShutdownTimeout 设置关闭超时时间
func (c *ActorSystemConfiguration) WithShutdownTimeout(timeout time.Duration) *ActorSystemConfiguration {
	c.shutdownTimeout = timeout
	return c
}
