package table

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"sort"
)

func GenerateConfigs(tables []Table, typeParser TypeParser, codeParser CodeParser, dataParser DataParser) *Configs {
	return (&Configs{
		tables:     tables,
		typeParser: typeParser,
		codeParser: codeParser,
		dataParser: dataParser,
		data:       make(map[string]map[string]any),
	}).gen()
}

type Configs struct {
	typeParser TypeParser
	codeParser CodeParser
	dataParser DataParser
	tables     []Table
	config     []*Config
	data       map[string]map[string]any
}

func (cs *Configs) GenerateCode() []byte {
	return []byte(cs.codeParser.Parse(cs.config))
}

func (cs *Configs) GenerateData() map[string]map[string]any {
	return cs.data
}

//goland:noinspection t
func (cs *Configs) gen() *Configs {
	for _, table := range cs.tables {
		config := &Config{
			name:      table.GetName(),
			desc:      table.GetDescribe(),
			fieldDesc: map[string]string{},
		}
		cs.config = append(cs.config, config)

		type indexField struct {
			idx int
			typ Type
		}
		var indexFields []indexField
		var fields []Field

		// 结构体解析
		configStruct := &StructType{Name: table.GetName()}
		fs := table.GetFields()
		for {
			field := fs.Next()
			if field == nil {
				break
			}
			if field.IsIgnore() {
				continue
			}
			fields = append(fields, field)

			config.fieldDesc[field.GetName()] = field.GetDesc()

			types := cs.typeParser.Parse(field.GetType())
			collection.ReverseSlice(&types)
			var parent string
			for i, t := range types {
				if i == 0 && field.GetIndex() > 0 {
					indexFields = append(indexFields, indexField{idx: field.GetIndex(), typ: t})
				}
				switch v := t.(type) {
				case *BasicType:
					configStruct.Fields = append(configStruct.Fields, &StructField{Name: field.GetName(), Type: v})
				case *ArrayType:
					configStruct.Fields = append(configStruct.Fields, &StructField{Name: field.GetName(), Type: v})
				case *StructType:
					if v.Name == "" {
						v.Name = config.name + "_" + field.GetName()
					} else {
						v.Name = parent + "_" + v.Name
					}
					v.Name = charproc.BigCamel(v.Name)
					parent = v.Name
					if i == 0 {
						configStruct.Fields = append(configStruct.Fields, &StructField{
							Name: field.GetName(),
							Type: v,
						})
					}
					config.types = append(config.types, v)
				}
			}
		}

		// 数据解析
		var tableData = make(map[string]any)
		var rowNo int
	scan:
		for {
			var indexes []Field
			var indexValue = make(map[string]map[string]any)
			var row = make(map[string]any)
		scanRow:
			for _, field := range fields {
				val, skip, end := field.Query(rowNo)
				if end {
					break scan
				}
				if skip {
					rowNo++
					continue scanRow
				}
				toolkit.MarshalToTargetWithJSON(val, &row)
				if field.GetIndex() > 0 {
					indexes = append(indexes, field)
					indexValue[field.GetName()] = val
				}
			}

			if len(indexes) == 0 {
				tableData = row
			}

			// index 处理
			sort.Slice(indexes, func(i, j int) bool {
				return indexes[i].GetIndex() >= indexes[j].GetIndex()
			})
			for i, index := range indexes {
				if i != len(indexes)-1 {
					row = map[string]any{
						fmt.Sprint(indexValue[index.GetName()][index.GetName()]): row,
					}
				} else {
					topKey := fmt.Sprint(indexValue[index.GetName()][index.GetName()])
					children, exist := tableData[topKey].(map[string]any)
					if !exist {
						children = make(map[string]any)
						tableData[topKey] = children
					}
					for k, v := range row {
						children[k] = v
					}
				}
			}

			rowNo++
		}

		cs.data[table.GetName()] = tableData

		sort.Slice(indexFields, func(i, j int) bool {
			return indexFields[i].idx < indexFields[j].idx // 升序
		})
		for _, f := range indexFields {
			config.indexes = append(config.indexes, f.typ)
		}
		config.types = append(config.types, configStruct)
	}

	return cs
}
