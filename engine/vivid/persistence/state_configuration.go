package persistence

func newStateConfiguration() *StateConfiguration {
	return &StateConfiguration{}
}

type StateConfiguration struct {
	storage Storage
}

func (c *StateConfiguration) WithStorage(storage Storage) *StateConfiguration {
	c.storage = storage
	return c
}
