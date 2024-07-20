package ecs

import (
	"github.com/kercylan98/minotaur/engine/ecs/storage"
)

type Storage = storage.Storage[entityId, ComponentId]
