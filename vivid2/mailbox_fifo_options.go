package vivid

func NewFIFOOptions() *FIFOOptions {
	return &FIFOOptions{
		BufferSize: 1024,
		StopMode:   FIFOStopModeGraceful,
	}
}

type FIFOOption func(*FIFOOptions)

type FIFOOptions struct {
	options    []FIFOOption
	BufferSize uint         // 消息队列的缓冲区大小
	StopMode   FifoStopMode // 消息队列的停止模式
}

// Apply 用于应用 FIFO 配置
func (o *FIFOOptions) Apply(opts ...*FIFOOptions) *FIFOOptions {
	for _, opt := range opts {
		for _, option := range opt.options {
			option(o)
		}
	}
	return o
}

// WithBufferSize 用于设置消息队列的缓冲区大小
func (o *FIFOOptions) WithBufferSize(size uint) *FIFOOptions {
	o.options = append(o.options, func(o *FIFOOptions) {
		o.BufferSize = size
	})
	return o
}

// WithStopMode 用于设置消息队列的停止模式
func (o *FIFOOptions) WithStopMode(mode FifoStopMode) *FIFOOptions {
	o.options = append(o.options, func(o *FIFOOptions) {
		o.StopMode = mode
	})
	return o
}
