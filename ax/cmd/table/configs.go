package table

import (
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
	}).gen()
}

type Configs struct {
	typeParser TypeParser
	codeParser CodeParser
	dataParser DataParser
	tables     []Table
	config     []*Config
}

func (cs *Configs) GenerateCode() string {
	return cs.codeParser.Parse(cs.config)
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

		// 结构体解析
		configStruct := &StructType{Name: table.GetName()}
		fs := table.GetFields()
		for {
			field := fs.Next()
			if field == nil {
				break
			}

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
					} else {
						config.types = append(config.types, v)
					}
				}
			}
		}
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
