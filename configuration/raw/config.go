package raw

// NewConfig 创建一个配置
func NewConfig(name, desc string) Config {
	var c = Config{
		name:            name,
		description:     desc,
		fields:          map[string]Field{},
		fieldStructures: map[string][]Structure{},
	}
	return c
}

type Config struct {
	name            string                 // 配置名称
	description     string                 // 配置描述
	fields          map[string]Field       // 原始字段描述
	fieldStructures map[string][]Structure // 原始字段描述中包含的结构体
}

// GetName 获取配置名称
func (c *Config) GetName() string {
	return c.name
}

// GetDescription 获取配置描述
func (c *Config) GetDescription() string {
	return c.description
}

// GetFields 获取原始字段描述
func (c *Config) GetFields() map[string]Field {
	return c.fields
}

// GetFieldsWithSort 获取原始字段描述并按照索引排序
func (c *Config) GetFieldsWithSort() []Field {
	var fields = make([]Field, len(c.fields))
	for _, f := range c.fields {
		fields[f.index] = f
	}
	return fields
}

// GetField 获取指定字段描述
func (c *Config) GetField(name string) Field {
	return c.fields[name]
}

// GetFieldNum 获取原始字段数量
func (c *Config) GetFieldNum() int {
	return len(c.fields)
}

// GetFieldStructures 获取原始字段描述中包含的结构体
func (c *Config) GetFieldStructures() map[string][]Structure {
	return c.fieldStructures
}

// GetFieldStructure 获取指定字段的结构体
func (c *Config) GetFieldStructure(name string) []Structure {
	return c.fieldStructures[name]
}
