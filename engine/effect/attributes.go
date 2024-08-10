package effect

type AttributeType = uint32 // 属性类型

func newAttributes(manager *Manager) *Attributes {
	return &Attributes{
		manager:    manager,
		attributes: make(map[AttributeType]Attribute),
	}
}

type Attributes struct {
	manager    *Manager
	attributes map[AttributeType]Attribute
}

func (as *Attributes) Clone() *Attributes {
	c := newAttributes(as.manager)
	for attributeType, attribute := range as.attributes {
		c.attributes[attributeType] = attribute
	}
	return c
}

func (as *Attributes) Get(attributeType AttributeType) Attribute {
	attr, exist := as.attributes[attributeType]
	if exist {
		return attr
	}
	return as.manager.config.defaultAttributes[attributeType]
}

func (as *Attributes) Set(attributeType AttributeType, attribute Attribute) {
	as.attributes[attributeType] = attribute
}
