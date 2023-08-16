package gateway

// Option 网关选项
type Option func(gateway *Gateway)

// WithEndpointSelector 设置端点选择器
//   - 默认情况下，网关会随机选择一个端点作为目标，如果需要自定义端点选择器，可以通过该选项设置
func WithEndpointSelector(selector func([]*Endpoint) *Endpoint) Option {
	return func(gateway *Gateway) {
		gateway.EndpointManager.selector = selector
	}
}
