package service

import (
	"context"
	"github.com/kercylan98/minotaur/skeleton/internal/repository"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
)

type ConfigurationService interface {
	application.Service

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
