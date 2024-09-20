package socket

// NewPacket 创建一个用于网络的数据包，它允许包含一个上下文对象。
// 上下文对象将被用于数据的写入，根据不同的 Writer 定义，它可能会拥有不同的要求标准
func NewPacket(data []byte, ctx any) *Packet {
	return &Packet{data: data, ctx: ctx}
}

// Packet 网络数据包
type Packet struct {
	data []byte
	ctx  any
}

// GetData 获取数据包数据
func (p *Packet) GetData() []byte {
	return p.data
}

// GetContext 获取数据包上下文
func (p *Packet) GetContext() any {
	return p.ctx
}
