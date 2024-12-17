package main

import (
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/internal/controller"
	"github.com/kercylan98/minotaur/skeleton/internal/repository"
	"github.com/kercylan98/minotaur/skeleton/internal/repository/repositoryimpl"
	"github.com/kercylan98/minotaur/skeleton/internal/service"
	"github.com/kercylan98/minotaur/skeleton/internal/service/serviceimpl"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"time"
)

type MyService struct {
	services struct {
		configuration service.ConfigurationService
	}
}

func (m *MyService) OnInitialize(app *application.Context, loader *application.RepositoryLoader) error {
	return nil
}

func (m *MyService) OnDependencyInitialize(loader *application.ServiceLoader) error {
	m.services.configuration = application.LoadService[service.ConfigurationService](loader)

	return nil
}

func (m *MyService) OnRun() {
	result, err := m.services.configuration.GetServiceConfig(context.Background(), &service.GetServiceConfigDTO{ServiceName: "example"})
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Config.GetString("value"))
}

func main() {
	ctx := application.New()

	application.RegisterRepository[repository.ConfigurationRepository](ctx, repositoryimpl.NewConfigurationFileWithTemplate())

	application.RegisterService[application.Service](ctx, new(MyService))
	application.RegisterService[service.ConfigurationService](ctx, serviceimpl.NewConfigurationService())
	application.RegisterController(ctx, controller.NewWebSocketController(vivid.NewActorSystem()))

	if err := ctx.Run(); err != nil {
		panic(err)
	}

	time.Sleep(chrono.Day)
}
