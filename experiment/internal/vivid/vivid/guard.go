package vivid

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid/supervision"
	"time"
)

type guard struct {
	supervision.Strategy
}

func (g *guard) OnReceive(ctx ActorContext) {
	switch ctx.Message().(type) {
	case *OnLaunch:
		g.Strategy = supervision.OneForOne(10, time.Millisecond*200, time.Second*3, supervision.FunctionalDecide(func(record *supervision.AccidentRecord) supervision.Directive {
			return supervision.DirectiveRestart
		}))
	}
}
