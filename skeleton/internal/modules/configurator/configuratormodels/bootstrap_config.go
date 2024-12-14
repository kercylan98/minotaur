package configuratormodels

import "github.com/kercylan98/minotaur/skeleton/internal/modules/configurator/configuratorcontants"

// BootstrapConfig 引导配置
type BootstrapConfig struct {
	App BootstrapAppConfig `yaml:"app" json:"app" mapstructure:"app" toml:"app"`
}

type BootstrapAppConfig struct {
	Env   configuratorcontants.Env `yaml:"env" json:"env" mapstructure:"env" toml:"env"` // 运行环境
	Actor BootstrapAppActorConfig  `yaml:"actor" json:"actor" mapstructure:"actor" toml:"actor"`
}

type BootstrapAppActorConfig struct {
	Name      string   `yaml:"name" json:"name" mapstructure:"name" toml:"name"`                     // Actor 名称
	Addr      string   `yaml:"addr" json:"addr" mapstructure:"addr" toml:"addr"`                     // Actor 地址
	SeedNodes []string `yaml:"seedNodes" json:"seedNodes" mapstructure:"seedNodes" toml:"seedNodes"` // 集群种子节点
}
