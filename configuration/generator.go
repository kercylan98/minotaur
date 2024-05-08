package configuration

import "github.com/kercylan98/minotaur/configuration/raw"

// Generator 代码生成器
type Generator interface {

	// Generate 生成代码
	Generate(table raw.Table) error
}

// GenerateCode 生成代码
func GenerateCode(generator Generator, scanners ...Scanner) error {
	var configs []raw.Config
	for _, scanner := range scanners {
		config, err := scanner.StructScan()
		if err != nil {
			return err
		}
		configs = append(configs, config)
	}

	table := raw.NewTable(configs...)
	return generator.Generate(table)
}
