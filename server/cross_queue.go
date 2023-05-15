package server

type CrossQueueName string

// CrossQueue 跨服消息队列接口
type CrossQueue interface {
	// GetName 获取跨服消息队列名称
	GetName() CrossQueueName
	// Init 初始化队列
	Init() error
	// Publish 发布跨服消息
	Publish(serverId int64, packet []byte) error
	// Subscribe 接收到跨服消息
	Subscribe(serverId int64, packetHandle func([]byte))
}
