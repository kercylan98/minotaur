package pulse

import "time"

type SubscribeOption func(*SubscribeOptions)

type SubscribeOptions struct {
	Priority        *int64         // 订阅优先级，优先级越高，订阅者越早收到事件，数值越小优先级越高
	PriorityTimeout *time.Duration // 优先级事件超时时间
	Producer        *Producer      // 事件生产者
}

func (o *SubscribeOptions) apply(opts []SubscribeOption) *SubscribeOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// WithSubscribePriority 设置订阅优先级
func WithSubscribePriority(priority int64, timeout time.Duration) SubscribeOption {
	return func(options *SubscribeOptions) {
		options.Priority = &priority
		options.PriorityTimeout = &timeout
	}
}

// WithSubscribeProducer 设置事件生产者
func WithSubscribeProducer(producer Producer) SubscribeOption {
	return func(options *SubscribeOptions) {
		options.Producer = &producer
	}
}
