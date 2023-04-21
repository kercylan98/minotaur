package lockstep

// Writer 游戏帧写入器，通常实现写入器的对象应该为包含网络连接的玩家
type Writer[ID comparable, FrameCommand any] interface {
	// GetID 游戏帧写入器ID
	GetID() ID
	// Healthy 检查写入器状态是否健康，例如离线、网络环境异常等
	Healthy() bool
	// Marshal 将多帧数据转换为流格式，以对游戏帧写入器进行写入
	Marshal(frames map[uint32]Frame[FrameCommand]) []byte
	// Write 向游戏帧写入器中写入数据
	Write(data []byte) error
}
