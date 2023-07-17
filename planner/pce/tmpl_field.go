package pce

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/str"
	"strings"
)

// TmplField 模板字段
type TmplField struct {
	Name    string      // 字段名称
	Desc    string      // 字段描述
	Type    string      // 字段类型
	Struct  *TmplStruct // 结构类型字段结构信息
	Index   int         // 字段索引
	slice   bool        // 是否是切片类型
	isIndex bool        // 是否是索引字段
}

// IsIndex 是否是索引字段
func (slf *TmplField) IsIndex() bool {
	return slf.isIndex
}

// IsStruct 是否是结构类型
func (slf *TmplField) IsStruct() bool {
	return slf.Struct != nil
}

// IsSlice 是否是切片类型
func (slf *TmplField) IsSlice() bool {
	return slf.slice
}

// setStruct 设置结构类型
func (slf *TmplField) setStruct(parent, name, desc, fieldType string, fields map[string]Field) {
	slf.Struct = &TmplStruct{
		Name: fmt.Sprintf("%s%s", parent, name),
		Desc: desc,
	}
	slf.handleStruct(slf.Struct.Name, fieldType, fields)
}

// handleStruct 处理结构类型
func (slf *TmplField) handleStruct(fieldName, fieldType string, fields map[string]Field) {
	if strings.HasPrefix(fieldType, "[]") {
		slf.handleSlice(fieldName, fieldType, fields)
		return
	} else if !strings.HasPrefix(fieldType, "{") || !strings.HasSuffix(fieldType, "}") {
		return
	}
	var s = strings.TrimSuffix(strings.TrimPrefix(fieldType, "{"), "}")
	var fs []string
	var field string
	var leftBrackets []int
	for i, c := range s {
		switch c {
		case ',':
			if len(leftBrackets) == 0 {
				fs = append(fs, field)
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
				fs = append(fs, field)
				field = ""
			}
		default:
			field += string(c)
		}
	}
	if len(field) > 0 {
		fs = append(fs, field)
	}

	for _, fieldInfo := range fs {
		fieldName, fieldType := str.KV(strings.TrimSpace(fieldInfo), ":")
		slf.Struct.addField(slf.Struct.Name, str.FirstUpper(fieldName), fieldName, fieldType, fields)
	}

	slf.Type = slf.Struct.Name
}

// handleSlice 处理切片类型
func (slf *TmplField) handleSlice(fieldName, fieldType string, fields map[string]Field) {
	if !strings.HasPrefix(fieldType, "[]") {
		return
	}
	slf.slice = true
	t := strings.TrimPrefix(fieldType, "[]")
	if hash.Exist(fields, t) {
		slf.Struct = nil
		slf.Type = t
	} else {
		slf.handleStruct(fieldName, t, fields)
	}
}
