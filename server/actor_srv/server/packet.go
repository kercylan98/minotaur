package server

// NewPacket 创建一个新的数据包
func NewPacket(data []byte) Packet {
	return new(packet).init(data)
}

// Packet 写入连接的数据包接口
type Packet interface {
	// GetBytes 获取数据包字节流
	GetBytes() []byte
	// SetContext 设置数据包上下文，上下文通常是受特定 Network 实现所限制的
	//  - 在内置的 network.WebSocket 实现中，上下文被用于指定连接发送数据的操作码
	SetContext(ctx any) Packet
	// GetContext 获取数据包上下文
	GetContext() any
}

type packet struct {
	ctx  any
	data []byte
}

func (m *packet) init(data []byte) Packet {
	m.data = data
	return m
}

func (m *packet) reset() {
	m.ctx = nil
	m.data = m.data[:0]
}

func (m *packet) GetBytes() []byte {
	return m.data
}

func (m *packet) SetContext(ctx any) Packet {
	m.ctx = ctx
	return m
}

func (m *packet) GetContext() any {
	return m.ctx
}
