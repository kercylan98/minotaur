package datasheet

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/maths"
	lua "github.com/yuin/gopher-lua"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DTypePrefab   DType = "prefab"   // 预制件数据表
	DTypeStandard DType = "standard" // 标准数据表
	DTypeIndex    DType = "index"    // 索引数据表
)

const (
	BasicTypeBoolean  = "boolean"
	BasicTypeBool     = "bool"
	BasicTypeByte     = "byte"
	BasicTypeShort    = "short"
	BasicTypeLong     = "long"
	BasicTypeInt      = "int"
	BasicTypeInt8     = "int8"
	BasicTypeInt16    = "int16"
	BasicTypeInt32    = "int32"
	BasicTypeInt64    = "int64"
	BasicTypeUint     = "uint"
	BasicTypeUint8    = "uint8"
	BasicTypeUint16   = "uint16"
	BasicTypeUint32   = "uint32"
	BasicTypeUint64   = "uint64"
	BasicTypeFloat    = "float"
	BasicTypeDouble   = "double"
	BasicTypeNumber   = "number"
	BasicTypeFloat32  = "float32"
	BasicTypeFloat64  = "float64"
	BasicTypeTime     = "time"     // 含时间和日期
	BasicTypeDateTime = "datetime" // 含日期和时间
	BasicTypeDateOnly = "dateonly" // 仅有日期
	BasicTypeTimeOnly = "timeonly" // 仅有时间
	BasicTypeString   = "string"
	BasicTypeStr      = "str"
	BasicTypeBigInt   = "bigint"
	BasicTypeBigFloat = "bigfloat"
	BasicTypeDuration = "duration"
)

var (
	// nameRegexp 匹配名称的正则表达式
	//  - [a-zA-Z] 首字母必须为字母
	//  - [a-zA-Z0-9_]* 后续字符必须为字母、数字或下划线
	nameRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
)

var timeFormats = []string{
	time.DateTime,
	time.DateOnly,
	time.TimeOnly,
	time.RFC3339,
	time.RFC3339Nano,
	time.ANSIC,
	time.UnixDate,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RubyDate,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	time.Layout,
}

var basicType = []string{
	BasicTypeBoolean,
	BasicTypeBool,
	BasicTypeByte,
	BasicTypeShort,
	BasicTypeLong,
	BasicTypeInt,
	BasicTypeInt8,
	BasicTypeInt16,
	BasicTypeInt32,
	BasicTypeInt64,
	BasicTypeUint,
	BasicTypeUint8,
	BasicTypeUint16,
	BasicTypeUint32,
	BasicTypeUint64,
	BasicTypeFloat,
	BasicTypeDouble,
	BasicTypeFloat32,
	BasicTypeFloat64,
	BasicTypeNumber,
	BasicTypeTime,
	BasicTypeString,
	BasicTypeStr,
	BasicTypeBigInt,
	BasicTypeBigFloat,
	BasicTypeDateTime,
	BasicTypeDateOnly,
	BasicTypeTimeOnly,
	BasicTypeDuration,
}
var basicTypeMap = make(map[string]struct{})

func init() {
	for _, t := range basicType {
		basicTypeMap[t] = struct{}{}
	}
}

func IsBasicType(input string) bool {
	_, exist := basicTypeMap[input]
	return exist
}

type DType = string

func newSet() *Set {
	return &Set{
		Prefabs:    make(map[string]*Prefab),
		Datasheets: make(map[DType][]*Datasheet),
	}
}

type Set struct {
	Prefabs    map[string]*Prefab
	Datasheets map[DType][]*Datasheet
}

type Datasheet struct {
	Name        string   // 字段名称
	Description string   // 结构描述
	Filepath    string   // 文件路径
	Fields      []*Field // 字段列表
}

type Field struct {
	Owner       *Datasheet // 所属结构
	Name        string     // 字段名称
	Optional    bool       // 是否可选
	Type        Type       // 字段类型
	Description string     // 字段描述
	Index       int        // 字段索引
	Groups      []string   // 字段分组
	Values      []string   // 字段字符值
}

type Prefab struct {
	Type
	Name        string // 预制体名称
	Description string // 预制体描述
}

type StructField struct {
	Owner    *Struct // 所属结构
	Name     string  // 字段名称
	Optional bool    // 是否可选
	Type     Type    // 字段类型
}

type Struct struct {
	Name   string         // 结构名称
	Fields []*StructField // 字段列表
}

type Array struct {
	Length   int  // 数组长度
	Optional bool // 可选元素类型
	Type     Type // 数组元素类型
}

type Slice struct {
	Optional bool // 可选元素类型
	Type     Type // 切片元素类型
}

type Map struct {
	KeyType       Type // 键类型
	ValueType     Type // 值类型
	ValueOptional bool // 值可选
}

