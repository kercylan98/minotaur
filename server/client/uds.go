package client

import (
	"net"
)

func NewUnixDomainSocket(addr string) *Client {
	return NewClient(&UnixDomainSocket{
		addr: addr,
	})
}

type UnixDomainSocket struct {
	conn   net.Conn
	addr   string
	closed bool
}

func (slf *UnixDomainSocket) Run(runState chan<- error, receive func(wst int, packet []byte)) {
	dial("unix", slf.addr, runState, receive, func(conn net.Conn) {
		slf.conn = conn
	}, func() bool {
		return slf.closed
	})
}

func (slf *UnixDomainSocket) Write(packet *Packet) error {
	_, err := slf.conn.Write(packet.data)
	return err
}

func (slf *UnixDomainSocket) Close() {
	slf.closed = true
}

func (slf *UnixDomainSocket) GetServerAddr() string {
	return slf.addr
}
