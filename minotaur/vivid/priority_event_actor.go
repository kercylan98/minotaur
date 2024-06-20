package vivid

import "time"

type priorityEventMessage struct {
	event Event
}

type priorityEventActor struct {
	subscribes []subscribeInfo
}

func (p *priorityEventActor) OnReceive(ctx MessageContext) {
	switch m := ctx.GetMessage().(type) {
	case priorityEventMessage:
		for _, info := range p.subscribes {
			var timeout time.Duration
			if info.priorityTimeout != nil {
				timeout = *info.priorityTimeout
			}
			info.subscriber.Ask(m.event, WithReplyTimeout(timeout))
		}
	}
}
