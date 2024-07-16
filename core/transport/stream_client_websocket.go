package transport

import (
	"github.com/fasthttp/websocket"
	"net/http"
)

var _ StreamClientCore = (*StreamClientWebSocketCore)(nil)

type StreamClientWebSocketCore struct {
	Url    string
	Dialer *websocket.Dialer // 默认将使用 websocket.DefaultDialer
	Header http.Header
	conn   *websocket.Conn
}

func (c *StreamClientWebSocketCore) OnConnect() error {
	if c.Dialer == nil {
		c.Dialer = websocket.DefaultDialer
	}
	conn, _, err := c.Dialer.Dial(c.Url, c.Header)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *StreamClientWebSocketCore) OnRead() (Packet, error) {
	messageType, data, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	return NewPacket(data).SetContext(messageType), nil
}

func (c *StreamClientWebSocketCore) OnWrite(p Packet) error {
	return c.conn.WriteMessage(p.GetContext().(int), p.GetBytes())
}

func (c *StreamClientWebSocketCore) OnClose() error {
	return c.conn.Close()
}
