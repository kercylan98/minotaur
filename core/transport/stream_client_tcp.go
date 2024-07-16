package transport

import (
	"bytes"
	"net"
)

var _ StreamClientCore = (*StreamClientTCPCore)(nil)

type StreamClientTCPCore struct {
	Addr   string
	conn   net.Conn
	buffer bytes.Buffer
}

func (c *StreamClientTCPCore) OnConnect() error {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

func (c *StreamClientTCPCore) OnRead() (Packet, error) {
	c.buffer.Reset()
	_, err := c.buffer.ReadFrom(c.conn)
	if err != nil {
		return nil, err
	}

	return NewPacket(c.buffer.Bytes()), nil
}

func (c *StreamClientTCPCore) OnWrite(p Packet) error {
	_, err := c.conn.Write(p.GetBytes())
	return err
}

func (c *StreamClientTCPCore) OnClose() error {
	return c.conn.Close()
}
