package raw

import (
	"fmt"
	"strings"
)

// AddField 添加字段到配置中
func AddField(config *Config, name, desc, typ, param string, index int, isKey bool, pos int) error {
	var f = Field{
		description: desc,
		name:        name,
		typ:         typ,
		param:       param,
		index:       index,
		key:         isKey,
		position:    pos,
	}

	_, exist := config.fieldStructures[name]
	if exist {
		return fmt.Errorf("field %s already exist", name)
	}

	config.fields[name] = f
	var err error
	config.fieldStructures[name], err = f.parseStructAlias()
	return err
}

// Field 原始字段描述
type Field struct {
	description string // 字段描述
	name        string // 字段名称
	typ         string // 字段类型
	param       string // 字段参数
	index       int    // 字段索引顺序
	key         bool   // 是否是索引字段
	position    int    // 字段在表中的位置（通常是列号或者行号）
}

// GetName 获取字段名称
func (f Field) GetName() string {
	return f.name
}

// GetDescription 获取字段描述
func (f Field) GetDescription() string {
	return f.description
}

// GetType 获取字段类型
func (f Field) GetType() string {
	var structs, err = f.parseStructAlias()
	if err != nil {
		panic(err)
	}
	for i, structure := range structs {
		if i == len(structs)-1 {
			return strings.ReplaceAll(f.typ, structure.rawType, structure.name)
		}
	}
	return f.typ
}

// GetParam 获取字段参数
func (f Field) GetParam() string {
	return f.param
}

// GetIndex 获取字段索引顺序
func (f Field) GetIndex() int {
	return f.index
}

// IsKey 是否是索引字段
func (f Field) IsKey() bool {
	return f.key
}

// GetPosition 获取字段在表中的位置
func (f Field) GetPosition() int {
	return f.position
}

// parseStruct 解析该字段中的结构体类型
func (f Field) parseStruct() ([]Structure, error) {
	// 列举:
	// string
	// [4]string
	// {id:int,name:string,info:{lv:int,exp:{mux:int,count:int}}}
	// []{id:int,name:string,info:{lv:int,exp:{mux:int,count:int}}}
	// map[string]string
	// map[string][]string
	// map[string]map[string]string
	// map[string]{id:int,name:string,info:{lv:int,exp:{mux:int,count:int}}}
	// map[string][]{id:int,name:string,info:{lv:int,exp:{mux:int,count:int}}}

	// 该函数目的是解析结构体类型的字段，即 {...} 部分
	var parser func(name, rawType string) ([]string, error)
	parser = func(name, rawType string) ([]string, error) {
		var res []string
		if IsBasicType(rawType) {
			return nil, nil
		}
		switch {
		case strings.HasPrefix(rawType, "[]"): // 切片
			parsed, err := parser(name, rawType[2:])
			if err != nil {
				return nil, err
			}
			res = append(res, parsed...)
		case strings.HasPrefix(rawType, "["): // 数组
			idx := strings.Index(rawType, "]")
			if idx == -1 {
				return nil, nil
			}
			parsed, err := parser(name, rawType[1:idx])
			if err != nil {
				return nil, err
			}
			res = append(res, parsed...)
		case strings.HasPrefix(rawType, "{"): // 结构体
			res = append(res, name+":"+rawType)
			var commaNum int
			for i := 1; i < len(rawType); i++ {
				c := rawType[i]
				if c == '{' {
					break
				} else if c == ',' {
					commaNum++
				}
			}
			fields := strings.SplitN(rawType[1:len(rawType)-1], ",", commaNum+1)
			for _, field := range fields {
				nameAndType := strings.SplitN(field, ":", 2)
				if len(nameAndType) != 2 {
					return nil, fmt.Errorf("invalid field %s", field)
				}
				parsed, err := parser(name+"_"+nameAndType[0], nameAndType[1])
				if err != nil {
					return nil, err
				}
				res = append(res, parsed...)
			}
		case strings.HasPrefix(rawType, "map["): // map
			// 找到第一个 ] 的位置
			idx := strings.Index(rawType, "]")
			if idx == -1 {
				return nil, nil
			}
			// 获取 key 的类型
			keyType := rawType[4:idx]
			if !IsBasicType(keyType) {
				return nil, fmt.Errorf("invalid map key type %s", keyType)
			}
			parsed, err := parser(name, rawType[idx+1:])
			if err != nil {
				return nil, err
			}
			res = append(res, parsed...)
		default:
			return nil, fmt.Errorf("invalid type %s", rawType)
		}

		return res, nil
	}
	// info:{id:int,name:string,info:{lv:int,exp:{mux:int,count:int}}}
	// info_info:{lv:int,exp:{mux:int,count:int}}
	// info_info_exp:{mux:int,count:int}
	structs, err := parser(f.name, f.typ)
	if err != nil {
		return nil, err
	}
	var res = make([]Structure, len(structs))
	for i := len(structs) - 1; i >= 0; i-- {
		nameAndType := strings.SplitN(structs[i], ":", 2)
		res[len(structs)-1-i] = Structure{
			name:    nameAndType[0],
			rawType: nameAndType[1],
		}
	}
	return res, nil
}

// parseStructAlias 解析该字段中的结构体类型并将嵌套结构体的类型替换为别名
func (f Field) parseStructAlias() ([]Structure, error) {
	structs, err := f.parseStruct()
	if err != nil || len(structs) == 0 {
		return nil, err
	}

	var rep = make([]Structure, len(structs))
	var curr = structs[0]
	curr.aliasType = curr.rawType
	rep[0] = curr
	for i := 1; i < len(structs); i++ {
		next := structs[i]
		rep[i] = Structure{
			name:      next.name,
			rawType:   next.rawType,
			aliasType: strings.ReplaceAll(next.rawType, curr.rawType, curr.name),
		}
		curr = next
	}

	return rep, nil
}
