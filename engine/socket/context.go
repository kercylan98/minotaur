package socket

// ContextEditor 是用于对 Socket 上下文进行编辑的接口，它将允许修改器提供的数据。
type ContextEditor interface {
	OnEdit(ctx *Context)
}

type FunctionalContextEditor func(ctx *Context)

func (f FunctionalContextEditor) OnEdit(ctx *Context) {
	f(ctx)
}

func newContext() *Context {
	return &Context{}
}

type Context struct {
	data map[string]any
}

func (c *Context) Set(key string, value any) {
	if c.data == nil {
		c.data = make(map[string]any)
	}
	c.data[key] = value
}

func (c *Context) Get(key string) any {
	return c.data[key]
}

func (c *Context) Has(key string) bool {
	_, ok := c.data[key]
	return ok
}

func (c *Context) Clear() {
	c.data = nil
}

func (c *Context) Delete(key string) {
	delete(c.data, key)
}

func (c *Context) Keys() []string {
	if len(c.data) == 0 {
		return nil
	}
	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}
