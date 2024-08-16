package table

type Table interface {
	// GetName 是用于代码生成的配置文件结构名称
	GetName() string

	// GetDescribe 用于生成的结构注释
	GetDescribe() string

	// GetIndex 获取索引数量
	GetIndex() int

	// GetFields 获取配置结构的字段
	GetFields() FieldScanner

	// IsIgnore 是否忽略该表
	IsIgnore() bool
}
