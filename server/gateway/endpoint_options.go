package gateway

// EndpointOption 网关端点选项
type EndpointOption func(endpoint *Endpoint)

// WithEndpointStateEvaluator 设置端点健康值评估函数
func WithEndpointStateEvaluator(evaluator func(costUnixNano float64) float64) EndpointOption {
	return func(endpoint *Endpoint) {
		endpoint.evaluator = evaluator
	}
}
