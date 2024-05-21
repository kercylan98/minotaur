package vivid

import "time"

type MessageOption func(*MessageOptions)

type MessageOptions struct {
	reply    bool
	replySeq uint64

	Sender       ActorRef
	ReplyTimeout time.Duration
}

func (o *MessageOptions) apply(options []MessageOption) *MessageOptions {
	for _, option := range options {
		option(o)
	}
	return o
}

func WithSender(sender ActorRef) MessageOption {
	return func(options *MessageOptions) {
		options.Sender = sender
	}
}

func WithReplyTimeout(timeout time.Duration) MessageOption {
	return func(options *MessageOptions) {
		options.ReplyTimeout = timeout
	}
}
