package serviceimpl

import (
	"context"
	"github.com/kercylan98/minotaur/skeleton/internal/repository"
	"github.com/kercylan98/minotaur/skeleton/internal/service"
	application2 "github.com/kercylan98/minotaur/skeleton/pkg/application"
)

var _ service.ConfigurationService = (*ConfigurationService)(nil)

func NewConfigurationService() *ConfigurationService {
	var s = &ConfigurationService{}
	return s
}

type ConfigurationService struct {
	repository struct {
		configuration repository.ConfigurationRepository
	}
}

func (c *ConfigurationService) OnInitialize(app *application2.Context, loader *application2.RepositoryLoader) error {
	c.repository.configuration = application2.LoadRepository[repository.ConfigurationRepository](loader)
	return nil
}

func (c *ConfigurationService) GetServiceConfig(ctx context.Context, dto *service.GetServiceConfigDTO) (result *service.GetServiceConfigResultDTO, err error) {
	var config *repository.GetServiceConfigResultDTO
	if config, err = c.repository.configuration.GetServiceConfig(ctx, &repository.GetServiceConfigDTO{ServiceName: dto.ServiceName}); err != nil {
		return
	}
	return &service.GetServiceConfigResultDTO{Config: config.Config, ServiceName: dto.ServiceName}, nil
}

func (c *ConfigurationService) LoadServiceConfig(ctx context.Context, dto *service.LoadServiceConfigDTO) (err error) {
	return c.repository.configuration.LoadServiceConfig(ctx, &repository.LoadServiceConfigDTO{ServiceName: dto.ServiceName, Target: dto.Target})
}
