package gateway

import (
	"github.com/alphadose/haxmap"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"time"
)

// NewEndpoint 创建网关端点
func NewEndpoint(gateway *Gateway, name string, client *client.Client, options ...EndpointOption) *Endpoint {
	endpoint := &Endpoint{
		gateway:     gateway,
		client:      client,
		name:        name,
		address:     client.GetServerAddr(),
		connections: haxmap.New[string, *server.Conn](),
	}
	for _, option := range options {
		option(endpoint)
	}
	if endpoint.evaluator == nil {
		endpoint.evaluator = func(costUnixNano float64) float64 {
			return 1 / (1 + 1.5*time.Duration(costUnixNano).Seconds())
		}
	}
	endpoint.client.RegConnectionClosedEvent(endpoint.onConnectionClosed)
	endpoint.client.RegConnectionReceivePacketEvent(endpoint.onConnectionReceivePacket)
	return endpoint
}

// Endpoint 网关端点
type Endpoint struct {
	gateway     *Gateway
	client      *client.Client                     // 端点客户端
	name        string                             // 端点名称
	address     string                             // 端点地址
	state       float64                            // 端点健康值（0为不可用，越高越优）
	offline     bool                               // 离线
	evaluator   func(costUnixNano float64) float64 // 端点健康值评估函数
	connections *haxmap.Map[string, *server.Conn]  // 连接列表
}

// Link 连接端点
func (slf *Endpoint) Link(conn *server.Conn) {
	slf.connections.Set(conn.GetID(), conn)
}

// Unlink 断开连接
func (slf *Endpoint) Unlink(conn *server.Conn) {
	slf.connections.Del(conn.GetID())
}

// GetLink 获取连接
func (slf *Endpoint) GetLink(id string) *server.Conn {
	conn, _ := slf.connections.Get(id)
	return conn
}

// Offline 离线
func (slf *Endpoint) Offline() {
	slf.offline = true
}

// Connect 连接端点
func (slf *Endpoint) Connect() {
	for {
		cur := time.Now().UnixNano()
		if err := slf.client.Run(); err == nil {
			slf.state = slf.evaluator(float64(time.Now().UnixNano() - cur))
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// Write 写入数据
func (slf *Endpoint) Write(packet []byte, callback ...func(err error)) {
	slf.client.Write(packet, callback...)
}

// WriteWS 写入 websocket 数据
func (slf *Endpoint) WriteWS(wst int, packet []byte, callback ...func(err error)) {
	slf.client.WriteWS(wst, packet, callback...)
}

// onConnectionClosed 与端点连接断开事件
func (slf *Endpoint) onConnectionClosed(conn *client.Client, err any) {
	if !slf.offline {
		go slf.Connect()
	}
}

// onConnectionReceivePacket 接收到来自端点的数据包事件
func (slf *Endpoint) onConnectionReceivePacket(conn *client.Client, wst int, packet []byte) {
	addr, sendTime, packet, err := UnmarshalGatewayInPacket(packet)
	if err != nil {
		panic(err)
	}
	slf.state = slf.evaluator(float64(time.Now().UnixNano() - sendTime))
	slf.GetLink(addr).SetWST(wst).Write(packet)
}
