package configuration

// Analyzer 分析器接口，通过分析特定文件产生配置文件
type Analyzer interface {
	Analyze(filePath string) (map[string]Configuration, error)
}
