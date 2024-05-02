package rpccore

import "github.com/kercylan98/minotaur/rpc"

type natsService struct {
	ServerInfo rpc.ServiceInfo `json:"server_info"`
	Routes     [][]rpc.Route   `json:"routes"`
}
