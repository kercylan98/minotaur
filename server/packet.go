package server

// NewPacket 创建一个数据包
func NewPacket(data []byte) Packet {
	return Packet{
		Data: data,
	}
}

// NewWSPacket 创建一个 websocket 数据包
func NewWSPacket(websocketType int, data []byte) Packet {
	return Packet{
		WebsocketType: websocketType,
		Data:          data,
	}
}

// NewPacketString 创建一个字符串数据包
func NewPacketString(data string) Packet {
	return Packet{
		Data: []byte(data),
	}
}

// NewWSPacketString 创建一个 websocket 字符串数据包
func NewWSPacketString(websocketType int, data string) Packet {
	return Packet{
		WebsocketType: websocketType,
		Data:          []byte(data),
	}
}

// Packet 数据包
type Packet struct {
	WebsocketType int    // websocket 消息类型
	Data          []byte // 数据
}

// String 转换为字符串
func (slf Packet) String() string {
	return string(slf.Data)
}
