package service

import (
	"context"
	"github.com/kercylan98/minotaur/skeleton/internal/repository"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
)

type ConfigurationService interface {
	application.Service

	GetBootstrapConfig(ctx context.Context) (result *GetBootstrapConfigResultDTO, err error)

	LoadBootstrapConfig(ctx context.Context, dto *LoadBootstrapConfigDTO) (err error)

	GetServiceConfig(ctx context.Context, dto *GetServiceConfigDTO) (result *GetServiceConfigResultDTO, err error)

	LoadServiceConfig(ctx context.Context, dto *LoadServiceConfigDTO) (err error)
}

type GetServiceConfigDTO struct {
	ServiceName string
}

type GetServiceConfigResultDTO struct {
	ServiceName string
	Config      *repository.Configuration
}

type LoadServiceConfigDTO struct {
	ServiceName string
	Target      any
}

type GetBootstrapConfigResultDTO struct {
	Config *repository.Configuration
}

type LoadBootstrapConfigDTO struct {
	Target any
}
