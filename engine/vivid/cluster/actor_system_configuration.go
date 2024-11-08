package cluster

import (
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
	shutdownTimeout time.Duration // 关闭超时时间
}

// WithShutdownTimeout 设置关闭超时时间
func (c *ActorSystemConfiguration) WithShutdownTimeout(timeout time.Duration) *ActorSystemConfiguration {
	c.shutdownTimeout = timeout
	return c
}
