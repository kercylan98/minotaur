package gateway

import (
	"fmt"
	"github.com/kercylan98/minotaur/core/transport"
	"github.com/kercylan98/minotaur/toolkit/router"
)

type bindServer struct {
	router *router.Multistage[func(kit *transport.GNETKit, conn *transport.Conn, request *RequestWrapper)]
}

func (b *bindServer) OnInit(kit *transport.GNETKit) {
	b.router = router.NewMultistage[func(kit *transport.GNETKit, conn *transport.Conn, request *RequestWrapper)]()
	kit.ConnectionPacketHook(b.onConnectionPacket)

	b.router.Route(MessageId_MI_Heartbeat, b.onHeartbeatRequest)
	b.router.Route(MessageId_MI_ServiceRegister, b.onServiceRegisterRequest)
}

func (b *bindServer) onConnectionPacket(kit *transport.GNETKit, conn *transport.Conn, packet transport.Packet) error {
	w, err := unwrapRequest(packet.GetBytes())
	if err != nil {
		return err
	}

	handler := b.router.Match(w.Id)
	if handler == nil {
		return fmt.Errorf("no handler for message id: %d", w.Id)
	}

	handler(kit, conn, w)
	return nil
}

func (b *bindServer) onHeartbeatRequest(kit *transport.GNETKit, conn *transport.Conn, request *RequestWrapper) {
	data, err := request.wrapResponse(MessageId_MI_Heartbeat, nil)
	if err != nil {
		panic(err)
		return
	}

	conn.Write(data)
}

func (b *bindServer) onServiceRegisterRequest(kit *transport.GNETKit, conn *transport.Conn, request *RequestWrapper) {
	var message ServiceRegisterRequest
	err := request.read(&message)
	if err != nil {
		panic(err)
	}

	// TODO: register service

	data, err := request.wrapResponse(MessageId_MI_ServiceRegister, &ServiceRegisterResponse{})
	if err != nil {
		panic(err)
	}

	conn.Write(data)
}
