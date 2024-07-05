package transport

import "github.com/gofiber/contrib/websocket"

type FiberContext struct {
	conn *websocket.Conn
}

func (c *FiberContext) Locals(key string, value ...interface{}) any {
	return c.conn.Locals(key, value...)
}

func (c *FiberContext) Params(key string, defaultValue ...string) string {
	return c.conn.Params(key, defaultValue...)
}

func (c *FiberContext) Query(key string, defaultValue ...string) string {
	return c.conn.Query(key, defaultValue...)
}

func (c *FiberContext) Cookies(key string, defaultValue ...string) string {
	return c.conn.Cookies(key, defaultValue...)
}

func (c *FiberContext) Headers(key string, defaultValue ...string) string {
	return c.conn.Headers(key, defaultValue...)
}
