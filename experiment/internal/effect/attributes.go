package effect

type AttributeType = int32 // 属性类型

func NewAttributes() *Attributes {
	return &Attributes{
		attributes: make(map[AttributeType]Attribute),
	}
}

type Attributes struct {
	attributes map[AttributeType]Attribute
}

func (as *Attributes) Clone() *Attributes {
	c := NewAttributes()
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
	return Attribute{}
}

func (as *Attributes) Set(attributeType AttributeType, attribute Attribute) {
	as.attributes[attributeType] = attribute
}
