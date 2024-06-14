package vivid

import "time"

type MessageOption func(*MessageOptions)

type MessageOptions struct {
	reply bool

	Sender       ActorRef
	ReplyTimeout time.Duration
	ContextHook  func(MessageContext)
	Priority     int64
	Instantly    bool // 是否立即执行
}

func (o *MessageOptions) apply(options []MessageOption) *MessageOptions {
	for _, option := range options {
		option(o)
	}
	return o
}

// WithInstantly 设置消息是否立即执行，如果设置为立即执行，消息将会被立即执行，否则将会被放入邮箱等待执行
//   - 由于没有放入邮箱等待执行，该消息是在当前线程中执行的，如果存在循环调用，可能会导致死锁
//   - 该可选项在跨网络调用时可能不会产生效果
//   - 该可选项将提供给 Dispatcher 进行处理，根据不同的 Dispatcher 实现，该可选项可能会被忽略
func WithInstantly(instantly bool) MessageOption {
	return func(options *MessageOptions) {
		options.Instantly = instantly
	}
}

// WithPriority 设置消息优先级，优先级越高的消息将会被优先处理
//   - 当 priority 的数值越小时，优先级越高
//   - 当邮箱类型为非优先级邮箱 PriorityMailboxFactoryId 时，该可选项会被忽略
func WithPriority(priority int64) MessageOption {
	return func(options *MessageOptions) {
		options.Priority = priority
	}
}

// WithSender 设置消息发送者，发送者可以有利于对消息流向的追踪
func WithSender(sender ActorRef) MessageOption {
	return func(options *MessageOptions) {
		options.Sender = sender
	}
}

// WithReplyTimeout 设置消息回复超时时间，当消息发送后等待回复的时间超过此时间时将会返回消息的零值
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
