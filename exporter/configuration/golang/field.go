package golang

import (
	"minotaur/exporter/configuration"
)

var fieldTypeMapper = map[string]configuration.FieldType{
	"string":  configuration.FieldTypeString,
	"int":     configuration.FieldTypeInt,
	"int64":   configuration.FieldTypeInt64,
	"float32": configuration.FieldTypeFloat32,
	"float64": configuration.FieldTypeFloat64,
	"bool":    configuration.FieldTypeBool,
	"byte":    configuration.FieldTypeByte,
}

func GetFieldType(fieldType string) configuration.FieldType {
	return fieldTypeMapper[fieldType]
}

func NewField(id int, name string, fieldType configuration.FieldType, isIndex bool) *Field {
	return &Field{
		id:        id,
		name:      name,
		fieldType: fieldType,
		isIndex:   isIndex,
	}
}

type Field struct {
	id        int
	name      string
	fieldType configuration.FieldType
	isIndex   bool
}

func (slf *Field) GetID() int {
	return slf.id
}

func (slf *Field) GetName() string {
	return slf.name
}

func (slf *Field) GetType() configuration.FieldType {
	return slf.fieldType
}

func (slf *Field) IsIndex() bool {
	return slf.isIndex
}
