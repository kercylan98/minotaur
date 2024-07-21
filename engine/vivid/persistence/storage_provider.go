package persistence

type StorageProviderName = string

type StorageProvider interface {
	GetStorageProviderName() StorageProviderName

	Provide() Storage
}

type FunctionalStorageProvider func() Storage

func (f FunctionalStorageProvider) GetStorageProviderName() StorageProviderName {
	return ""
}

func (f FunctionalStorageProvider) Provide() Storage {
	return f()
}
