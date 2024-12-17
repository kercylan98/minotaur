package configuratorcontants

type Env = string

const (
	BootstrapConfigName = "bootstrap"
	RuntimeConfigName   = "application"
)

const (
	EnvDevelopment Env = "dev"
	EnvTest        Env = "test"
	EnvProduction  Env = "prod"
)

var (
	SupportConfigFileTypes = []string{
		"yaml", "yml", "toml", "json",
	}
)
