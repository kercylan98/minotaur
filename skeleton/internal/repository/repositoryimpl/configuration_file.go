package repositoryimpl

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"github.com/kercylan98/minotaur/skeleton/internal/repository"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/utils/configutils"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var _ repository.ConfigurationRepository = (*ConfigurationFileRepository)(nil)

//go:embed "templates/bootstrap.template.yaml"
var bootstrapConfigTemplate []byte

//go:embed "templates/application.template.yaml"
var applicationConfigTemplate []byte

func NewConfigurationFileWithTemplate() *ConfigurationFileRepository {
	var r = &ConfigurationFileRepository{
		template: true,
	}
	return r
}

func NewConfigurationFileRepository(configFileDir, bootstrapConfigFileName, applicationConfigFileName string) repository.ConfigurationRepository {
	var r = &ConfigurationFileRepository{
		configFileDir:             configFileDir,
		bootstrapConfigFileName:   bootstrapConfigFileName,
		applicationConfigFileName: applicationConfigFileName,
	}
	return r
}

type ConfigurationFileRepository struct {
	configFileDir             string
	bootstrapConfigFileName   string
	applicationConfigFileName string
	template                  bool

	bootstrapConfig   *repository.Configuration
	applicationConfig *repository.Configuration
}

func (c *ConfigurationFileRepository) OnInitialize(ctx *application.Context) (err error) {
	if err = c.initBootstrapConfig(); err != nil {
		return
	}

	if err = c.initApplicationConfig(); err != nil {
		return
	}

	return
}

func (c *ConfigurationFileRepository) GetBootstrapConfig(ctx context.Context) (result *repository.GetBootstrapConfigResultDTO, err error) {
	return &repository.GetBootstrapConfigResultDTO{Config: c.bootstrapConfig}, nil
}

func (c *ConfigurationFileRepository) LoadBootstrapConfig(ctx context.Context, dto *repository.LoadBootstrapConfigDTO) (err error) {
	return c.bootstrapConfig.Unmarshal(dto.Target)
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

func (c *ConfigurationFileRepository) initBootstrapConfig() (err error) {
	bootstrapViper := viper.New()
	bootstrapViper.SetConfigFile(c.bootstrapConfigFileName)
	bootstrapViper.SetConfigType("yaml")
	bootstrapViper.AddConfigPath(c.configFileDir)
	bootstrapViper.AddConfigPath(".")

	bootstrapViper.SetDefault(repository.ConfigurationEnvKey, "dev")
	bootstrapViper.SetDefault(repository.ConfigurationLoggerLevelKey, "info")

	if c.template {
		bootstrapViper.SetConfigType("yaml")
		reader := bytes.NewReader(bootstrapConfigTemplate)
		if err = bootstrapViper.ReadConfig(reader); err != nil {
			return err
		}
	} else {
		if err = bootstrapViper.ReadInConfig(); err != nil {
			return
		}

		if err = configutils.CheckEnv(c.bootstrapConfig.GetString(repository.ConfigurationEnvKey)); err != nil {
			return
		}

		if err = configutils.CheckLoggerLevel(c.bootstrapConfig.GetString(repository.ConfigurationLoggerLevelKey)); err != nil {
			return
		}
	}

	c.bootstrapConfig = &repository.Configuration{Viper: bootstrapViper}
	return
}

func (c *ConfigurationFileRepository) initApplicationConfig() (err error) {
	if c.bootstrapConfig == nil {
		return
	}

	applicationViper := viper.New()
	applicationViper.SetConfigType("yaml")
	applicationViper.AddConfigPath(c.configFileDir)
	applicationViper.AddConfigPath(".")

	if c.template {
		reader := bytes.NewReader(applicationConfigTemplate)
		if err = applicationViper.ReadConfig(reader); err != nil {
			return err
		}
	} else {
		fileName := filepath.Base(c.applicationConfigFileName)
		fileExt := filepath.Ext(fileName)
		env := fmt.Sprintf("%s-%s%s", fileName, strings.ToLower(c.bootstrapConfig.GetString(repository.ConfigurationEnvKey)), fileExt)

		applicationViper.SetConfigFile(env)
		if err = applicationViper.ReadInConfig(); err != nil {
			applicationViper.SetConfigFile(c.applicationConfigFileName)
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
