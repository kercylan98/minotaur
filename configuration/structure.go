package configuration

import (
	"errors"
	"strconv"
	"strings"
)

const (
	SourceFieldStructureTypeInvalid StructureType = iota
	SourceFieldStructureTypeString
	SourceFieldStructureTypeByte
	SourceFieldStructureTypeRune
	SourceFieldStructureTypeInt
	SourceFieldStructureTypeInt8
	SourceFieldStructureTypeInt32
	SourceFieldStructureTypeInt64
	SourceFieldStructureTypeUint
	SourceFieldStructureTypeUint8
	SourceFieldStructureTypeUint32
	SourceFieldStructureTypeUint64
	SourceFieldStructureTypeFloat32
	SourceFieldStructureTypeFloat64
	SourceFieldStructureTypeBool
	SourceFieldStructureTypeTime
	SourceFieldStructureTypeDuration
	SourceFieldStructureTypeSlice
	SourceFieldStructureTypeArray
	SourceFieldStructureTypeMap
	SourceFieldStructureTypeStruct
)

var (
	ErrInvalidStructure = errors.New("invalid structure")
)

// StructureType 代表了配置源字段的结构类型
type StructureType = int32

// ParseStructure 解析字段结构
func ParseStructure(structureType string) (structure *Structure, err error) {
	structureType = strings.TrimSpace(structureType)
	lower := strings.ToLower(structureType)
	structure = new(Structure)
	structure.Raw = structureType
	switch lower {
	case "string", "str", "text", "varchar":
		structure.Type = SourceFieldStructureTypeString
	case "byte":
		structure.Type = SourceFieldStructureTypeByte
	case "rune":
		structure.Type = SourceFieldStructureTypeRune
	case "int":
		structure.Type = SourceFieldStructureTypeInt
	case "int8":
		structure.Type = SourceFieldStructureTypeInt8
	case "int32":
		structure.Type = SourceFieldStructureTypeInt32
	case "int64":
		structure.Type = SourceFieldStructureTypeInt64
	case "uint":
		structure.Type = SourceFieldStructureTypeUint
	case "uint8":
		structure.Type = SourceFieldStructureTypeUint8
	case "uint32":
		structure.Type = SourceFieldStructureTypeUint32
	case "uint64":
		structure.Type = SourceFieldStructureTypeUint64
	case "float32":
		structure.Type = SourceFieldStructureTypeFloat32
	case "float64", "double", "number":
		structure.Type = SourceFieldStructureTypeFloat64
	case "bool":
		structure.Type = SourceFieldStructureTypeBool
	default:
		var next string
		switch {
		case strings.HasPrefix(lower, "[]"): // 切片
			structure.Type = SourceFieldStructureTypeSlice
			next = structureType[2:]
		case strings.HasPrefix(lower, "["): // 数组
			var i = strings.IndexByte(structureType, ']')
			if i < 0 {
				return nil, ErrInvalidStructure
			}
			length, err := strconv.Atoi(structureType[1:i])
			if err != nil {
				return nil, err
			}

			structure.ArrayLength = length
			structure.Type = SourceFieldStructureTypeArray
			next = structureType[i+1:]
		case strings.HasPrefix(lower, "map["):
			var i = strings.IndexByte(structureType, ']')
			if i < 0 {
				return nil, ErrInvalidStructure
			}
			key := structureType[4:i]
			value := structureType[i+1:]
			keyStructure, err := ParseStructure(key)
			if err != nil {
				return nil, err
			}

			valueStructure, err := ParseStructure(value)
			if err != nil {
				return nil, err
			}

			structure.MapKey = keyStructure
			structure.MapValue = valueStructure
			structure.Type = SourceFieldStructureTypeMap
			next = structureType[4:]
		case strings.HasPrefix(lower, "{"):
			// {id:int,name:string,info:{lv:int,exp:{mux:int,count:int}}}
			var i = strings.LastIndexByte(structureType, '}')
			if i < 0 {
				return nil, ErrInvalidStructure
			}

			// 寻找下一个 '{' 前有多少个 ',' 分隔符，以此来判断是否有下一级结构
			num := 0
			find := structureType[1:i]
			for _, c := range find {
				if c == ',' {
					num++
				} else if c == '{' {
					break
				}
			}
			kvs := strings.SplitN(structureType[1:i], ",", num+1)
			structure.StructFields = make([]*StructureStructField, 0, len(kvs))
			for _, kv := range kvs {
				kvp := strings.SplitN(kv, ":", 2)
				if len(kvp) != 2 {
					return nil, ErrInvalidStructure
				}

				key := strings.TrimSpace(kvp[0])
				value := strings.TrimSpace(kvp[1])
				valueStructure, err := ParseStructure(value)
				if err != nil {
					return nil, err
				}

				structure.StructFields = append(structure.StructFields, &StructureStructField{
					Name:      key,
					Structure: valueStructure,
				})
			}

			structure.Type = SourceFieldStructureTypeStruct
			next = structureType[i+1:]
		}

		if next != "" {
			nextStructure, err := ParseStructure(next)
			if err != nil {
				return nil, err
			}
			structure.Next = nextStructure
		}
	}
	return structure, err
}

// Structure 字段结构信息
type Structure struct {
	Raw          string                  // 原始结构
	Type         StructureType           // 字段类型
	Next         *Structure              // 下一级结构
	ArrayLength  int                     // 数组长度
	MapKey       *Structure              // Map 键
	MapValue     *Structure              // Map 值
	StructFields []*StructureStructField // 结构体字段
}

// StructureStructField 结构体字段
type StructureStructField struct {
	Name      string     // 字段名称
	Structure *Structure // 字段结构
}
