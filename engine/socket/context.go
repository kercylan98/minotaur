package socket

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
	if c.data == nil {
		return nil
	}
	return c.data[key]
}

func (c *Context) Has(key string) bool {
	if c.data == nil {
		return false
	}
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
	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}
