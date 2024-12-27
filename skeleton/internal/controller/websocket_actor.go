package controller

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/api/rpcmessage"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/toolkit"
)

func newWebSocketActor(rpc module.RPCModule) *webSocketActor {
	return &webSocketActor{
		rpc: rpc,
	}
}

type webSocketActor struct {
	rpc module.RPCModule
}

func (w *webSocketActor) OnOpened(ctx vivid.ActorContext, socket socket.Socket) {

}

func (w *webSocketActor) OnClose(ctx vivid.ActorContext, socket socket.Socket, err error) {

}

func (w *webSocketActor) OnPacket(ctx vivid.ActorContext, socket socket.Socket, packet *socket.Packet) {
	result, err := w.rpc.Request(&rpcmessage.RPCOAuthLoginRequest{
		Account:  "account",
		Password: "password",
	}).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(toolkit.MarshalJSON(result)))

	socket.WritePacket(packet)
}

func (w *webSocketActor) OnReceive(ctx vivid.ActorContext) {
	
}
