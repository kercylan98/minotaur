package cluster

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	"time"
)

func newActorSystemConfiguration() *ActorSystemConfiguration {
	asc := &ActorSystemConfiguration{
		ActorSystemConfiguration: vivid.NewActorSystemConfiguration(),
		shutdownTimeout:          time.Minute,
	}
	return asc
}

type ActorSystemConfiguration struct {
	*vivid.ActorSystemConfiguration
	shutdownTimeout    time.Duration                // 关闭超时时间
	onlyActorProviders map[string]OnlyActorProvider // 集群内唯一的 Actor
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
