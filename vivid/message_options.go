package vivid

type MessageOption func(opts *MessageOptions)

type MessageOptions struct {
	Reply bool // 是否需要回复

	sender ActorRef // 发送者
}

// WithMessageOptions 设置消息选项
func WithMessageOptions(options *MessageOptions) MessageOption {
	return func(opts *MessageOptions) {
		opts.sender = options.sender
		opts.Reply = options.Reply
	}
}

func (o *MessageOptions) apply(options ...MessageOption) *MessageOptions {
	for _, option := range options {
		option(o)
	}
	return o
}

// WithMessageSender 设置发送者
func WithMessageSender(sender ActorRef) MessageOption {
	return func(opts *MessageOptions) {
		opts.sender = sender
	}
}

// WithMessageReply 设置是否需要回复
func WithMessageReply(reply bool) MessageOption {
	return func(opts *MessageOptions) {
		opts.Reply = reply
	}
}
