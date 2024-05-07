package balancer

type SelectOptions struct {
	ConsistencyKey string // ConsistencyKey 一致性哈希的键
}

// NewSelectOptions 创建一个新的 SelectOptions
func NewSelectOptions() *SelectOptions {
	return &SelectOptions{}
}

// WithConsistencyKey 设置 ConsistencyKey
func (o *SelectOptions) WithConsistencyKey(consistencyKey string) *SelectOptions {
	o.ConsistencyKey = consistencyKey
	return o
}

// Apply 应用多个选项
func (o *SelectOptions) Apply(opts ...*SelectOptions) *SelectOptions {
	for _, opt := range opts {
		o.ConsistencyKey = opt.ConsistencyKey
	}
	return o
}
