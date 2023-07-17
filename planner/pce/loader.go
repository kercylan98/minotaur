package pce

import (
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tidwall/gjson"
	"strings"
)

// NewLoader 创建加载器
//   - 加载器被用于加载配置表的数据和结构信息
func NewLoader(fields []Field) *Loader {
	loader := &Loader{
		fields: make(map[string]Field),
	}
	for _, f := range fields {
		loader.fields[f.TypeName()] = f
	}
	return loader
}

// Loader 配置加载器
type Loader struct {
	fields map[string]Field
}

// LoadStruct 加载结构
func (slf *Loader) LoadStruct(config Config) *TmplStruct {
	var tmpl = &TmplStruct{
		Name:       str.FirstUpper(config.GetConfigName()),
		Desc:       config.GetDescription(),
		IndexCount: config.GetIndexCount(),
	}
	for i, field := range config.GetFields() {
		f := tmpl.addField(tmpl.Name, str.FirstUpper(field.Name), field.Desc, field.Type, slf.fields)
		if i < tmpl.IndexCount {
			f.isIndex = true
		}
	}
	return tmpl
}

// LoadData 加载配置并得到配置数据
func (slf *Loader) LoadData(config Config) map[any]any {
	var source = make(map[any]any)
	var action = source
	var indexCount = config.GetIndexCount() - 1
	for _, row := range config.GetData() {
		var item = make(map[any]any)
		for index, field := range row {
			bind, exist := slf.fields[field.Type]
			var value any
			if exist {
				if len(field.Value) == 0 {
					value = bind.Zero()
				} else {
					value = bind.Parse(field.Value)
				}
			} else {
				value = slf.structInterpreter(field.Type, field.Value)
			}

			if value != nil {
				item[field.Name] = value
			}

			if index < indexCount {
				m, exist := action[value]
				if !exist {
					m = map[any]any{}
					action[value] = m
					action = m.(map[any]any)
				} else {
					action = m.(map[any]any)
				}
			} else if index == indexCount {
				action[value] = item
			} else if indexCount == -1 {
				source = item
			}
		}
		action = source
	}

	return source
}

// sliceInterpreter 切片解释器
func (slf *Loader) sliceInterpreter(fieldType, fieldValue string) any {
	if !strings.HasPrefix(fieldType, "[]") {
		return nil
	}
	t := strings.TrimPrefix(fieldType, "[]")
	var data = map[any]any{}
	gjson.ForEachLine(fieldValue, func(line gjson.Result) bool {
		line.ForEach(func(key, value gjson.Result) bool {
			field, exist := slf.fields[t]
			if exist {
				data[len(data)] = field.Parse(value.String())
			} else {
				data[len(data)] = slf.structInterpreter(t, value.String())
			}
			return true
		})
		return true
	})
	return data
}

// structInterpreter 结构体解释器
//   - {id:int,name:string,info:{lv:int,exp:int}}
func (slf *Loader) structInterpreter(fieldType, fieldValue string) any {
	if strings.HasPrefix(fieldType, "[]") {
		return slf.sliceInterpreter(fieldType, fieldValue)
	} else if !strings.HasPrefix(fieldType, "{") || !strings.HasSuffix(fieldType, "}") {
		return nil
	}
	var s = strings.TrimSuffix(strings.TrimPrefix(fieldType, "{"), "}")
	var data = map[any]any{}
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
		field, exist := slf.fields[fieldType]
		if exist {
			data[fieldName] = field.Parse(gjson.Get(fieldValue, fieldName).String())
		} else {
			data[fieldName] = slf.structInterpreter(fieldType, gjson.Get(fieldValue, fieldName).String())
		}
	}

	return data
}

// DataInfo 配置数据
type DataInfo struct {
	DataField        // 字段
	Value     string // 字段值
}

// DataField 配置数据字段
type DataField struct {
	Index      int    // 字段索引
	Name       string // 字段名称
	Desc       string // 字段描述
	Type       string // 字段类型
	ExportType string // 导出类型
}
