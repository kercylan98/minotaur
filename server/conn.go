package server

import (
	"github.com/gorilla/websocket"
	"github.com/panjf2000/gnet"
	"github.com/xtaci/kcp-go/v5"
)

func newKcpConn(session *kcp.UDPSession) *Conn {
	return &Conn{
		ip:  session.RemoteAddr().String(),
		kcp: session,
		write: func(data []byte) error {
			_, err := session.Write(data)
			return err
		},
	}
}

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
	kcp   *kcp.UDPSession
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
	} else if slf.kcp != nil {
		slf.kcp.Close()
	}
	slf.write = nil
}
