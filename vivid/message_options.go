package vivid

import "time"

type MessageOption func(*MessageOptions)

type MessageOptions struct {
	reply bool

	Sender       ActorRef
	ReplyTimeout time.Duration
	ContextHook  func(MessageContext)
}

func (o *MessageOptions) apply(options []MessageOption) *MessageOptions {
	for _, option := range options {
		option(o)
	}
	return o
}

// WithSender 设置消息发送者，发送者可以有利于对消息流向的追踪
func WithSender(sender ActorRef) MessageOption {
	return func(options *MessageOptions) {
		options.Sender = sender
	}
}

// WithReplyTimeout 设置消息回复超时时间，当消息发送后等待回复的时间超过此时间时将会返回 nil
func WithReplyTimeout(timeout time.Duration) MessageOption {
	return func(options *MessageOptions) {
		options.ReplyTimeout = timeout
	}
}

// WithContextHook 设置消息上下文钩子，用于在消息发送前获取到消息上下文进行特殊处理
func WithContextHook(hook func(MessageContext)) MessageOption {
	return func(options *MessageOptions) {
		options.ContextHook = hook
	}
}
