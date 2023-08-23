package gateway_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server"
	"github.com/kercylan98/minotaur/server/client"
	"github.com/kercylan98/minotaur/server/gateway"
	"testing"
	"time"
)

func TestGateway_RunEndpointServer(t *testing.T) {
	srv := server.New(server.NetworkWebsocket, server.WithDeadlockDetect(time.Second*3))
	srv.RegConnectionClosedEvent(func(srv *server.Server, conn *server.Conn, err any) {
		fmt.Println(err)
	})
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
		fmt.Println("endpoint receive packet", string(packet))
		conn.Write(packet)
	})
	if err := srv.Run(":8889"); err != nil {
		panic(err)
	}
}

func TestGateway_Run(t *testing.T) {
	srv := server.New(server.NetworkWebsocket, server.WithDeadlockDetect(time.Second*3))
	gw := gateway.NewGateway(srv)
	srv.RegStartFinishEvent(func(srv *server.Server) {
		if err := gw.AddEndpoint(gateway.NewEndpoint(gw, "test", client.NewWebsocket("ws://127.0.0.1:8889"))); err != nil {
			panic(err)
		}
	})
	if err := gw.Run(":8888"); err != nil {
		panic(err)
	}
}
