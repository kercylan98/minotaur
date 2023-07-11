package server

// Packet 数据包
type Packet struct {
	WebsocketType int    // websocket 消息类型
	Data          []byte // 数据
}

func (slf Packet) String() string {
	return string(slf.Data)
}
