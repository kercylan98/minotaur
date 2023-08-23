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
	c, err := net.Dial("unix", slf.addr)
	if err != nil {
		runState <- err
		return
	}
	slf.conn = c
	runState <- nil
	packet := make([]byte, 1024)
	for !slf.closed {
		n, readErr := slf.conn.Read(packet)
		if readErr != nil {
			panic(readErr)
		}
		receive(0, packet[:n])
	}
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
