package client

import "net"

func NewTCP(addr string) *Client {
	return NewClient(&TCP{
		addr: addr,
	})
}

type TCP struct {
	conn   net.Conn
	addr   string
	closed bool
}

func (slf *TCP) Run(runState chan<- error, receive func(wst int, packet []byte)) {
	dial("tcp", slf.addr, runState, receive, func(conn net.Conn) {
		slf.conn = conn
	}, func() bool {
		return slf.closed
	})
}

func (slf *TCP) Write(packet *Packet) error {
	_, err := slf.conn.Write(packet.data)
	return err
}

func (slf *TCP) Close() {
	slf.closed = true
}

func (slf *TCP) GetServerAddr() string {
	return slf.addr
}

func (slf *TCP) Clone() Core {
	return &TCP{
		addr: slf.addr,
	}
}
