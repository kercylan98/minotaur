package server

type Cross interface {
	// Init 初始化跨服
	//  - serverId: 本服id
	//  - packetHandle.serverId: 发送跨服消息的服务器id
	//  - packetHandle.packet: 数据包
	Init(serverId int64, packetHandle func(serverId int64, packet []byte))
	// PushMessage 推送跨服消息
	//  - serverId: 目标服务器id
	PushMessage(serverId int64, packet []byte) error
}
