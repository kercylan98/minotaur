package rpc

// DiscoveryClientFactory 用于创建一个新的 RPC 客户端
type DiscoveryClientFactory func(address string) (Client, error)

// Discovery 服务发现接口
type Discovery interface {
	// GetInstance 用于获取一个服务实例
	GetInstance(name string) (Client, error)
}

// DiscoveryClient 用于创建一个新的基于服务发现的 RPC 客户端
type DiscoveryClient struct {
	discovery Discovery
}

// NewDiscoveryClient 用于创建一个新的基于服务发现的 RPC 客户端
func NewDiscoveryClient(d Discovery) *DiscoveryClient {
	return &DiscoveryClient{
		discovery: d,
	}
}

// Tell 用于向指定的服务发起一个 RPC 调用，该调用不需要返回值
func (c *DiscoveryClient) Tell(name string, route Route, data any) error {
	instance, err := c.discovery.GetInstance(name)
	if err != nil {
		return err
	}

	return instance.Tell(route, data)
}

// AsyncTell 用于向指定的服务发起一个异步 RPC 调用，该调用不需要返回值，当调用失败后会通过回调函数告知错误信息
func (c *DiscoveryClient) AsyncTell(name string, route Route, data any, callback ...func(err error)) {
	instance, err := c.discovery.GetInstance(name)
	if err != nil {
		if len(callback) > 0 {
			callback[0](err)
		}
		return
	}

	instance.AsyncTell(route, data, callback...)
}

// Ask 用于向指定的服务发起一个 RPC 调用，该调用需要返回值
func (c *DiscoveryClient) Ask(name string, route Route, data any) (Response, error) {
	instance, err := c.discovery.GetInstance(name)
	if err != nil {
		return nil, err
	}

	return instance.Ask(route, data)
}
