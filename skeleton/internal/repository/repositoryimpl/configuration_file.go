package repositoryimpl

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"github.com/kercylan98/minotaur/skeleton/internal/repository"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/spf13/viper"
	"strings"
)

var _ repository.ConfigurationRepository = (*ConfigurationFileRepository)(nil)

//go:embed "templates/application.template.yaml"
var applicationConfigTemplate []byte

func NewConfigurationFileRepository() repository.ConfigurationRepository {
	var r = &ConfigurationFileRepository{}
	return r
}

type ConfigurationFileRepository struct {
	applicationConfig *repository.Configuration
}

func (c *ConfigurationFileRepository) OnInitialize(ctx *application.Context) (err error) {
	if err = c.initApplicationConfig(ctx); err != nil {
		return
	}
	return
}

func (c *ConfigurationFileRepository) GetServiceConfig(ctx context.Context, dto *repository.GetServiceConfigDTO) (result *repository.GetServiceConfigResultDTO, err error) {
	service := viper.New()
	if err = service.MergeConfigMap(c.applicationConfig.GetStringMap(dto.ServiceName)); err != nil {
		return
	}
	return &repository.GetServiceConfigResultDTO{Config: &repository.Configuration{Viper: service}}, nil
}

func (c *ConfigurationFileRepository) LoadServiceConfig(ctx context.Context, dto *repository.LoadServiceConfigDTO) (err error) {
	var config *repository.GetServiceConfigResultDTO
	if config, err = c.GetServiceConfig(ctx, &repository.GetServiceConfigDTO{ServiceName: dto.ServiceName}); err != nil {
		return
	}
	return config.Config.Unmarshal(dto.Target)
}

func (c *ConfigurationFileRepository) initApplicationConfig(ctx *application.Context) (err error) {
	applicationViper := viper.New()
	applicationViper.SetConfigType("yaml")
	applicationViper.AddConfigPath(ctx.GetConfigDir())
	applicationViper.AddConfigPath(".")

	if ctx.IsTemplateConfig() {
		reader := bytes.NewReader(applicationConfigTemplate)
		if err = applicationViper.ReadConfig(reader); err != nil {
			return err
		}
	} else {
		env := fmt.Sprintf("application-%s.yaml", strings.ToLower(ctx.GetBootstrapConfig().GetString(application.ConfigurationEnvKey)))

		applicationViper.SetConfigFile(env)
		if err = applicationViper.ReadInConfig(); err != nil {
			applicationViper.SetConfigFile("application.yaml")
			if err = applicationViper.ReadInConfig(); err != nil {
				return
			} else {
				return
			}
		}
	}

	c.applicationConfig = &repository.Configuration{Viper: applicationViper}
	return
}
