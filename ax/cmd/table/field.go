package table

type Field interface {
	// GetIndex 获取字段索引值
	//  - 当返回值 <= 0 时为非索引字段
	GetIndex() int

	// GetName 字段名称
	GetName() string

	// GetDesc 字段描述
	GetDesc() string

	// GetType 字段类型
	GetType() string

	// GetParam 字段参数
	GetParam() string

	// IsIgnore 是否忽略
	IsIgnore() bool

	// Query 数据查询
	Query(pos int) (val map[string]any, skip, end bool)
}
