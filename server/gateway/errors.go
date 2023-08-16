package gateway

import "errors"

var (
	// ErrEndpointAlreadyExists 网关端点已存在
	ErrEndpointAlreadyExists = errors.New("gateway: endpoint already exists")
	// ErrCannotAddRunningEndpoint 无法添加一个正在运行的网关端点
	ErrCannotAddRunningEndpoint = errors.New("gateway: cannot add a running endpoint")
	// ErrEndpointNotExists 该名称下不存在任何端点
	ErrEndpointNotExists = errors.New("gateway: endpoint not exists")
)
