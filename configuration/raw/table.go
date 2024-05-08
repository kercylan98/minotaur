package raw

// NewTable 创建一个表
func NewTable(configs ...Config) Table {
	t := Table{
		configs: make(map[string]Config),
	}

	for _, config := range configs {
		t.configs[config.GetName()] = config
	}

	return t
}

type Table struct {
	configs map[string]Config
}

// GetConfig 获取配置
func (t *Table) GetConfig(name string) Config {
	return t.configs[name]
}

// GetConfigs 获取所有配置
func (t *Table) GetConfigs() map[string]Config {
	return t.configs
}
