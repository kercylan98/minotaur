package cluster

import (
	"github.com/kercylan98/minotaur/engine/vivid/cluster/internal/cm"
	"sync"
)

func newActorSystemState() *actorSystemState {
	return &actorSystemState{
		data: new(cm.State),
	}
}

type actorSystemState struct {
	rw   sync.RWMutex
	data *cm.State
}
