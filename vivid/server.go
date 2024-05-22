package vivid

type Server interface {
	// C 用于获取一个用于接收远程消息的通道
	C() <-chan []byte
}
