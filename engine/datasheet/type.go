package datasheet

import "regexp"

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
	BasicTypeBigInt,
	BasicTypeBigFloat,
	BasicTypeDateTime,
	BasicTypeDateOnly,
	BasicTypeTimeOnly,
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
}

func (t *Struct) isTodoPrefab() bool     { return false }
func (t *Array) isTodoPrefab() bool      { return false }
func (t *Slice) isTodoPrefab() bool      { return false }
func (t *Map) isTodoPrefab() bool        { return false }
func (t *Basic) isTodoPrefab() bool      { return false }
func (t *TodoPrefab) isTodoPrefab() bool { return true }

func (d *Datasheet) HasIndex() bool {
	for _, field := range d.Fields {
		if field.Index == 1 {
			return true
		}
	}
	return false
}
