package transport

import (
	"github.com/fasthttp/websocket"
	"github.com/kercylan98/minotaur/core/vivid"
	"net/url"
)

func NewWebSocketClient(u url.URL, config ...WebSocketClientConfig) *WebSocketClient {
	wsc := &WebSocketClient{
		url: u,
	}
	if len(config) > 0 {
		wsc.config = config[0]
	} else {
		wsc.config = WebSocketClientConfig{}
	}
	return wsc
}

type WebSocketClient struct {
	url    url.URL
	config WebSocketClientConfig
	conn   *websocket.Conn
	err    error
}

func (c *WebSocketClient) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case vivid.OnLaunch:
		c.onLaunch(ctx)
	case ReactPacket:
		if c.config.ConnectionPacketHandler != nil {
			c.config.ConnectionPacketHandler(ctx, m)
		}
	case Packet:
		c.onWritePacket(ctx, m)
	case vivid.FutureForwardMessage:
		ctx.Tell(ctx.Ref(), m.Error)
	case error:
		c.err = m
		ctx.Terminate(ctx.Ref())
	case vivid.OnTerminate:
		c.onTerminate(ctx)
	}
}

func (c *WebSocketClient) onLaunch(ctx vivid.ActorContext) {
	conn, _, err := websocket.DefaultDialer.Dial(c.url.String(), c.config.Header)
	if err != nil {
		panic(err)
	}
	c.conn = conn

	if c.config.ConnectionOpenedHandler != nil {
		c.config.ConnectionOpenedHandler(ctx)
	}

	ctx.AwaitForward(ctx.Ref(), func() vivid.Message {
		var messageType int
		var data []byte
		for {
			messageType, data, err = c.conn.ReadMessage()
			if err != nil {
				return err
			}

			ctx.Tell(ctx.Ref(), ReactPacket{NewPacket(data).SetContext(messageType)})
		}
	})
}

func (c *WebSocketClient) onWritePacket(ctx vivid.ActorContext, m Packet) {
	if err := c.conn.WriteMessage(m.GetContext().(int), m.GetBytes()); err != nil {
		ctx.Tell(ctx.Ref(), err)
	}
}

func (c *WebSocketClient) onTerminate(ctx vivid.ActorContext) {
	if c.config.ConnectionClosedHandler != nil {
		c.config.ConnectionClosedHandler(ctx, c.err)
	}
	if c.conn != nil {
		_ = c.conn.Close()
	}
}
