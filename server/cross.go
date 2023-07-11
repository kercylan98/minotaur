package server

// Cross 跨服接口
type Cross interface {
	// Init 初始化跨服
	//  - serverId: 本服id
	//  - packetHandle.serverId: 发送跨服消息的服务器id
	//  - packetHandle.packet: 数据包
	Init(server *Server, packetHandle func(serverId int64, packet []byte)) error
	// PushMessage 推送跨服消息
	//  - serverId: 目标服务器id
	PushMessage(serverId int64, packet []byte) error
	// Release 释放资源
	Release()
}
