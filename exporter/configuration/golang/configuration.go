package golang

import (
	"minotaur/exporter/configuration"
	"minotaur/utils/hash"
)

func NewConfiguration(name string) *Configuration {
	return &Configuration{
		name:   name,
		fields: hash.NewSortMap[int, configuration.Field](),
	}
}

type Configuration struct {
	name   string
	fields *hash.SortMap[int, configuration.Field]
}

func (slf *Configuration) GetName() string {
	return slf.name
}

func (slf *Configuration) GetFields() []configuration.Field {
	return slf.fields.ToSliceSort()
}

func (slf *Configuration) AddField(field configuration.Field) {
	slf.fields.Set(field.GetID(), field)
}
