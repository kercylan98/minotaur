package lockstep

// Client 帧同步客户端接口定义
//   - 客户端应该具备ID及写入数据包的实现
type Client[ID comparable] interface {
	// GetID 用户玩家ID
	GetID() ID
	// Write 写入数据包
	Write(packet []byte, callback ...func(err error))
}