type Basic struct {
	Name string // 基本类型名称
}

type TodoPrefab struct {
	Optional bool   // 可选的
	Name     string // 预制名称
}

type Type interface {
	isTodoPrefab() bool
	parseData(luaState *lua.LState, value string) (any, error)
}

func (t *Struct) isTodoPrefab() bool { return false }

func (t *Struct) parseData(luaState *lua.LState, value string) (result any, err error) {
	var bytes []byte
	if isValidLuaTableOrArray(value) {
		bytes, err = lua2json(luaState, value)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	bytes = []byte(value)
	err = toolkit.UnmarshalJSONE(bytes, &result)
	return result, nil
}

func (t *Array) isTodoPrefab() bool { return false }

func (t *Array) parseData(luaState *lua.LState, value string) (result any, err error) {
	var bytes []byte
	if isValidLuaTableOrArray(value) {
		bytes, err = lua2json(luaState, value)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	bytes = []byte(value)
	err = toolkit.UnmarshalJSONE(bytes, &result)
	return result, nil
}

func (t *Slice) isTodoPrefab() bool { return false }

func (t *Slice) parseData(luaState *lua.LState, value string) (result any, err error) {
	var bytes []byte
	if isValidLuaTableOrArray(value) {
		bytes, err = lua2json(luaState, value)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	bytes = []byte(value)
	err = toolkit.UnmarshalJSONE(bytes, &result)
	return result, nil
}

func (t *Map) isTodoPrefab() bool { return false }

func (t *Map) parseData(luaState *lua.LState, value string) (result any, err error) {
	var bytes []byte
	if isValidLuaTableOrArray(value) {
		bytes, err = lua2json(luaState, value)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	bytes = []byte(value)
	err = toolkit.UnmarshalJSONE(bytes, &result)
	return result, nil
}

func (t *Basic) isTodoPrefab() bool { return false }

func (t *Basic) parseData(luaState *lua.LState, value string) (result any, err error) {
	switch t.Name {
	case BasicTypeBoolean, BasicTypeBool:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = false
			return
		}
		result, err = strconv.ParseBool(tsv)
	case BasicTypeByte:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = byte(0)
			return
		}
		result, err = strconv.ParseUint(tsv, 10, 8)
		if err != nil {
			return nil, err
		}
		result = byte(result.(uint64))
	case BasicTypeShort, BasicTypeInt16:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = int16(0)
			return
		}
		result, err = strconv.ParseInt(tsv, 10, 16)
		if err != nil {
			return nil, err
		}
		result = int16(result.(int64))
	case BasicTypeLong, BasicTypeInt64:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = int64(0)
			return
		}
		result, err = strconv.ParseInt(tsv, 10, 64)
		if err != nil {
			return nil, err
		}
		result = result.(int64)
	case BasicTypeInt:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = 0
			return
		}
		result, err = strconv.Atoi(tsv)
	case BasicTypeInt8:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = int8(0)
			return
		}
		result, err = strconv.ParseInt(tsv, 10, 8)
		if err != nil {
			return nil, err
		}
		result = int8(result.(int64))
	case BasicTypeInt32:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = int32(0)
			return
		}
		result, err = strconv.ParseInt(tsv, 10, 32)
		if err != nil {
			return nil, err
		}
		result = int32(result.(int64))
	case BasicTypeUint:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = uint(0)
			return
		}
		result, err = strconv.ParseUint(tsv, 10, 0)
		if err != nil {
			return nil, err
		}
		result = uint(result.(uint64))
	case BasicTypeUint8:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = uint8(0)
			return
		}
		result, err = strconv.ParseUint(tsv, 10, 8)
		if err != nil {
			return nil, err
		}
		result = uint8(result.(uint64))
	case BasicTypeUint16:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = uint16(0)
			return
		}
		result, err = strconv.ParseUint(tsv, 10, 16)
		if err != nil {
			return nil, err
		}
		result = uint16(result.(uint64))
	case BasicTypeUint32:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = uint32(0)
			return
		}
		result, err = strconv.ParseUint(tsv, 10, 32)
		if err != nil {
			return nil, err
		}
		result = uint32(result.(uint64))
	case BasicTypeUint64:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = uint64(0)
			return
		}
		result, err = strconv.ParseUint(tsv, 10, 64)
		if err != nil {
			return nil, err
		}
		result = result.(uint64)
	case BasicTypeFloat, BasicTypeFloat32:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = float32(0)
			return
		}
		result, err = strconv.ParseFloat(tsv, 32)
		if err != nil {
			return nil, err
		}
		result = float32(result.(float64))
	case BasicTypeDouble, BasicTypeFloat64, BasicTypeNumber:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = float64(0)
			return
		}
		result, err = strconv.ParseFloat(tsv, 64)
	case BasicTypeTime:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = time.Time{}
			return
		}
		// 遍历所有格式进行解析
		for _, format := range timeFormats {
			result, err = time.Parse(format, tsv)
			if err == nil {
				return
			}
		}
	case BasicTypeDateTime:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = time.Time{}
			return
		}
		result, err = time.Parse(time.DateTime, tsv)
	case BasicTypeTimeOnly:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = time.Time{}
			return
		}
		result, err = time.Parse(time.TimeOnly, tsv)
	case BasicTypeString, BasicTypeStr:
		// 去除首尾 '"'
		if len(value) > 2 && value[0] == '"' && value[len(value)-1] == '"' {
			value = value[1 : len(value)-1]
		}
		result = value
	case BasicTypeBigInt:
		var bi = new(BigInt)
		if err = bi.UnmarshalJSON([]byte(value)); err != nil {
			return nil, err
		}
		result = bi
	case BasicTypeBigFloat:
		var bf = new(BigFloat)
		if err = bf.UnmarshalJSON([]byte(value)); err != nil {
			return nil, err
		}
		result = bf
	case BasicTypeDuration:
		tsv := strings.TrimSpace(value)
		if tsv == "" {
			result = time.Duration(0)
			return
		}
		// 检查是否是纯数字
		if maths.IsPureNumber(tsv) {
			result, err = strconv.ParseInt(tsv, 10, 64)
			if err != nil {
				return nil, err
			}
			result = time.Duration(result.(int64))
			return
		}
		result, err = time.ParseDuration(tsv)
	default:
		return nil, fmt.Errorf("unknown basic type %s", t.Name)
	}

	return
}

