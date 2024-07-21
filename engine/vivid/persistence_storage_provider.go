package vivid

import (
	"github.com/kercylan98/minotaur/engine/vivid/persistence"
	"github.com/kercylan98/minotaur/toolkit"
)

var defaultPersistenceStorageProviderInstance = toolkit.NewInertiaSingleton(func() persistence.StorageProvider {
	return new(defaultPersistenceStorageProvider)
})

func GetDefaultPersistenceStorageProvider() persistence.StorageProvider {
	return defaultPersistenceStorageProviderInstance.Get()
}

type defaultPersistenceStorageProvider struct{}

func (d *defaultPersistenceStorageProvider) GetStorageProviderName() persistence.StorageProviderName {
	return "__default"
}

func (d *defaultPersistenceStorageProvider) Provide() persistence.Storage {
	return persistence.NewMemoryStorage()
}
