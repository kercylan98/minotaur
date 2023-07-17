package pce

type Config interface {
	// GetConfigName 配置名称
	GetConfigName() string
	// GetDisplayName 配置显示名称
	GetDisplayName() string
	// GetDescription 配置描述
	GetDescription() string
	// GetIndexCount 索引数量
	GetIndexCount() int
	// GetFields 获取字段
	GetFields() []dataField
	// GetData 获取数据
	GetData() [][]dataInfo
}
