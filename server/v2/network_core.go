package server

import (
	"github.com/kercylan98/minotaur/utils/hub"
	"golang.org/x/net/context"
	"net"
)

type ConnWriter func(message Packet) error

type NetworkCore interface {
	OnConnectionOpened(ctx context.Context, conn net.Conn, writer ConnWriter)

	OnConnectionClosed(conn Conn)

	OnReceivePacket(packet Packet)

	GeneratePacket(data []byte) Packet
}

type networkCore struct {
	*server
	packetPool *hub.ObjectPool[*packet]
}

func (ne *networkCore) init(srv *server) *networkCore {
	ne.server = srv
	ne.packetPool = hub.NewObjectPool(func() *packet {
		return new(packet)
	}, func(data *packet) {
		data.reset()
	})
	return ne
}

func (ne *networkCore) OnConnectionOpened(ctx context.Context, conn net.Conn, writer ConnWriter) {

}

func (ne *networkCore) OnConnectionClosed(conn Conn) {

}

func (ne *networkCore) OnReceivePacket(packet Packet) {

}

func (ne *networkCore) GeneratePacket(data []byte) Packet {
	return ne.packetPool.Get().init(data)
}
