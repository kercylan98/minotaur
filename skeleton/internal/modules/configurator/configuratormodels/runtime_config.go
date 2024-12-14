package configuratormodels

import "fmt"

// RuntimeConfig 运行时配置
type RuntimeConfig struct {
	Database RuntimeDatabaseConfig `yaml:"database" json:"database" mapstructure:"database" toml:"database"`
}

type RuntimeDatabaseConfig struct {
	MySQL RuntimeDatabaseMySQLConfig `yaml:"mysql" json:"mysql" mapstructure:"mysql" toml:"mysql"`
}

type RuntimeDatabaseMySQLConfig struct {
	Enabled  bool   `yaml:"enabled" json:"enabled" mapstructure:"enabled" toml:"enabled"`
	Host     string `yaml:"host" json:"host" mapstructure:"host" toml:"host"`
	Port     int    `yaml:"port" json:"port" mapstructure:"port" toml:"port"`
	Username string `yaml:"username" json:"username" mapstructure:"username" toml:"username"`
	Password string `yaml:"password" json:"password" mapstructure:"password" toml:"password"`
	Database string `yaml:"database" json:"database" mapstructure:"database" toml:"database"`
}

func (c *RuntimeDatabaseMySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database)
}
