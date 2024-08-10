package effect

func newManagerConfiguration() *ManagerConfiguration {
	return &ManagerConfiguration{}
}

type ManagerConfiguration struct {
	defaultAttributes map[AttributeType]Attribute // 属性默认值
}

// WithDefaultAttribute 设置属性默认值
func (c *ManagerConfiguration) WithDefaultAttribute(attributeType AttributeType, attribute Attribute) *ManagerConfiguration {
	if c.defaultAttributes == nil {
		c.defaultAttributes = make(map[AttributeType]Attribute)
	}
	c.defaultAttributes[attributeType] = attribute
	return c
}

// WithDefaultAttributes 设置属性默认值
func (c *ManagerConfiguration) WithDefaultAttributes(attributes map[AttributeType]Attribute) *ManagerConfiguration {
	if c.defaultAttributes == nil {
		c.defaultAttributes = make(map[AttributeType]Attribute)
	}
	for attributeType, attribute := range attributes {
		c.defaultAttributes[attributeType] = attribute
	}
	return c
}
