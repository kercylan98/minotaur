package traffickers

import (
	"github.com/panjf2000/gnet/v2"
	"time"
)

type websocketConn struct {
	gnet.Conn
	deadline time.Time
}

func (c *websocketConn) SetDeadline(t time.Time) error {
	c.deadline = t
	return nil
}
