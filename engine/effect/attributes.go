package effect

type AttributeType = uint32 // 属性类型

type Attributes map[AttributeType]Attribute // 属性集合

// Get 获取特定属性的属性值
func (as Attributes) Get(attributeType AttributeType) Attribute {
	return as[attributeType]
}

// Set 设置特定属性的属性值
func (as Attributes) Set(attributeType AttributeType, value Attribute) {
	as[attributeType] = value
}
