package rpc

type contextReader func(dst any) error

// Context 是一个 RPC 调用上下文的接口，该接口用于在 RPC 调用过程中传递上下文信息
type Context interface {

	// ReadTo 将上下文信息读取到指定的目标中
	ReadTo(dst any) error

	// MustReadTo 将上下文信息读取到指定的目标中，如果读取失败则会 panic
	MustReadTo(dst any)

	// Reply 将指定的数据写入到上下文信息中
	Reply(src any)
}

// newContext 用于创建一个新的 RPC 调用上下文
func newContext(reader contextReader) *context {
	return &context{
		reader: reader,
	}
}

type context struct {
	reply  any
	reader contextReader
}

func (c *context) ReadTo(dst any) error {
	return c.reader(dst)
}

func (c *context) MustReadTo(dst any) {
	if err := c.reader(dst); err != nil {
		panic(err)
	}
}

func (c *context) Reply(src any) {
	c.reply = src
}
