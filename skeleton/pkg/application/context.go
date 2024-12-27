package application

import (
	"context"
	"reflect"
)

func New(configurator ...Configurator) *Context {
	configuration := newConfiguration()
	for _, c := range configurator {
		c.Configure(configuration)
	}

	ctx := &Context{
		config: configuration,
		ctx:    context.Background(),
	}
	return ctx
}

type Context struct {
	config          *Configuration
	ctx             context.Context
	bootstrapConfig *BootstrapConfig
	modules         map[reflect.Type]Module
	services        map[reflect.Type][]Service
	repositoryList  map[reflect.Type][]Repository
	controllers     []Controller
}

func (c *Context) GetBootstrapConfig() *BootstrapConfig {
	return c.bootstrapConfig
}

func (c *Context) IsTemplateConfig() bool {
	return c.config.templateConfig
}

func (c *Context) GetConfigDir() string {
	return c.config.configDir
}

func (c *Context) Run() (err error) {

	if err = initBootstrapConfig(c, c.GetConfigDir(), "bootstrap.yaml"); err != nil {
		return
	}

	if err = runModules(c); err != nil {
		return
	}

	if err = runRepositoryList(c); err != nil {
		return
	}

	if err = runServices(c); err != nil {
		return
	}

	if err = runControllers(c); err != nil {
		return
	}

	if err = setupModules(c); err != nil {
		return
	}

	return nil
}
