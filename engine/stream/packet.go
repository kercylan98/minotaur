package stream

func NewPacket() *Packet {
	return &Packet{}
}

func NewPacketS(data string) *Packet {
	return NewPacketD([]byte(data))
}

func NewPacketD(data []byte) *Packet {
	return &Packet{data: data}
}

func NewPacketC(ctx any) *Packet {
	return &Packet{ctx: ctx}
}

func NewPacketDC(data []byte, ctx any) *Packet {
	return &Packet{
		data: data,
		ctx:  ctx,
	}
}

func NewPacketSC(data string, ctx any) *Packet {
	return NewPacketDC([]byte(data), ctx)
}

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
