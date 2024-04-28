package pce

import (
"github.com/kercylan98/minotaur/toolkit/collection"
)

// TmplStruct 模板结构
type TmplStruct struct {
	Name       string       // 结构名称
	Desc       string       // 结构描述
	Fields     []*TmplField // 字段列表
	IndexCount int          // 索引数量
}

// addField 添加字段
func (slf *TmplStruct) addField(parent, name, desc, fieldType string, fields map[string]Field) *TmplField {
	field := &TmplField{
		Name: name,
		Desc: desc,
		Type: fieldType,
	}
	if !collection.KeyInMap(fields, fieldType) {
		field.setStruct(parent, name, desc, fieldType, fields)
	} else {
		field.Type = GetFieldGolangType(fields[fieldType])
	}
	slf.Fields = append(slf.Fields, field)
	return field
}

// AllChildren 获取所有子结构
func (slf *TmplStruct) AllChildren() []*TmplStruct {
	var children []*TmplStruct
	for _, field := range slf.Fields {
		if field.IsStruct() {
			children = append(children, field.Struct)
			children = append(children, field.Struct.AllChildren()...)
		}
	}
	return children
}
