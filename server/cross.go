package server

// Cross 跨服功能接口实现
type Cross interface {
	// PushPacket 推送数据包
	PushPacket(serverId int64, packet []byte) error

	//
}
