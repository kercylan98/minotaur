package vivid

import (
	"github.com/kercylan98/minotaur/engine/vivid/supervision"
	"github.com/kercylan98/minotaur/toolkit/log"
	"time"
)

type guard struct {
	supervision.Strategy
}

func (g *guard) OnReceive(ctx ActorContext) {
	switch ctx.Message().(type) {
	case *OnLaunch:
		g.Strategy = supervision.OneForOne(10, time.Millisecond*200, time.Second*3, supervision.FunctionalDecide(func(record *supervision.AccidentRecord) supervision.Directive {
			ctx.System().Logger().Warn("ActorSystem", log.String("system", ctx.System().PhysicalAddress()), log.String("info", "unsupervised accidents"), log.String("actor", record.Victim.GetLogicalAddress()), log.Any("reason", record.Reason))
			if len(record.Stack) > 0 {
				ctx.System().Logger().Error("trace", log.StackData("stack", record.Stack))
			}
			return supervision.DirectiveRestart
		}))
	}
}
