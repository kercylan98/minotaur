package persistence

type StorageProvider interface {
	Provide() Storage
}

type FunctionalStorageProvider func() Storage

func (f FunctionalStorageProvider) Provide() Storage {
	return f()
}
