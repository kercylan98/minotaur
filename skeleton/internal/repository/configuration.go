package repository

import (
	"context"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/spf13/viper"
)

const (
	ConfigurationEnvKey         = "app.env"
	ConfigurationLoggerLevelKey = "app.logger.level"
)

type Configuration struct {
	*viper.Viper `mapstructure:"-" json:"-"`
}

type ConfigurationRepository interface {
	application.Repository

	// GetBootstrapConfig 获取引导配置
	GetBootstrapConfig(ctx context.Context) (result *GetBootstrapConfigResultDTO, err error)

	// LoadBootstrapConfig 加载引导配置到指定目标
	LoadBootstrapConfig(ctx context.Context, dto *LoadBootstrapConfigDTO) (err error)

	// GetServiceConfig 获取服务配置
	GetServiceConfig(ctx context.Context, dto *GetServiceConfigDTO) (result *GetServiceConfigResultDTO, err error)

	// LoadServiceConfig 加载服务配置到指定目标
	LoadServiceConfig(ctx context.Context, dto *LoadServiceConfigDTO) (err error)
}

type GetBootstrapConfigResultDTO struct {
	Config *Configuration
}

type GetServiceConfigDTO struct {
	ServiceName string
}

type GetServiceConfigResultDTO struct {
	Config *Configuration
}

type LoadBootstrapConfigDTO struct {
	Target any
}

type LoadServiceConfigDTO struct {
	ServiceName string
	Target      any
}