func (t *TodoPrefab) isTodoPrefab() bool { return true }

func (t *TodoPrefab) parseData(luaState *lua.LState, value string) (result any, err error) {
	return nil, fmt.Errorf("prefab can't parse data")
}

func (d *Datasheet) HasIndex() bool {
	for _, field := range d.Fields {
		if field.Index == 1 {
			return true
		}
	}
	return false
}

// LoadData 加载数据
func (s *Set) LoadData() (map[string]any, error) {
	state := lua.NewState()
	defer state.Close()
	datasheetDataMap := make(map[string]any)
	for typ, datasheets := range s.Datasheets {
		for _, datasheet := range datasheets {
			switch typ {
			case DTypeStandard:
				if v, err := loadDataFromStandardDatasheet(state, datasheet); err != nil {
					return nil, err
				} else {
					datasheetDataMap[datasheet.Name] = v
				}
			case DTypeIndex:
				if v, err := loadDataFromIndexDatasheet(state, datasheet); err != nil {
					return nil, err
				} else {
					datasheetDataMap[datasheet.Name] = v
				}
			default:
				return nil, fmt.Errorf("load datasheet %s(%s) data failed, %w", datasheet.Name, datasheet.Filepath, ErrorNotSupportDatasheetType)
			}
		}
	}
	return datasheetDataMap, nil
}

func loadDataFromIndexDatasheet(state *lua.LState, datasheet *Datasheet) (map[string]any, error) {
	var maxRow int
	for _, field := range datasheet.Fields {
		if len(field.Values) > maxRow {
			maxRow = len(field.Values)
		}
	}

	var rawDatasheetData = make(map[string]any)
	var curr = rawDatasheetData
	for i := 0; i < maxRow; i++ {
		var row = make(map[string]any)
		var indexValue = make(map[int]string)
		for _, field := range datasheet.Fields {
			if value, err := field.Type.parseData(state, field.Values[i]); err != nil {
				return nil, err
			} else {
				row[field.Name] = value
				indexValue[field.Index] = fmt.Sprint(value)
			}
		}

		for v := 0; v < len(datasheet.Fields); v++ {
			field := datasheet.Fields[v]
			if field.Index == 0 {
				break
			}
			value := indexValue[field.Index]

			temp, exist := curr[value]
			if !exist {
				temp = make(map[string]any)
				if nextV := v + 1; nextV < len(datasheet.Fields) && datasheet.Fields[nextV].Index == 0 {
					curr[value] = row
				} else {
					curr[value] = temp
				}
				curr = temp.(map[string]any)
			} else {
				curr = temp.(map[string]any)
			}
		}
		curr = rawDatasheetData
	}

	return rawDatasheetData, nil
}

func loadDataFromStandardDatasheet(state *lua.LState, datasheet *Datasheet) (map[any]any, error) {
	var datasheetData = make(map[any]any)
	for _, field := range datasheet.Fields {
		if value, err := field.Type.parseData(state, field.Values[0]); err != nil {
			return nil, err
		} else {
			datasheetData[field.Name] = value
		}
	}
	return datasheetData, nil
}
