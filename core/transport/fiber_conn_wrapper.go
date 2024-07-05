package transport

import (
	"github.com/gofiber/contrib/websocket"
	"time"
)

type fiberConnWrapper struct {
	*websocket.Conn
}

func (f fiberConnWrapper) Read(b []byte) (n int, err error) {
	panic("implement me")
}

func (f fiberConnWrapper) Write(b []byte) (n int, err error) {
	panic("implement me")
}

func (f fiberConnWrapper) SetDeadline(t time.Time) error {
	panic("implement me")
}
