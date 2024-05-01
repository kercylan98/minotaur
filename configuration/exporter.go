package configuration

// Exporter 数据导出器
type Exporter interface {

	// Export 导出数据
	Export(rows []SourceRow) error
}
