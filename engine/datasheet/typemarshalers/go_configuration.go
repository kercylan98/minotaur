package typemarshalers

func newGoConfiguration() *GoConfiguration {
	return &GoConfiguration{}
}

type GoConfiguration struct {
	typePrefix bool // 是否需要类型前缀
}

func (c *GoConfiguration) WithTypePrefix() *GoConfiguration {
	c.typePrefix = true
	return c
}
