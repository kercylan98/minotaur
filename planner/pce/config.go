package pce

// Config 配置解析接口
//   - 用于将配置文件解析为可供分析的数据结构
//   - 可以在 cs 包中找到内置提供的实现及其模板，例如 cs.XlsxIndexConfig
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
	GetFields() []DataField
	// GetData 获取数据
	GetData() [][]DataInfo
}
