package logger

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
)

type Actor struct {
	*log.Logger
}

func (a *Actor) OnReceive(ctx vivid.MessageContext) {

}
