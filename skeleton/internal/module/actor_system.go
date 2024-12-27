package module

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
)

type ActorSystemModule interface {
	application.Module

	ActorSystem() *vivid.ActorSystem

	Cluster() *cluster.ActorSystem
}
