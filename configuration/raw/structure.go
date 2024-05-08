package raw

import "strings"

// Structure 结构体描述
type Structure struct {
	name      string // 结构体名称
	rawType   string // 原始结构体类型
	aliasType string // 采用别名的结构体类型
}

// GetName 获取结构体名称
func (s Structure) GetName() string {
	return s.name
}

// GetFields 获取结构体字段描述
func (s Structure) GetFields() []StructureField {
	var res []StructureField
	for _, nameAndTypeStr := range strings.Split(s.aliasType[1:len(s.aliasType)-1], ",") {
		nameAndType := strings.SplitN(nameAndTypeStr, ":", 2)
		if len(nameAndType) != 2 {
			continue
		}

		name, typ := strings.TrimSpace(nameAndType[0]), strings.TrimSpace(nameAndType[1])
		res = append(res, StructureField{
			name: name,
			typ:  typ,
		})
	}
	return res
}
