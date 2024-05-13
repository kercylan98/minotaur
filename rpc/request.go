package rpc

import "encoding/json"

// NewRequest 用于创建一个新的 RPC 请求
func NewRequest(route Route, data []byte) *Request {
	return &Request{
		Route: route,
		Data:  data,
	}
}

// Request 是用于发起 RPC 调用的请求，其中 RPC 的调用细节
type Request struct {
	Route Route           `json:"route"`
	Data  json.RawMessage `json:"data"`
}
