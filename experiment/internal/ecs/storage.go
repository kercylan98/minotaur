package ecs

import (
	"github.com/kercylan98/minotaur/experiment/internal/ecs/storage"
)

type Storage = storage.Storage[entityId, ComponentId]
