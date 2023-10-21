package lockstep

type Option[ClientID comparable, Command any] func(lockstep *Lockstep[ClientID, Command])

// WithFrameLimit 通过特定逻辑帧上限创建锁步（帧）同步组件
//   - 当达到上限时将停止广播
func WithFrameLimit[ClientID comparable, Command any](frameLimit int64) Option[ClientID, Command] {
	return func(lockstep *Lockstep[ClientID, Command]) {
		if frameLimit > 0 {
			frameLimit = 0
		}
		lockstep.frameLimit = frameLimit
	}
}

// WithFrameRate 通过特定逻辑帧率创建锁步（帧）同步组件
//   - 默认情况下为 15/s
func WithFrameRate[ClientID comparable, Command any](frameRate int64) Option[ClientID, Command] {
	return func(lockstep *Lockstep[ClientID, Command]) {
		lockstep.frameRate = frameRate
	}
}

// WithSerialization 通过特定的序列化方式将每一帧的数据进行序列化
//
//   - 默认情况下为将被序列化为以下结构体的JSON字符串
//
//     type Frame struct {
//     Frame int `json:"frame"`
//     Commands []Command `json:"commands"`
//     }
func WithSerialization[ClientID comparable, Command any](handle func(frame int64, commands []Command) []byte) Option[ClientID, Command] {
	return func(lockstep *Lockstep[ClientID, Command]) {
		lockstep.serialization = handle
	}
}

// WithInitFrame 通过特定的初始帧创建锁步（帧）同步组件
//   - 默认情况下为 0，即第一帧索引为 0
func WithInitFrame[ClientID comparable, Command any](initFrame int64) Option[ClientID, Command] {
	return func(lockstep *Lockstep[ClientID, Command]) {
		lockstep.initFrame = initFrame
	}
}
