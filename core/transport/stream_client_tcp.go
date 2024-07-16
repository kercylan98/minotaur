package transport

import (
	"net"
)

var _ StreamClientCore = (*StreamClientTCPCore)(nil)

type StreamClientTCPCore struct {
	Addr   string
	conn   net.Conn
	buffer []byte
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
	c.buffer = make([]byte, 1024)
	_, err := c.conn.Read(c.buffer)
	return NewPacket(c.buffer), err
}

func (c *StreamClientTCPCore) OnWrite(p Packet) error {
	_, err := c.conn.Write(p.GetBytes())
	return err
}

func (c *StreamClientTCPCore) OnClose() error {
	return c.conn.Close()
}
