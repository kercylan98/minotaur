package vivid

func NewPriorityOptions() *PriorityOptions {
	return &PriorityOptions{
		BufferSize: 1024,
		StopMode:   PriorityStopModeGraceful,
	}
}

type PriorityOption func(*PriorityOptions)

type PriorityOptions struct {
	options    []PriorityOption
	BufferSize uint             // 消息队列的缓冲区大小
	StopMode   PriorityStopMode // 消息队列的停止模式
}

// Apply 用于应用 Priority 配置
func (o *PriorityOptions) Apply(opts ...*PriorityOptions) *PriorityOptions {
	for _, opt := range opts {
		for _, option := range opt.options {
			option(o)
		}
	}
	return o
}

// WithBufferSize 用于设置消息队列的缓冲区大小
func (o *PriorityOptions) WithBufferSize(size uint) *PriorityOptions {
	o.options = append(o.options, func(o *PriorityOptions) {
		o.BufferSize = size
	})
	return o
}

// WithStopMode 用于设置消息队列的停止模式
func (o *PriorityOptions) WithStopMode(mode PriorityStopMode) *PriorityOptions {
	o.options = append(o.options, func(o *PriorityOptions) {
		o.StopMode = mode
	})
	return o
}
