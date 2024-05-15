package vivid

import "time"

// NewActorSystemOptions 用于创建一个 ActorSystemOptions
func NewActorSystemOptions() *ActorSystemOptions {
	return (&ActorSystemOptions{
		Dispatchers: make(map[string]Dispatcher),
	}).
		WithNetwork("tcp").
		WithDispatcher("default", newDispatcher()).
		WithAskDefaultTimeout(time.Second * 2)
}

// ActorSystemOptions 是 ActorSystem 的配置选项
type ActorSystemOptions struct {
	ClusterName       string                // 集群名称
	Network           string                // 网络协议
	Host              string                // 主机地址
	Port              uint16                // 端口
	Server            Server                // 远程调用服务端
	ClientFactory     ClientFactory         // 远程调用客户端
	Dispatchers       map[string]Dispatcher // 消息分发器
	AskDefaultTimeout time.Duration         // 默认的 Ask 超时时间
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
		for r, d := range option.Dispatchers {
			o.Dispatchers[r] = d
		}
		if option.AskDefaultTimeout != 0 {
			o.AskDefaultTimeout = option.AskDefaultTimeout
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

// WithDispatcher 绑定消息分发器
func (o *ActorSystemOptions) WithDispatcher(name string, dispatcher Dispatcher) *ActorSystemOptions {
	o.Dispatchers[name] = dispatcher
	return o
}

// WithAskDefaultTimeout 设置默认的 Ask 超时时间
func (o *ActorSystemOptions) WithAskDefaultTimeout(timeout time.Duration) *ActorSystemOptions {
	o.AskDefaultTimeout = timeout
	return o
}
