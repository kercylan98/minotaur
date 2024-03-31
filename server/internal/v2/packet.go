package server

func NewPacket(data []byte) Packet {
	return new(packet).init(data)
}

type Packet interface {
	GetBytes() []byte
	SetContext(ctx any) Packet
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
