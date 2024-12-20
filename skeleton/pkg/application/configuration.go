package application

func newConfiguration() *Configuration {
	return &Configuration{
		configDir: ".",
	}
}

type Configuration struct {
	configDir      string // 配置文件目录
	templateConfig bool   // 是否使用模板配置
}

func (c *Configuration) WithConfigDir(configDir string) {
	c.configDir = configDir
}

func (c *Configuration) WithTemplateConfig(enable bool) {
	c.templateConfig = enable
}
