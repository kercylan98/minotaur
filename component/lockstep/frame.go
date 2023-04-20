package lockstep

type Frame[Command any] interface {
	// GetIndex 获取这一帧的索引
	GetIndex() uint32
	// GetCommands 获取这一帧的数据
	GetCommands() []Command
	// Marshal 序列化帧数据
	Marshal() ([]byte, error)
}
