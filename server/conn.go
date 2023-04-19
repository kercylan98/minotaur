package server

import (
	"github.com/gorilla/websocket"
	"github.com/panjf2000/gnet"
)

func newGNetConn(conn gnet.Conn) *Conn {
	return &Conn{
		ip: conn.RemoteAddr().String(),
		gn: conn,
		write: func(data []byte) error {
			return conn.AsyncWrite(data)
		},
	}
}

func newWebsocketConn(ws *websocket.Conn) *Conn {
	return &Conn{
		ws: ws,
		write: func(data []byte) error {
			return ws.WriteMessage(websocket.BinaryMessage, data)
		},
	}
}

type Conn struct {
	ip    string
	ws    *websocket.Conn
	gn    gnet.Conn
	write func(data []byte) error
}

func (slf *Conn) Write(data []byte) error {
	return slf.write(data)
}

func (slf *Conn) Close() {
	if slf.ws != nil {
		slf.ws.Close()
	} else if slf.gn != nil {
		slf.gn.Close()
	}
	slf.write = nil
}
