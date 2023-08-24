package gateway

import "time"

// EndpointOption 网关端点选项
type EndpointOption func(endpoint *Endpoint)

// WithEndpointStateEvaluator 设置端点健康值评估函数
func WithEndpointStateEvaluator(evaluator func(costUnixNano float64) float64) EndpointOption {
	return func(endpoint *Endpoint) {
		endpoint.evaluator = evaluator
	}
}

// WithReconnectInterval 设置端点重连间隔
//   - 默认为 DefaultEndpointReconnectInterval
//   - 端点在连接失败后会在该间隔后重连，如果 <= 0 则不会重连
func WithReconnectInterval(interval time.Duration) EndpointOption {
	return func(endpoint *Endpoint) {
		endpoint.rci = interval
	}
}
