package minotaur

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

type applicationActor struct {
	*Application
}

func (a *applicationActor) OnReceive(ctx vivid.MessageContext) {
	a.onReceive(ctx)
}
