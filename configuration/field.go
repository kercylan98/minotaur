package configuration

// Field 解析后的字段
type Field struct {
	Source    *SourceField // 配置源字段信息
	Structure *Structure   // 字段结构
}

// ExpandStructures 展开所有结构，将所有子结构作为字段返回
func (f *Field) ExpandStructures(preName ...string) []*Field {
	fields := make([]*Field, 0)

	switch f.Structure.Type {
	case SourceFieldStructureTypeSlice, SourceFieldStructureTypeArray:
		structureField := &Field{
			Source: &SourceField{
				DisplayName: "",
				Name:        f.Source.Name,
				Index:       -1,
				Structure:   f.Source.Structure,
				Type:        f.Source.Type,
			},
			Structure: f.Structure.Next,
		}
		// 数组字段如果下一级是结构体，那么展开
		if structureField.Structure.Type == SourceFieldStructureTypeStruct {
			fields = append(fields, structureField.ExpandStructures()...)
		}
	case SourceFieldStructureTypeMap:
		keyField := &Field{
			Source: &SourceField{
				DisplayName: "",
				Name:        f.Source.Name + "Key",
				Index:       -1,
				Structure:   f.Source.Structure,
				Type:        f.Source.Type,
			},
			Structure: f.Structure.MapKey,
		}
		valueField := &Field{
			Source: &SourceField{
				DisplayName: "",
				Name:        f.Source.Name + "Value",
				Index:       -1,
				Structure:   f.Source.Structure,
				Type:        f.Source.Type,
			},
			Structure: f.Structure.MapValue,
		}
		if len(preName) > 0 {
			keyField.Source.Name = preName[0] + ":" + keyField.Source.Name
			valueField.Source.Name = preName[0] + ":" + valueField.Source.Name
		}
		fields = append(fields, keyField)
		fields = append(fields, keyField.ExpandStructures(keyField.Source.Name)...)
		fields = append(fields, valueField)
		fields = append(fields, valueField.ExpandStructures(valueField.Source.Name)...)
	case SourceFieldStructureTypeStruct:
		structureField := &Field{
			Source: &SourceField{
				DisplayName: "",
				Name:        f.Source.Name,
				Index:       -1,
				Structure:   f.Source.Structure,
				Type:        f.Source.Type,
			},
			Structure: f.Structure,
		}
		if len(preName) > 0 {
			structureField.Source.Name = preName[0] + ":" + structureField.Source.Name
		}
		fields = append(fields, structureField)

		for _, structure := range f.Structure.StructFields {
			childField := &Field{
				Source: &SourceField{
					DisplayName: "",
					Name:        structureField.Source.Name + ":" + structure.Name,
					Index:       -1,
					Structure:   structure.Structure.Raw,
					Type:        f.Source.Type,
				},
				Structure: structure.Structure,
			}

			fields = append(fields, childField.ExpandStructures()...)
		}
	default:

	}

	return fields
}
