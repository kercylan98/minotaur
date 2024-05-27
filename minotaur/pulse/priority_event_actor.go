package pulse

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"time"
)

type priorityEventMessage struct {
	event Event
}

type priorityEventActor struct {
	subscribes []subscribeInfo
}

func (p *priorityEventActor) OnReceive(ctx vivid.MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case Event:
		for _, info := range p.subscribes {
			var timeout time.Duration
			if info.priorityTimeout != nil {
				timeout = *info.priorityTimeout
			}
			info.subscriber.Ask(m, vivid.WithReplyTimeout(timeout))
		}
	}
}
