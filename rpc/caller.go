package rpc

type Caller interface {
	// GetRoute 获取调用路由
	GetRoute() []Route

	// GetPacket 获取数据包
	GetPacket() []byte

	// Respond 响应调用
	Respond(packet []byte) error
}
