package rpccore

import (
	"encoding/json"
	"github.com/kercylan98/minotaur/rpc"
)

type natsPacket struct {
	IsRequest bool               `json:"r"`           // 是否为请求
	Request   natsPacketRequest  `json:"i,omitempty"` // 请求数据
	Response  natsPacketResponse `json:"o,omitempty"` // 响应数据
}

type natsPacketRequest struct {
	Routes []rpc.Route     `json:"r"`
	Data   json.RawMessage `json:"q"`
}

type natsPacketResponse struct {
}
