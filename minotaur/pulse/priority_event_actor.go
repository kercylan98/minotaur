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
	for _, info := range p.subscribes {
		var timeout time.Duration
		if info.priorityTimeout != nil {
			timeout = *info.priorityTimeout
		}
		info.subscriber.Ask(ctx.GetMessage(), vivid.WithReplyTimeout(timeout))
	}
}
