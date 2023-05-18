package internal

import (
	"bytes"
	"fmt"
	"github.com/kercylan98/minotaur/utils/str"
	"html/template"
	"strings"
)

func NewField(configName, fieldName, fieldType string) *Field {
	sourceType := fieldType
	var handleStruct = func(configName, fieldType string) []*Field {
		var fs []*Field
		var s = strings.TrimSuffix(strings.TrimPrefix(fieldType, "{"), "}")
		var fields []string
		var field string
		var leftBrackets []int
		for i, c := range s {
			switch c {
			case ',':
				if len(leftBrackets) == 0 {
					fields = append(fields, field)
					field = ""
				} else {
					field += string(c)
				}
			case '{':
				leftBrackets = append(leftBrackets, i)
				field += string(c)
			case '}':
				leftBrackets = leftBrackets[:len(leftBrackets)-1]
				field += string(c)
				if len(leftBrackets) == 0 {
					fields = append(fields, field)
					field = ""
				}
			default:
				field += string(c)
			}
		}
		if len(field) > 0 {
			fields = append(fields, field)
		}

		for _, fieldInfo := range fields {
			fieldName, fieldType := str.KV(strings.TrimSpace(fieldInfo), ":")
			fs = append(fs, NewField(configName, str.FirstUpper(fieldName), fieldType))

		}
		return fs
	}

	var fs []*Field
	var sliceField *Field
	if t, exist := basicTypeName[fieldType]; exist {
		fieldType = t
		goto end
	}
	if strings.HasPrefix(fieldType, "[]") {
		fieldType = strings.TrimPrefix(fieldType, "[]")
		if t, exist := basicTypeName[fieldType]; exist {
			fieldType = fmt.Sprintf("map[int]%s", t)
		} else {
			sliceField = NewField(configName, fieldName, fieldType)
			fieldType = fmt.Sprintf("map[int]*%s", configName+fieldName)
		}
		goto end
	}

	fs = handleStruct(configName+fieldName, fieldType)
	fieldType = fmt.Sprintf("*%s", configName+fieldName)

end:
	{
		return &Field{
			Name:        fieldName,
			Type:        fieldType,
			SourceType:  sourceType,
			TypeNotStar: strings.TrimPrefix(fieldType, "*"),
			Fields:      fs,
			SliceField:  sliceField,
		}
	}
}

type Field struct {
	Name        string
	Type        string
	SourceType  string
	TypeNotStar string
	Fields      []*Field
	SliceField  *Field
	ExportParam string
	Server      bool
	Describe    string
	Ignore      bool
}

func (slf *Field) Struct() string {
	if len(slf.Fields) == 0 && slf.SliceField == nil {
		return ""
	}
	tmpl, err := template.New("struct").Parse(generateGoStructTemplate)
	if err != nil {
		return ""
	}
	var buf bytes.Buffer
	if len(slf.Fields) > 0 {
		if err = tmpl.Execute(&buf, slf); err != nil {
			return ""
		}
	} else if slf.SliceField != nil {
		if err = tmpl.Execute(&buf, slf.SliceField); err != nil {
			return ""
		}
	}
	s := buf.String()
	return s
}

func (slf *Field) String() string {
	if len(slf.Fields) == 0 && slf.SliceField == nil {
		return ""
	}
	tmpl, err := template.New("struct").Parse(generateGoStructTemplate)
	if err != nil {
		return ""
	}

	var buf bytes.Buffer
	if len(slf.Fields) > 0 {
		if err = tmpl.Execute(&buf, slf); err != nil {
			return ""
		}
	} else if slf.SliceField != nil {
		if err = tmpl.Execute(&buf, slf.SliceField); err != nil {
			return ""
		}
	}
	s := buf.String()
	for _, field := range slf.Fields {
		s += fmt.Sprintf("%s", field.String())
	}
	return s
}
