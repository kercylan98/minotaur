package client

import (
	"github.com/gorilla/websocket"
	"github.com/kercylan98/minotaur/server"
)

// NewWebsocket 创建 websocket 客户端
func NewWebsocket(addr string) *Client {
	return NewClient(&Websocket{
		addr: addr,
	})
}

// Websocket websocket 客户端
type Websocket struct {
	addr   string
	conn   *websocket.Conn
	closed bool
}

func (slf *Websocket) Run(runState chan<- error, receive func(wst int, packet []byte)) {
	ws, _, err := websocket.DefaultDialer.Dial(slf.addr, nil)
	if err != nil {
		runState <- err
		return
	}
	slf.conn = ws
	slf.closed = false
	runState <- nil
	for !slf.closed {
		messageType, packet, readErr := ws.ReadMessage()
		if readErr != nil {
			panic(readErr)
		}
		receive(messageType, packet)
	}
}

func (slf *Websocket) Write(packet *Packet) error {
	if packet.wst == 0 {
		packet.wst = server.WebsocketMessageTypeBinary
	}
	return slf.conn.WriteMessage(packet.wst, packet.data)
}

func (slf *Websocket) Close() {
	slf.closed = true
}

func (slf *Websocket) GetServerAddr() string {
	return slf.addr
}

func (slf *Websocket) Clone() Core {
	return &Websocket{
		addr: slf.addr,
	}
}
