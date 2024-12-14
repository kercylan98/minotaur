package configurator

import (
	"github.com/kercylan98/minotaur/skeleton/internal/modules"
	"github.com/kercylan98/minotaur/skeleton/internal/modules/configurator/configuratormodels"
)

var _ modules.ConfiguratorModule = (*configuratorServices)(nil)

func newConfiguratorServices(module *configuratorModule) *configuratorServices {
	return &configuratorServices{module: module}
}

type configuratorServices struct {
	module                           *configuratorModule
	runtimeConfigChangedEventHandles []func(runtimeConfig configuratormodels.RuntimeConfig)
}

func (c *configuratorServices) GetRuntimeConfig() configuratormodels.RuntimeConfig {
	c.module.runtimeRWMutex.RLock()
	defer c.module.runtimeRWMutex.RUnlock()
	return *c.module.runtimeConfig
}

func (c *configuratorModule) WatchRuntimeConfigChangedEvent(handler func(runtimeConfig configuratormodels.RuntimeConfig)) {
	c.services.runtimeConfigChangedEventHandles = append(c.services.runtimeConfigChangedEventHandles, handler)
}
