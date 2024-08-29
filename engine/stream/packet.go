package stream

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
func NewPacket() *Packet {
	return &Packet{}
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
func NewPacketS(data string) *Packet {
	return NewPacketD([]byte(data))
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
func NewPacketD(data []byte) *Packet {
	return &Packet{data: data}
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
func NewPacketC(ctx any) *Packet {
	return &Packet{ctx: ctx}
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
func NewPacketDC(data []byte, ctx any) *Packet {
	return &Packet{
		data: data,
		ctx:  ctx,
	}
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
func NewPacketSC(data string, ctx any) *Packet {
	return NewPacketDC([]byte(data), ctx)
}

// Deprecated: 该设计加大了理解成本，且不易于使用，将考虑新的方案用于处理网络连接。至 v0.7.0 版本及以后，stream 包将被移除。
type Packet struct {
	data []byte // 数据包数据
	ctx  any    // 数据包上下文
}

// Derivation 继承数据包上下文派生一个新的数据包
func (p *Packet) Derivation(data []byte) *Packet {
	return NewPacketDC(data, p.ctx)
}

func (p *Packet) SetData(data []byte) *Packet {
	p.data = data
	return p
}

func (p *Packet) SetString(data string) *Packet {
	p.data = []byte(data)
	return p
}

func (p *Packet) SetContext(ctx any) *Packet {
	p.ctx = ctx
	return p
}

func (p *Packet) Data() []byte {
	return p.data
}

func (p *Packet) Context() any {
	return p.ctx
}
