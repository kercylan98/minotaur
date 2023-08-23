package client

type Core interface {
	// Run 启动客户端
	//  - runState: 运行状态，当客户端启动完成时，应该向该通道发送 error 或 nil
	//  - receive: 接收到数据包时应该将数据包发送到该函数，wst 表示 websocket 的数据类型，data 表示数据包
	Run(runState chan<- error, receive func(wst int, packet []byte))

	// Write 向客户端写入数据包
	Write(packet *Packet) error

	// Close 关闭客户端
	Close()

	// GetServerAddr 获取服务器地址
	GetServerAddr() string
}
