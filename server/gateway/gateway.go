package gateway

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/utils/random"
	"math"
	"sync"
	"time"
)

type (
	// EndpointSelector 端点选择器，用于从多个端点中选择一个可用的端点，如果没有可用的端点则返回 nil
	EndpointSelector func(endpoints []*Endpoint) *Endpoint
)

// NewGateway 基于 server.Server 创建网关服务器
//   - behaviorController 行为控制函数决定了客户端与网关服务器建立连接及接收数据包后的行为
func NewGateway(srv *server.Server, scanner Scanner, options ...Option) *Gateway {
	gateway := &Gateway{
		events:  newEvents(),
		srv:     srv,
		scanner: scanner,
		es:      make(map[string]map[string]*Endpoint),
		ess: func(endpoints []*Endpoint) *Endpoint {
			return endpoints[random.Int(0, len(endpoints)-1)]
		},
		cce: make(map[string]*Endpoint),
	}
	for _, option := range options {
		option(gateway)
	}
	return gateway
}

// Gateway 网关
type Gateway struct {
	*events
	srv     *server.Server                  // 网关服务器核心
	scanner Scanner                         // 端点扫描器
	es      map[string]map[string]*Endpoint // 端点列表 [name][address]
	esm     sync.Mutex                      // 端点列表锁
	ess     EndpointSelector                // 端点选择器
	closed  bool                            // 网关是否已关闭
	running bool                            // 网关是否正在运行
	cce     map[string]*Endpoint            // 连接当前连接的端点 [conn.ID]
	cceLock sync.RWMutex                    // 连接当前连接的端点锁
}

// Run 运行网关
func (slf *Gateway) Run(addr string) error {
	if slf.closed {
		return ErrGatewayClosed
	}
	if slf.running {
		return ErrGatewayRunning
	}
	slf.srv.RegStartFinishEvent(func(srv *server.Server) {
		go func() {
			for !slf.closed {
				endpoints, err := slf.scanner.GetEndpoints()
				if err != nil {
					continue
				}
				slf.esm.Lock()
				for _, endpoint := range endpoints {
					es, exist := slf.es[endpoint.GetName()]
					if !exist {
						es = make(map[string]*Endpoint)
						slf.es[endpoint.GetName()] = es
					}
					e, exist := es[endpoint.GetAddress()]
					if !exist {
						e = endpoint
						es[endpoint.GetAddress()] = e
						go e.connect(slf)
					}
				}
				slf.esm.Unlock()
				time.Sleep(slf.scanner.GetInterval())
			}
		}()
	}, math.MinInt)
	slf.srv.RegStopEvent(func(srv *server.Server) {
		slf.Shutdown()
	}, math.MinInt)
	slf.srv.RegConnectionOpenedEvent(func(srv *server.Server, conn *server.Conn) {
		slf.OnConnectionOpenedEvent(slf, conn)
	}, math.MinInt)
	slf.srv.RegConnectionClosedEvent(func(srv *server.Server, conn *server.Conn, err any) {
		slf.OnConnectionClosedEvent(slf, conn)
	}, math.MinInt)
	slf.srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		slf.OnConnectionReceivePacketEvent(slf, conn, packet)
	}, math.MinInt)
	slf.running = true
	if err := slf.srv.Run(addr); err != nil {
		return err
	}
	slf.running = false
	return nil
}

// Shutdown 关闭网关
func (slf *Gateway) Shutdown() {
	if !slf.closed {
		return
	}
	slf.closed = true
	slf.srv.Shutdown()
}

// Server 获取网关服务器核心
func (slf *Gateway) Server() *server.Server {
	return slf.srv
}

// GetEndpoint 获取一个可用的端点
//   - name: 端点名称
func (slf *Gateway) GetEndpoint(name string) (*Endpoint, error) {
	slf.esm.Lock()
	endpoints, exist := slf.es[name]
	if !exist || len(endpoints) == 0 {
		delete(slf.es, name)
		slf.esm.Unlock()
		return nil, ErrEndpointNotExists
	}

	var available = make([]*Endpoint, 0, len(endpoints))
	for _, e := range endpoints {
		if e.GetState() > 0 {
			available = append(available, e)
		}
	}
	slf.esm.Unlock()
	if len(available) == 0 {
		return nil, ErrEndpointNotExists
	}

	endpoint := slf.ess(available)
	if endpoint == nil {
		return nil, ErrEndpointNotExists
	}
	return endpoint, nil
}

// GetConnEndpoint 获取一个可用的端点，如果客户端已经连接到了某个端点，将优先返回该端点
//   - 当连接到的端点不可用或没有连接记录时，效果同 GetEndpoint 相同
//   - 当连接行为为有状态时，推荐使用该方法
func (slf *Gateway) GetConnEndpoint(name string, conn *server.Conn) (*Endpoint, error) {
	slf.cceLock.RLock()
	endpoint, exist := slf.cce[conn.GetID()]
	slf.cceLock.RUnlock()
	if exist && endpoint.GetState() > 0 {
		return endpoint, nil
	}
	return slf.GetEndpoint(name)
}

// SwitchEndpoint 将端点端点的所有连接切换到另一个端点
func (slf *Gateway) SwitchEndpoint(source, dest *Endpoint) {
	if source.name == dest.name && source.address == dest.address || source.GetState() <= 0 || dest.GetState() <= 0 {
		return
	}
	slf.cceLock.Lock()
	for id, endpoint := range slf.cce {
		if endpoint == source {
			slf.cce[id] = dest
		}
	}
	slf.cceLock.Unlock()
}
