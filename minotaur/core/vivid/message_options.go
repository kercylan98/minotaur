package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/pools"
	"time"
)

type MessageOption func(options *MessageOptions)

type (
	RegulatoryMessageHook func(*RegulatoryMessage)
	MessageHook           func(message Message, cover func(Message))
)

type MessageOptions struct {
	options       []MessageOption
	FutureTimeout time.Duration
	AskReplyAgent ActorRef

	RegulatoryMessageHooks []RegulatoryMessageHook
	MessageHooks           []MessageHook
}

// WithMessageHook 配置处理 Message 的钩子函数，这个钩子函数将在发送 Message 前调用
func (o *MessageOptions) WithMessageHook(hook MessageHook) *MessageOptions {
	o.options = append(o.options, func(options *MessageOptions) {
		options.MessageHooks = append(options.MessageHooks, hook)
	})
	return o
}

// WithRegulatoryMessageHook 配置处理 RegulatoryMessage 的钩子函数，这个钩子函数将在 ActorContext.Ask 方法发送 RegulatoryMessage 前调用
func (o *MessageOptions) WithRegulatoryMessageHook(hook RegulatoryMessageHook) *MessageOptions {
	o.options = append(o.options, func(options *MessageOptions) {
		options.RegulatoryMessageHooks = append(options.RegulatoryMessageHooks, hook)
	})
	return o
}

// WithReplyAgent 配置通过 ActorContext.Ask 方法响应消息的代理 Actor，这个代理 Actor 将接替原 Actor 处理响应
func (o *MessageOptions) WithReplyAgent(agent ActorRef) *MessageOptions {
	o.options = append(o.options, func(options *MessageOptions) {
		options.WithRegulatoryMessageHook(func(message *RegulatoryMessage) {
			message.Sender = agent
		})
	})
	return o
}

// WithFutureTimeout 配置通过 ActorContext.FutureAsk 方法发送消息的超时时间
//   - 默认情况下的超时时间为 1 秒，假如超时时间小于等于 0，这个消息将无限期等待
func (o *MessageOptions) WithFutureTimeout(timeout time.Duration) {
	o.options = append(o.options, func(options *MessageOptions) {
		options.FutureTimeout = timeout
	})
}

func (o *MessageOptions) apply() *MessageOptions {
	for _, option := range o.options {
		option(o)
	}
	return o
}

var messageOptionsPool = pools.NewObjectPool[MessageOptions](func() *MessageOptions {
	return &MessageOptions{
		FutureTimeout: time.Second,
	}
}, func(data *MessageOptions) {
	data.options = data.options[:0]
	data.FutureTimeout = time.Second
	data.AskReplyAgent = nil
	data.MessageHooks = data.MessageHooks[:0]
})

func generateMessageOptions(options ...MessageOption) *MessageOptions {
	var opts = messageOptionsPool.Get()
	for _, opt := range options {
		opt(opts)
	}
	return opts.apply()
}

func releaseMessageOptions(opts *MessageOptions) {
	messageOptionsPool.Put(opts)
}

func (o *MessageOptions) hookRegulatoryMessage(m *RegulatoryMessage) {
	for _, hook := range o.RegulatoryMessageHooks {
		hook(m)
	}
}

func (o *MessageOptions) hookMessage(m Message, cover func(message Message)) {
	for _, hook := range o.MessageHooks {
		hook(m, cover)
	}
}
