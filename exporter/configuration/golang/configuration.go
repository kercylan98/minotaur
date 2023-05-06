package golang

import (
	"fmt"
	"github.com/kercylan98/minotaur/exporter/configuration"
	"github.com/kercylan98/minotaur/utils/hash"
	"strings"
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

func (slf *Configuration) GetStruct() string {
	var index string
	var structStr string
	for _, field := range slf.GetFields() {
		if field.IsIndex() {
			index += fmt.Sprintf("map[%s]", field.GetType())
		}
		structStr += fmt.Sprintf("%s %s\r\n", strings.ToUpper(field.GetName()[:1])+field.GetName()[1:], field.GetType().String())
	}
	index = fmt.Sprintf("%s*%s", index, strings.ToUpper(slf.GetName()[:1])+slf.GetName()[1:])
	structStr = fmt.Sprintf("type %s struct {\r\n %s \r\n}", strings.ToUpper(slf.GetName()[:1])+slf.GetName()[1:], structStr)
	return fmt.Sprintf("%s \r\n%s", index, structStr)
}
