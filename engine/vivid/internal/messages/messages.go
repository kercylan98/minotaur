package messages

import messagesv1 "github.com/kercylan98/minotaur/engine/vivid/internal/messages/v1"

type (
	Terminated              = messagesv1.Terminated
	Watch                   = messagesv1.Watch
	Unwatch                 = messagesv1.Unwatch
	SlowProcess             = messagesv1.SlowProcess
	SubscribeRequest        = messagesv1.SubscribeRequest
	UnsubscribeRequest      = messagesv1.UnsubscribeRequest
	Subscription            = messagesv1.Subscription
	PublishRequestBroadcast = messagesv1.PublishRequestBroadcast
	AbyssMessageEvent       = messagesv1.AbyssMessageEvent
)

type LocalPublishRequest struct {
	Topic   string
	Message any
}
