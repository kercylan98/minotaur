package raw

// StructureField 结构体字段描述
type StructureField struct {
	name string // 字段名称
	typ  string // 字段类型
}

// GetName 获取字段名称
func (sf StructureField) GetName() string {
	return sf.name
}

// GetType 获取字段类型
func (sf StructureField) GetType() string {
	return sf.typ
}
