package unsafevivid

import "github.com/kercylan98/minotaur/vivid/vivids"

type SubscribeEventMessage struct {
	Subscriber vivids.ActorId
	Event      vivids.Event
}
