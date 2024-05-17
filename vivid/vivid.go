package vivid

import (
	"github.com/kercylan98/minotaur/vivid/unsafevivid"
	"github.com/kercylan98/minotaur/vivid/vivids"
)

func NewActorSystem(name string, opts ...*vivids.ActorSystemOptions) vivids.ActorSystem {
	return unsafevivid.NewActorSystem(name, opts...)
}
