package module

import "github.com/kercylan98/minotaur/engine/vivid"

type Application interface {
	ActorSystem() *vivid.ActorSystem
}
