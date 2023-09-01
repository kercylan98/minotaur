package gateway_test

import (
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"github.com/kercylan98/minotaur/server/gateway"
	"github.com/kercylan98/minotaur/utils/log"
	"testing"
	"time"
)

type Scanner struct {
}

func (slf *Scanner) GetEndpoints() ([]*gateway.Endpoint, error) {
	return []*gateway.Endpoint{
		gateway.NewEndpoint("test", client.NewWebsocket("ws://127.0.0.1:8889"), gateway.WithEndpointConnectionPoolSize(10)),
	}, nil
}

func (slf *Scanner) GetInterval() time.Duration {
	return time.Second
}

func TestGateway_RunEndpointServer(t *testing.T) {
	srv := server.New(server.NetworkWebsocket, server.WithDeadlockDetect(time.Second*3))
	srv.RegConnectionPacketPreprocessEvent(func(srv *server.Server, conn *server.Conn, packet []byte, abort func(), usePacket func(newPacket []byte)) {
		addr, packet, err := gateway.UnmarshalGatewayOutPacket(packet)
		if err != nil {
			// 非网关的普通数据包
			return
		}
		usePacket(packet)
		conn.SetMessageData("gw-addr", addr)
	})
	srv.RegConnectionWritePacketBeforeEvent(func(srv *server.Server, conn *server.Conn, packet []byte) []byte {
		addr, ok := conn.GetMessageData("gw-addr").(string)
		if !ok {
			return packet
		}
		packet, err := gateway.MarshalGatewayInPacket(addr, time.Now().Unix(), packet)
		if err != nil {
			panic(err)
		}
		return packet
	})
	srv.RegConnectionReceivePacketEvent(func(srv *server.Server, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	srv.RegConnectionOpenedEvent(func(srv *server.Server, conn *server.Conn) {
		log.Info("connection opened", log.String("conn", conn.GetID()))
	})
	if err := srv.Run(":8889"); err != nil {
		panic(err)
	}
}

func TestGateway_Run(t *testing.T) {
	gw := gateway.NewGateway(server.New(server.NetworkWebsocket, server.WithDeadlockDetect(time.Second*3)), new(Scanner))
	gw.RegConnectionReceivePacketEventHandle(func(gateway *gateway.Gateway, conn *server.Conn, packet []byte) {
		endpoint, err := gateway.GetEndpoint("test")
		if err == nil {
			endpoint.Forward(conn, packet)
		}
	})
	gw.RegEndpointConnectReceivePacketEventHandle(func(gateway *gateway.Gateway, endpoint *gateway.Endpoint, conn *server.Conn, packet []byte) {
		conn.Write(packet)
	})
	if err := gw.Run(":8888"); err != nil {
		panic(err)
	}
}
