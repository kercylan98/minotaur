package server

import (
	"context"
	"github.com/kercylan98/minotaur/utils/log"
	"net"
	"unsafe"
)

type Conn interface {
	net.Conn
}

type conn struct {
	net.Conn
	cs     *connections
	ctx    context.Context
	cancel context.CancelFunc
	idx    int
}

func (c *conn) init(ctx context.Context, cs *connections, conn net.Conn, idx int) *conn {
	c.Conn = conn
	c.cs = cs
	c.ctx, c.cancel = context.WithCancel(ctx)
	c.idx = idx
	return c
}

func (c *conn) awaitRead() {
	defer func() { _ = c.Close() }()

	const bufferSize = 4096
	buf := make([]byte, bufferSize) // 避免频繁的内存分配，初始化一个固定大小的缓冲区
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			ptr := unsafe.Pointer(&buf[0])
			n, err := c.Read((*[bufferSize]byte)(ptr)[:])
			if err != nil {
				log.Error("READ", err)
				return
			}

			if n > 0 {
				if _, err := c.Write(buf[:n]); err != nil {
					log.Error("Write", err)
				}
			}
		}
	}
}

func (c *conn) Close() (err error) {
	c.cs.Event() <- c
	return
}
