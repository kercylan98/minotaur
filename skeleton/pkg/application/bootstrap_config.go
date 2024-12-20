package application

import (
	"bytes"
	_ "embed"
	"github.com/kercylan98/minotaur/skeleton/pkg/utils/configutils"
	"github.com/spf13/viper"
)

const (
	ConfigurationEnvKey         = "app.env"
	ConfigurationLoggerLevelKey = "app.logger.level"
)

//go:embed "bootstrap.template.yaml"
var bootstrapConfigTemplate []byte

type BootstrapConfig struct {
	*viper.Viper `mapstructure:"-" json:"-"`
}

func initBootstrapConfig(ctx *Context, configFileDir, bootstrapConfigFileName string) (err error) {
	bootstrapViper := viper.New()
	bootstrapViper.SetConfigFile(bootstrapConfigFileName)
	bootstrapViper.SetConfigType("yaml")
	bootstrapViper.AddConfigPath(configFileDir)
	bootstrapViper.AddConfigPath(".")

	bootstrapViper.SetDefault(ConfigurationEnvKey, "dev")
	bootstrapViper.SetDefault(ConfigurationLoggerLevelKey, "info")

	if ctx.IsTemplateConfig() {
		bootstrapViper.SetConfigType("yaml")
		reader := bytes.NewReader(bootstrapConfigTemplate)
		if err = bootstrapViper.ReadConfig(reader); err != nil {
			return err
		}
	} else {
		if err = bootstrapViper.ReadInConfig(); err != nil {
			return
		}

		if err = configutils.CheckEnv(bootstrapViper.GetString(ConfigurationEnvKey)); err != nil {
			return
		}

		if err = configutils.CheckLoggerLevel(bootstrapViper.GetString(ConfigurationLoggerLevelKey)); err != nil {
			return
		}
	}
	ctx.bootstrapConfig = &BootstrapConfig{Viper: bootstrapViper}
	return
}
