package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"sync/atomic"
)

type gnetConnContext struct {
	ref    vivid.ActorRef
	status *atomic.Uint32
	actor  *gnetConnActor
}
