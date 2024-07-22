package stream

func NewPacket() *Packet {
	return &Packet{}
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

type Packet struct {
	data []byte // 数据包数据
	ctx  any    // 数据包上下文
}

func (p *Packet) SetData(data []byte) *Packet {
	p.data = data
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
