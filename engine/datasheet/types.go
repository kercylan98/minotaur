package datasheet

type BasicType struct {
	TypeName string // 基本类型名称
}

type Struct struct {
	StructName string   // 结构体名称
	Anonymity  bool     // 匿名结构体
	Fields     []*Field // 字段
}

type Array struct {
	Length    int  // 长度
	ValueType Type // 数组元素类型
}

type Slice struct {
	ValueType Type
}

type Map struct {
	KeyType   Type
	ValueType Type
}

type Field struct {
	Owner     *Struct // 宿主
	FieldName string  // 字段名称
	FieldType Type    // 字段类型
}
