package nexus

import "time"

const (
	DefaultLowHandlerThreshold = 200 * time.Millisecond
)

func NewEventOptions() *EventOptions {
	return &EventOptions{
		LowHandlerThreshold: DefaultLowHandlerThreshold,
		LowHandlerTrace:     false,
	}
}

type EventOptions struct {
	LowHandlerThreshold        time.Duration                          // 慢消息处理器检查阈值
	LowHandlerThresholdHandler func(cost time.Duration)               // 慢消息处理器检查阈值处理器
	LowHandlerTrace            bool                                   // 是否开启慢消息处理器跟踪
	LowHandlerTraceHandler     func(cost time.Duration, stack []byte) // 慢消息处理器跟踪处理器
}

// WithLowHandlerThreshold 设置慢消息处理器检查阈值
func (o *EventOptions) WithLowHandlerThreshold(d time.Duration, handler func(cost time.Duration)) *EventOptions {
	o.LowHandlerThreshold = d
	o.LowHandlerThresholdHandler = handler
	return o
}

// WithLowHandlerTrace 设置是否开启慢消息处理器跟踪
func (o *EventOptions) WithLowHandlerTrace(enable bool, handler func(cost time.Duration, stack []byte)) *EventOptions {
	o.LowHandlerTrace = enable
	o.LowHandlerTraceHandler = handler
	return o
}

// Apply 应用选项
func (o *EventOptions) Apply(opts ...*EventOptions) *EventOptions {
	if o == nil {
		return o
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		o.LowHandlerThreshold = opt.LowHandlerThreshold
		o.LowHandlerTrace = opt.LowHandlerTrace
	}

	return o
}
