package server

type Packet interface {
	GetBytes() []byte
	SetContext(ctx any)
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

func (m *packet) SetContext(ctx any) {
	m.ctx = ctx
}

func (m *packet) GetContext() any {
	return m.ctx
}
