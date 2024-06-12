package vivid

import (
	"time"
)

type priorityEventMessage struct {
	event Event
}

type priorityEventActor struct {
	subscribes []subscribeInfo
}

func (p *priorityEventActor) OnReceive(ctx MessageContext) {
	for _, info := range p.subscribes {
		var timeout time.Duration
		if info.priorityTimeout != nil {
			timeout = *info.priorityTimeout
		}
		info.subscriber.Ask(ctx.GetMessage(), WithReplyTimeout(timeout))
	}
}
