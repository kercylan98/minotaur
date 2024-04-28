package gateway

import (
	"time"
)

const (
	DefaultEndpointReconnectInterval  = time.Second
	DefaultEndpointConnectionPoolSize = 1
)

// EndpointOption 网关端点选项
type EndpointOption func(endpoint *Endpoint)

// WithEndpointStateEvaluator 设置端点健康值评估函数
func WithEndpointStateEvaluator(evaluator func(costUnixNano float64) float64) EndpointOption {
	return func(endpoint *Endpoint) {
		endpoint.evaluator = evaluator
	}
}

// WithEndpointConnectionPoolSize 设置端点连接池大小
//   - 默认为 DefaultEndpointConnectionPoolSize
//   - 端点连接池大小决定了网关服务器与端点服务器建立的连接数，如果 <= 0 则会使用默认值
//   - 在网关服务器中，多个客户端在发送消息到端点服务器时，会共用一个连接，适当的增大连接池大小可以提高网关服务器的承载能力
func WithEndpointConnectionPoolSize(size int) EndpointOption {
	return func(endpoint *Endpoint) {
		endpoint.cps = size
	}
}

// WithEndpointReconnectInterval 设置端点重连间隔
//   - 默认为 DefaultEndpointReconnectInterval
//   - 端点在连接失败后会在该间隔后重连，如果 <= 0 则不会重连
func WithEndpointReconnectInterval(interval time.Duration) EndpointOption {
	return func(endpoint *Endpoint) {
		endpoint.rci = interval
	}
}
