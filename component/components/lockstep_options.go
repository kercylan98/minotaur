package components

type LockstepOption[ClientID comparable, Command any] func(lockstep *Lockstep[ClientID, Command])

// WithLockstepFrameLimit 通过特定逻辑帧上限创建锁步（帧）同步组件
//   - 当达到上限时将停止广播
func WithLockstepFrameLimit[ClientID comparable, Command any](frameLimit int) LockstepOption[ClientID, Command] {
	return func(lockstep *Lockstep[ClientID, Command]) {
		if frameLimit > 0 {
			frameLimit = 0
		}
		lockstep.frameLimit = frameLimit
	}
}

// WithLockstepFrameRate 通过特定逻辑帧率创建锁步（帧）同步组件
//   - 默认情况下为 15/s
func WithLockstepFrameRate[ClientID comparable, Command any](frameRate int) LockstepOption[ClientID, Command] {
	return func(lockstep *Lockstep[ClientID, Command]) {
		lockstep.frameRate = frameRate
	}
}

// WithLockstepSerialization 通过特定的序列化方式将每一帧的数据进行序列化
//
//   - 默认情况下为将被序列化为以下结构体的JSON字符串
//
//     type Frame struct {
//     Frame int `json:"frame"`
//     Commands []Command `json:"commands"`
//     }
func WithLockstepSerialization[ClientID comparable, Command any](handle func(frame int, commands []Command) []byte) LockstepOption[ClientID, Command] {
	return func(lockstep *Lockstep[ClientID, Command]) {
		lockstep.serialization = handle
	}
}
