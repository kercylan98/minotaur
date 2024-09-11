package socket

func newPacket(data []byte, ctx any) *Packet {
	return &Packet{data: data, ctx: ctx}
}

type Packet struct {
	data []byte
	ctx  any
}

func (p *Packet) GetData() []byte {
	return p.data
}

func (p *Packet) GetContext() any {
	return p.ctx
}
