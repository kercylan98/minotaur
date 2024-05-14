package vivid

// ClientFactory 是 ActorSystem 远程调用的客户端工厂
type ClientFactory func(network, host string, port uint16) (Client, error)

func (f ClientFactory) NewClient(network, host string, port uint16) (Client, error) {
	return f(network, host, port)
}

// Client 是 ActorSystem 远程调用的客户端
type Client interface {
	// Ask 执行远程调用，并返回结果
	Ask(data []byte) ([]byte, error)

	// Tell 执行远程调用，但不返回结果
	Tell(data []byte) error
}
