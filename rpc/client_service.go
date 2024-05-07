package rpc

type ClientService struct {
	CallableService
}

func (c *ClientService) GetId() InstanceId {
	return c.GetServiceInfo().InstanceId
}

func (c *ClientService) GetWeight() int {
	return 1
}
