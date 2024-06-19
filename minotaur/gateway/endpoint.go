package gateway

import (
	"net"
	"strconv"
)

type EndpointId = int64

type Endpoint struct {
	Id     EndpointId `json:"id"`     // 端点标识
	Host   Host       `json:"addr"`   // 端点地址
	Port   Port       `json:"port"`   // 端点端口
	Weight uint64     `json:"weight"` // 端点权重
}

func (e *Endpoint) GetId() EndpointId {
	return e.Id
}

func (e *Endpoint) GetWeight() int {
	return int(e.Weight)
}

func (e *Endpoint) GetAddress() Address {
	return net.JoinHostPort(e.Host, strconv.Itoa(int(e.Port)))
}
