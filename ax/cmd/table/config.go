package table

type Config struct {
	name      string
	desc      string
	types     []Type
	fieldDesc map[string]string
	indexes   []Type
}

func (c *Config) GetName() string {
	return c.name
}

func (c *Config) GetDesc() string {
	return c.desc
}

func (c *Config) GetTypes() []Type {
	return c.types
}

func (c *Config) GetFieldDesc(field string) string {
	return c.fieldDesc[field]
}

func (c *Config) GetIndexes() []Type {
	return c.indexes
}
