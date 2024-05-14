package vivid

type MessageOption func(opts *MessageOptions)

type MessageOptions struct {
	Sender ActorRef `json:"-"` // 发送者
}

// WithMessageOptions 设置消息选项
func WithMessageOptions(options *MessageOptions) MessageOption {
	return func(opts *MessageOptions) {
		opts.Sender = options.Sender
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
		opts.Sender = sender
	}
}
