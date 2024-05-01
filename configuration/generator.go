package configuration

// Generator 代码生成器
type Generator interface {

	// Generate 生成代码
	Generate(fields []*Field) error
}
