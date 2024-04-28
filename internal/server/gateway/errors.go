package gateway

import "errors"

var (
	// ErrEndpointNotExists 该名称下不存在任何端点
	ErrEndpointNotExists = errors.New("gateway: endpoint not exists")
	// ErrGatewayClosed 网关已关闭
	ErrGatewayClosed = errors.New("gateway: gateway closed")
	// ErrGatewayRunning 网关正在运行
	ErrGatewayRunning = errors.New("gateway: gateway running")
	// ErrConnectionNotFount 该端点下不存在该连接
	ErrConnectionNotFount = errors.New("gateway: connection not found")
)
