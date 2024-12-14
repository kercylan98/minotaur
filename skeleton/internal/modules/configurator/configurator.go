package configurator

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster"
	"github.com/kercylan98/minotaur/skeleton/internal/modules/configurator/configuratorcontants"
	"github.com/kercylan98/minotaur/skeleton/internal/modules/configurator/configuratormodels"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/configutils"
	"github.com/kercylan98/minotaur/skeleton/pkg/module"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/spf13/viper"
	"sync"
)

//go:embed "bootstrap.template.yaml"
var bootstrapConfigTemplate []byte

//go:embed "application.template.yaml"
var runtimeConfigTemplate []byte

func NewConfiguratorModule(binder application.ActorSystemBinder) module.Module[*application.Context] {
	return &configuratorModule{
		binder: binder,
	}
}

type configuratorModule struct {
	app             *application.Context
	services        *configuratorServices
	binder          application.ActorSystemBinder
	bootstrapViper  *viper.Viper
	bootstrapConfig *configuratormodels.BootstrapConfig
	runtimeViper    *viper.Viper
	runtimeConfig   *configuratormodels.RuntimeConfig
	runtimeRWMutex  sync.RWMutex
}

func (c *configuratorModule) OnInitialize(app *application.Context) error {
	c.app = app
	c.services = newConfiguratorServices(c)
	if err := c.initBootstrapConfig(); err != nil {
		return err
	}

	if err := c.initRuntimeConfig(); err != nil {
		return err
	}
	return nil
}

func (c *configuratorModule) initBootstrapConfig() error {
	bootstrapViper := viper.New()
	bootstrapConfig := &configuratormodels.BootstrapConfig{}
	var withDefault bool
	if err := configutils.LoadConfigWithViper(bootstrapViper, configuratorcontants.BootstrapConfigName, bootstrapConfig, configuratorcontants.SupportConfigFileTypes...); err != nil {
		bootstrapViper.SetConfigType("yaml")
		reader := bytes.NewReader(bootstrapConfigTemplate)
		if err = bootstrapViper.ReadConfig(reader); err != nil {
			return err
		}
		if err = bootstrapViper.Unmarshal(bootstrapConfig); err != nil {
			return err
		}
		withDefault = true
	}
	c.bootstrapViper = bootstrapViper
	c.bootstrapConfig = bootstrapConfig
	c.onBindActorSystem()
	if withDefault {
		c.app.ActorSystem().Logger().Warn("ConfiguratorModule", log.String("status", "bootstrap config not found, use default config"), log.Any("config", bootstrapConfig))
	}
	return nil
}

func (c *configuratorModule) initRuntimeConfig() error {
	runtimeViper := viper.New()
	runtimeConfig := &configuratormodels.RuntimeConfig{}
	runtimeFileNames := []string{fmt.Sprintf("%s.%s", configuratorcontants.RuntimeConfigName, c.bootstrapConfig.App.Env), configuratorcontants.RuntimeConfigName}
	found := false
	for _, fileName := range runtimeFileNames {
		if err := configutils.LoadConfigWithViper(runtimeViper, fileName, runtimeConfig, configuratorcontants.SupportConfigFileTypes...); err == nil {
			found = true
			break
		}
	}
	if !found {
		runtimeViper.SetConfigType("yaml")
		reader := bytes.NewReader(runtimeConfigTemplate)
		if err := runtimeViper.ReadConfig(reader); err != nil {
			return err
		}
		if err := runtimeViper.Unmarshal(runtimeConfig); err != nil {
			return err
		}
		c.app.ActorSystem().Logger().Warn("ConfiguratorModule", log.String("status", "runtime config not found, use default config"), log.Any("config", runtimeConfig))
	}

	c.runtimeViper = runtimeViper
	c.runtimeConfig = runtimeConfig
	if found {
		c.runtimeViper.OnConfigChange(func(in fsnotify.Event) {
			c.runtimeRWMutex.Lock()
			defer c.runtimeRWMutex.Unlock()
			if err := c.runtimeViper.ReadInConfig(); err != nil {
				c.app.ActorSystem().Logger().Error("ConfiguratorModule", log.String("status", "reload runtime config failed"), log.Err(err))
			} else {
				c.app.ActorSystem().Logger().Info("ConfiguratorModule", log.String("status", "reload runtime config success"), log.Any("config", runtimeConfig))

				for _, handle := range c.services.runtimeConfigChangedEventHandles {
					func() {
						defer func() {
							if r := recover(); r != nil {
								c.app.ActorSystem().Logger().Error("ConfiguratorModule", log.String("status", "reload runtime config panic"), log.Any("panic", r), log.Any("handler", fmt.Sprintf("%T", handle)))
							}
						}()
						handle(*c.runtimeConfig)
					}()
				}
			}
		})
		c.runtimeViper.WatchConfig()
	}
	return nil
}

func (c *configuratorModule) onBindActorSystem() {
	actorConfig := c.bootstrapConfig.App.Actor
	if len(actorConfig.SeedNodes) == 0 {
		var configurator = vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
			if actorConfig.Name != "" {
				config.WithName(actorConfig.Name)
			}
			if actorConfig.Addr != "" {
				config.WithShared(actorConfig.Addr)
			}
		})

		c.binder(vivid.NewActorSystem(configurator), nil)
	} else {
		var configurator = cluster.FunctionalActorSystemConfigurator(func(config *cluster.ActorSystemConfiguration) {
			if actorConfig.Name != "" {
				config.WithName(actorConfig.Name)
			}

		})
		actorSystemCluster := cluster.NewActorSystem(actorConfig.Addr, actorConfig.SeedNodes, configurator)
		c.binder(actorSystemCluster.ActorSystem, actorSystemCluster)
	}
}
