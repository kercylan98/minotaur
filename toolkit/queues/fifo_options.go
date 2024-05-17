package queues

func NewFIFOOptions() *FIFOOptions {
	return &FIFOOptions{
		BufferSize:        1024,
		StopMode:          FIFOStopModeGraceful,
		DequeueBufferSize: 0,
		DequeueFullBlock:  true,
		PickUpDiscarded:   false,
	}
}

type FIFOOptions struct {
	BufferSize        uint         // 消息队列的缓冲区大小
	StopMode          FifoStopMode // 消息队列的停止模式
	DequeueBufferSize uint         // 出列通道的缓冲区大小
	DequeueFullBlock  bool         // 出列通道满时是否阻塞，默认为阻塞，如果设置为 false，则当出列通道满时会丢弃新的消息
	PickUpDiscarded   bool         // 是否将被丢弃的消息重新放回队列中
}

// Apply 用于应用 FIFO 配置
func (o *FIFOOptions) Apply(opts ...*FIFOOptions) *FIFOOptions {
	for _, opt := range opts {
		o.BufferSize = opt.BufferSize
		o.StopMode = opt.StopMode
		o.DequeueBufferSize = opt.DequeueBufferSize
		o.DequeueFullBlock = opt.DequeueFullBlock
	}
	return o
}

// WithBufferSize 用于设置消息队列的缓冲区大小
func (o *FIFOOptions) WithBufferSize(size uint) *FIFOOptions {
	o.BufferSize = size
	return o
}

// WithStopMode 用于设置消息队列的停止模式
func (o *FIFOOptions) WithStopMode(mode FifoStopMode) *FIFOOptions {
	o.StopMode = mode
	return o
}

// WithDequeueBufferSize 用于设置出列通道的缓冲区大小
func (o *FIFOOptions) WithDequeueBufferSize(size uint) *FIFOOptions {
	o.DequeueBufferSize = size
	return o
}

// WithDequeueFullBlock 用于设置出列通道满时是否阻塞
//   - 当设置为 true 时，出列通道满时会阻塞，否则会丢弃新的消息
func (o *FIFOOptions) WithDequeueFullBlock(block bool) *FIFOOptions {
	o.DequeueFullBlock = block
	return o
}

// WithPickUpDiscarded 用于设置是否将被丢弃的消息重新放回队列的尾部
//   - 当消息被丢弃时，如果设置为 true，则会将被丢弃的消息重新放回队列中
//   - 由于队列的缓冲区是无限的，因此可能会导致队列的消息堆积
func (o *FIFOOptions) WithPickUpDiscarded(pickUp bool) *FIFOOptions {
	o.PickUpDiscarded = pickUp
	return o
}
