package component

// LockstepClient 帧同步客户端接口定义
//   - 客户端应该具备ID及写入数据包的实现
type LockstepClient[ID comparable] interface {
	// GetID 用户玩家ID
	GetID() ID
	// Send 发送数据包
	//   - messageType: websocket模式中指定消息类型
	Send(packet []byte, messageType ...int)
}
