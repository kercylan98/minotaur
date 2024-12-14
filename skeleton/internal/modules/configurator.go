package modules

import "github.com/kercylan98/minotaur/skeleton/internal/modules/configurator/configuratormodels"

type ConfiguratorModule interface {
	GetRuntimeConfig() configuratormodels.RuntimeConfig

	WatchRuntimeConfigChangedEvent(handler func(runtimeConfig configuratormodels.RuntimeConfig))
}
