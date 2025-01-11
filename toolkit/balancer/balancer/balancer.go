package balancer

// Balancer 负载均衡器通用接口
type Balancer[I Instance] interface {
	// Select 选择一个实例，如果没有实例则返回其零值
	Select() (selected I)

	// RegisterInstance 注册示例到负载均衡器
	RegisterInstance(instance I)

	// DeregisterInstance 根据实例 ID 移除一个实例，如果没有找到则不执行任何操作
	DeregisterInstance(id InstanceID)

	// Count 返回实例数量
	Count() int
}
