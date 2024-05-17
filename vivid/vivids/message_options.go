package vivids

import "time"

type MessageOption func(opts *MessageOptions)

type MessageOptions struct {
	ReplyTimeout time.Duration // 回复超时时间
	SenderId     ActorId       // 显式发送者
}

// WithMessageOptions 设置消息选项
func WithMessageOptions(options *MessageOptions) MessageOption {
	return func(opts *MessageOptions) {
		opts.SenderId = options.SenderId
		opts.ReplyTimeout = options.ReplyTimeout
	}
}

func (o *MessageOptions) Apply(options ...MessageOption) *MessageOptions {
	for _, option := range options {
		option(o)
	}
	return o
}

// WithMessageSender 设置发送者
func WithMessageSender(sender ActorContext) MessageOption {
	return func(opts *MessageOptions) {
		opts.SenderId = sender.GetActorId()
	}
}

// WithMessageReply 设置是否需要回复，当超时时间 <= 0 时，表示不需要回复
func WithMessageReply(timeout time.Duration) MessageOption {
	return func(opts *MessageOptions) {
		opts.ReplyTimeout = timeout
	}
}
