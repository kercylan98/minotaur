package configuration

const (
	SourceFieldTypeInvalid SourceFieldType = iota // 无效字段
	SourceFieldTypeServer                         // 服务器端字段
	SourceFieldTypeClient                         // 客户端字段
	SourceFieldTypeCommon                         // 通用字段
)

// SourceFieldType 代表了配置源字段的类型
type SourceFieldType = int8

// Source 配置源表示了游戏配置的来源，通常为 XLSX 文件等
type Source interface {
	// DisplayName 获取配置源适合人类阅读的名称
	DisplayName() string
	// Name 获取配置源所表示的配置名称
	Name() string
	// Description 获取配置源的描述，该描述通常用于注释生成等场景
	Description() string
	// Fields 获取配置源中的字段
	Fields() []*SourceField
	// Rows 获取配置源中的数据行
	Rows(fields []*SourceField) []SourceRow
}

// SourceField 配置源字段表示了配置源中的字段
type SourceField struct {
	DisplayName string          // 适合人类阅读的名称
	Name        string          // 字段名称
	Index       int             // 获取配置源该字段的索引值
	Structure   string          // 字段结构
	Type        SourceFieldType // 获取配置源字段的类型
}

// SourceRow 配置源行表示了配置源中的一组完整数据
type SourceRow []*SourceCell

// SourceCell 配置源单元表示了配置源中的一个单元格
type SourceCell struct {
	Field *SourceField // 单元格所属的字段
	Value string       // 单元格的值
}
