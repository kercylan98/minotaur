package vivid

// Server 是 ActorSystem 用于接收远程调用的服务端
//   - 该服务端仅需提供一个可用于接收远程消息的通道即可
type Server interface {
	// C 用于获取一个用于接收远程消息的通道
	C() <-chan []byte
}
