package cluster

import (
	"github.com/kercylan98/minotaur/engine/vivid"
)

type component struct {
	cluster *ActorSystem
}

func (c *component) OnInitialize(actorSystem *vivid.ActorSystem) error {
	return nil
}

func (c *component) OnShutdownBefore(actorSystem *vivid.ActorSystem) error {
	c.cluster.onShutdown()
	return nil
}
