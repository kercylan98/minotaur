package vivid

// NewActorSystemOptions 用于创建一个 ActorSystemOptions
func NewActorSystemOptions() *ActorSystemOptions {
	return &ActorSystemOptions{
		Network: "tcp",
	}
}

// ActorSystemOptions 是 ActorSystem 的配置选项
type ActorSystemOptions struct {
	ClusterName   string        // 集群名称
	Network       string        // 网络协议
	Host          string        // 主机地址
	Port          uint16        // 端口
	Server        Server        // 远程调用服务端
	ClientFactory ClientFactory // 远程调用客户端
}

// Apply 用于应用配置选项
func (o *ActorSystemOptions) Apply(options ...*ActorSystemOptions) *ActorSystemOptions {
	for _, option := range options {
		if option.ClusterName != "" {
			o.ClusterName = option.ClusterName
		}
		if option.Network != "" {
			o.Network = option.Network
		}
		if option.Host != "" {
			o.Host = option.Host
		}
		if option.Port != 0 {
			o.Port = option.Port
		}
		if option.Server != nil {
			o.Server = option.Server
		}
		if option.ClientFactory != nil {
			o.ClientFactory = option.ClientFactory
		}
	}
	return o
}

// WithNetwork 设置网络协议
func (o *ActorSystemOptions) WithNetwork(network string) *ActorSystemOptions {
	o.Network = network
	return o
}

// WithCluster 设置集群信息，该可选项用于 ActorSystem 的集群配置
//   - 该选项将会覆盖 WithAddress 的配置
func (o *ActorSystemOptions) WithCluster(Server Server, clusterName, host string, port uint16) *ActorSystemOptions {
	o.ClusterName = clusterName
	return o.WithAddress(Server, host, port)
}

// WithAddress 设置主机地址
func (o *ActorSystemOptions) WithAddress(Server Server, host string, port uint16) *ActorSystemOptions {
	o.Host = host
	o.Port = port
	o.Server = Server
	return o
}

// WithClientFactory 设置远程调用客户端工厂
func (o *ActorSystemOptions) WithClientFactory(ClientFactory ClientFactory) *ActorSystemOptions {
	o.ClientFactory = ClientFactory
	return o
}
