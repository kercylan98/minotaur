package golang

import (
	"minotaur/exporter/configuration"
)

func NewField(name string, fieldType configuration.FieldType) *Field {
	return &Field{
		name:      name,
		fieldType: fieldType,
	}
}

type Field struct {
	id        int
	name      string
	fieldType configuration.FieldType
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
