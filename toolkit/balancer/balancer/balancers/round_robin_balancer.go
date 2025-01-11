package balancers

import "github.com/kercylan98/minotaur/toolkit/balancer/balancer"

var (
	_ balancer.Balancer[balancer.Instance]  = (*implOfRoundRobinBalancer[balancer.Instance])(nil)
	_ RoundRobinBalancer[balancer.Instance] = (*implOfRoundRobinBalancer[balancer.Instance])(nil)
)

// NewRoundRobinBalancerBuilder 创建一个轮询负载均衡器构建器(RoundRobinBalancerBuilder)
func NewRoundRobinBalancerBuilder[I balancer.Instance]() RoundRobinBalancerBuilder[I] {
	return &implOfRoundRobinBalancerBuilder[I]{}
}

// RoundRobinBalancerBuilder 轮询负载均衡器构建器接口，用于构建 RoundRobinBalancer
type RoundRobinBalancerBuilder[I balancer.Instance] interface {
	// Build 构建一个 RoundRobinBalancer
	Build() RoundRobinBalancer[I]

	// InstancesOf 通过实例列表构建一个 RoundRobinBalancer
	InstancesOf(instances ...I) RoundRobinBalancer[I]
}

type implOfRoundRobinBalancerBuilder[I balancer.Instance] struct{}

func (r *implOfRoundRobinBalancerBuilder[I]) Build() RoundRobinBalancer[I] {
	return &implOfRoundRobinBalancer[I]{
		index: -1,
	}
}

func (r *implOfRoundRobinBalancerBuilder[I]) InstancesOf(instances ...I) RoundRobinBalancer[I] {
	b := r.Build()
	for _, instance := range instances {
		b.RegisterInstance(instance)
	}
	return b
}

// RoundRobinBalancer 轮询负载均衡器接口
//   - 该负载均衡器会依次选择每个实例，当选择到最后一个实例后，再次选择第一个实例
//
// marks.ConcurrencyUnsafe
type RoundRobinBalancer[I balancer.Instance] interface {
	balancer.Balancer[I]

	// Select 选择一个实例，如果没有实例则返回 I 的零值
	//
	//  - marks.ConcurrencyUnsafe
	Select() (selected I)

	// RegisterInstance 添加一个实例到负载均衡器的实例列表尾部
	//
	//  - marks.ConcurrencyUnsafe
	//  - marks.ExcludeDuplicateElements
	RegisterInstance(instance I)

	// DeregisterInstance 根据实例 ID 移除一个实例，如果没有找到则不执行任何操作
	//
	//  - marks.ConcurrencyUnsafe
	DeregisterInstance(id balancer.InstanceID)

	// Count 返回实例数量
	//
	//  - marks.ConcurrencyUnsafe
	Count() int
}

type implOfRoundRobinBalancer[I balancer.Instance] struct {
	instances []I
	index     int
}

func (r *implOfRoundRobinBalancer[I]) Select() (selected I) {
	if len(r.instances) == 0 {
		return
	}

	r.index = (r.index + 1) % len(r.instances)
	selected = r.instances[r.index]
	return
}

func (r *implOfRoundRobinBalancer[I]) RegisterInstance(instance I) {
	for _, i := range r.instances {
		if i.GetID() == instance.GetID() {
			return
		}
	}
	r.instances = append(r.instances, instance)
}

func (r *implOfRoundRobinBalancer[I]) DeregisterInstance(id balancer.InstanceID) {
	for i, instance := range r.instances {
		if instance.GetID() == id {
			r.instances = append(r.instances[:i], r.instances[i+1:]...)
			return
		}
	}
}

func (r *implOfRoundRobinBalancer[I]) Count() int {
	return len(r.instances)
}
