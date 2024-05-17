package vivids

type Receiver interface {
	Tell(v Message, opts ...MessageOption) error
	Ask(v Message, opts ...MessageOption) (any, error)
}

type Future interface {
	// IsCompleted 判断 Future 是否已完成
	IsCompleted() bool
}
