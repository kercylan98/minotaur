package transport

func newFiberConfiguration() *FiberConfiguration {
	return &FiberConfiguration{}
}

type FiberConfiguration struct {
	services []FiberService
}

func (c *FiberConfiguration) WithServices(service ...FiberService) *FiberConfiguration {
	c.services = append(c.services, service...)
	return c
}
