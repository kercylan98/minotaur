package vivid

// ClientFactory 是 ActorSystem 远程调用的客户端工厂
type ClientFactory func(system *ActorSystem, network, host string, port uint16) (Client, error)

func (f ClientFactory) NewClient(system *ActorSystem, network, host string, port uint16) (Client, error) {
	return f(system, network, host, port)
}

// Client 是 ActorSystem 远程调用的客户端
type Client interface {
	// Exec 执行远程调用
	Exec(data []byte) error
}
