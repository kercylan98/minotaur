package transport

func newFiberConfiguration() *FiberConfiguration {
	return &FiberConfiguration{}
}

type FiberConfiguration struct {
	services []FiberService
}

func (c *FiberConfiguration) WithServices(services ...FiberService) *FiberConfiguration {
	c.services = append(c.services, services...)
	return c
}
