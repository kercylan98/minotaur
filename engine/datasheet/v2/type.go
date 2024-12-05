package datasheet

import "regexp"

const (
	DTypePrefab   DType = "prefab"   // 预制件数据表
	DTypeStandard DType = "standard" // 标准数据表
	DTypeIndex    DType = "index"    // 索引数据表
)

var (
	// nameRegexp 匹配名称的正则表达式
	//  - [a-zA-Z] 首字母必须为字母
	//  - [a-zA-Z0-9_]* 后续字符必须为字母、数字或下划线
	nameRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
)

var basicType = []string{
	"boolean", "bool", "int", "int8", "int16", "int32", "int64", "uint",
	"uint8", "uint16", "uint32", "uint64", "float", "double", "byte", "time",
	"float32", "float64", "short", "long", "string", "bigint", "bigfloat",
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
