package configuration

type Analyzer interface {
	Analyze(filePath string) (map[string]Configuration, error)
}
