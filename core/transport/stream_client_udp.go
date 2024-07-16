package transport

import (
	"bytes"
	"net"
)

var _ StreamClientCore = (*StreamClientUDPCore)(nil)

type StreamClientUDPCore struct {
	Addr   string
	conn   *net.UDPConn
	buffer bytes.Buffer
}

func (c *StreamClientUDPCore) OnConnect() error {
	s, err := net.ResolveUDPAddr("udp4", c.Addr)
	if err != nil {
		return err
	}
	c.conn, err = net.DialUDP("udp4", nil, s)
	if err != nil {
		return err
	}
	return nil
}

func (c *StreamClientUDPCore) OnRead() (Packet, error) {
	c.buffer.Reset()
	_, err := c.buffer.ReadFrom(c.conn)
	if err != nil {
		return nil, err
	}

	return NewPacket(c.buffer.Bytes()), nil
}

func (c *StreamClientUDPCore) OnWrite(p Packet) error {
	_, err := c.conn.Write(p.GetBytes())
	return err
}

func (c *StreamClientUDPCore) OnClose() error {
	return c.conn.Close()
}
