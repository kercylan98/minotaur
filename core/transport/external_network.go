package transport

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"reflect"
	"sync"
	"sync/atomic"
)

var externalNetworkNum atomic.Int32
var externalNetworkLaunchedNum atomic.Int32
var externalNetworkOnceLaunchInfo sync.Once

func NewWebSocket(addr string, pattern ...string) *ExternalNetwork {
	return newExternalNetwork(schemaWebSocket, addr, pattern...)
}

func NewTCP(addr string) *ExternalNetwork {
	return newExternalNetwork(schemaTcp, addr)
}

func NewTCP4(addr string) *ExternalNetwork {
	return newExternalNetwork(schemaTcp4, addr)
}

func NewTCP6(addr string) *ExternalNetwork {
	return newExternalNetwork(schemaTcp6, addr)
}

func NewUDP(addr string) *ExternalNetwork {
	return newExternalNetwork(schemaUdp, addr)
}

func NewUDP4(addr string) *ExternalNetwork {
	return newExternalNetwork(schemaUdp4, addr)
}

func NewUDP6(addr string) *ExternalNetwork {
	return newExternalNetwork(schemaUdp6, addr)
}

func NewUnix(addr string) *ExternalNetwork {
	return newExternalNetwork(schemaUnix, addr)
}

func newExternalNetwork(schema, addr string, pattern ...string) *ExternalNetwork {
	n := &ExternalNetwork{
		packetHandler:     func(conn *Conn, packet Packet) {},
		connOpenedHandler: func(conn *Conn) {},
		connClosedHandler: func(conn *Conn, err error) {},
	}
	n.producer = func() vivid.Actor {
		return newGnetEngine(n, schema, addr, pattern...)
	}
	return n
}

var _ vivid.Module = &ExternalNetwork{}

type ExternalNetworkPacketHandler func(conn *Conn, packet Packet)
type ExternalNetworkConnOpenedHandler func(conn *Conn)
type ExternalNetworkConnClosedHandler func(conn *Conn, err error)

type ExternalNetwork struct {
	support           *vivid.ModuleSupport
	producer          vivid.ActorProducer
	packetHandler     ExternalNetworkPacketHandler
	connOpenedHandler ExternalNetworkConnOpenedHandler
	connClosedHandler ExternalNetworkConnClosedHandler
}

func (n *ExternalNetwork) SetConnOpenedHandler(handler ExternalNetworkConnOpenedHandler) *ExternalNetwork {
	n.connOpenedHandler = handler
	return n
}

func (n *ExternalNetwork) SetConnClosedHandler(handler ExternalNetworkConnClosedHandler) *ExternalNetwork {
	n.connClosedHandler = handler
	return n
}

func (n *ExternalNetwork) SetPacketHandler(handler ExternalNetworkPacketHandler) *ExternalNetwork {
	n.packetHandler = handler
	return n
}

func (n *ExternalNetwork) OnLoad(support *vivid.ModuleSupport, hasTransportModule bool) {
	n.support = support
	actorType := reflect.TypeOf(n.producer()).Elem().Name()
	n.support.System().ActorOf(n.producer, func(options *vivid.ActorOptions) {
		options.WithNamePrefix(actorType)
	})
}
