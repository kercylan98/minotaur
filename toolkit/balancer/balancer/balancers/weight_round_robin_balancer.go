package balancers

import (
	"github.com/kercylan98/minotaur/toolkit/balancer/balancer"
	"sync"
)

const (
	// WeightRoundRobinBalancerDefaultWeight 默认权重
	WeightRoundRobinBalancerDefaultWeight = 10
	// WeightRoundRobinBalancerMaxWeight 最大权重
	WeightRoundRobinBalancerMaxWeight = 100
	// WeightRoundRobinBalancerMinWeight 最小权重
	WeightRoundRobinBalancerMinWeight = 0
)

var (
	_ balancer.Balancer[WeightRoundRobinBalancerInstance]        = (*implOfWeightRoundRobinBalancer[WeightRoundRobinBalancerInstance])(nil)
	_ WeightRoundRobinBalancer[WeightRoundRobinBalancerInstance] = (*implOfWeightRoundRobinBalancer[WeightRoundRobinBalancerInstance])(nil)
)

// NewWeightRoundRobinBalancerBuilder 创建一个加权轮询负载均衡器构建器(WeightRoundRobinBalancerBuilder)
func NewWeightRoundRobinBalancerBuilder[I WeightRoundRobinBalancerInstance]() WeightRoundRobinBalancerBuilder[I] {
	return &implOfWeightRoundRobinBalancerBuilder[I]{}
}

// WeightRoundRobinBalancerBuilder 加权轮询负载均衡器构建器接口，用于构建 WeightRoundRobinBalancer
type WeightRoundRobinBalancerBuilder[I WeightRoundRobinBalancerInstance] interface {
	// Build 构建一个 WeightRoundRobinBalancer
	Build() WeightRoundRobinBalancer[I]

	// InstancesOf 通过实例列表构建一个 WeightRoundRobinBalancer
	InstancesOf(instances ...I) WeightRoundRobinBalancer[I]
}

// implOfWeightRoundRobinBalancerBuilder 加权轮询负载均衡器构建器
type implOfWeightRoundRobinBalancerBuilder[I WeightRoundRobinBalancerInstance] struct{}

// Build 构建一个 WeightRoundRobinBalancer
func (i *implOfWeightRoundRobinBalancerBuilder[I]) Build() WeightRoundRobinBalancer[I] {
	return &implOfWeightRoundRobinBalancer[I]{
		weights:    make(map[balancer.InstanceID]int),
		rawWeights: make(map[balancer.InstanceID]int),
	}
}

// InstancesOf 通过实例列表构建一个 WeightRoundRobinBalancer
func (i *implOfWeightRoundRobinBalancerBuilder[I]) InstancesOf(instances ...I) WeightRoundRobinBalancer[I] {
	b := i.Build()
	for _, instance := range instances {
		b.RegisterInstance(instance)
	}
	return b
}

// WeightRoundRobinBalancerInstance 加权轮询负载均衡器实例接口
type WeightRoundRobinBalancerInstance interface {
	balancer.Instance

	// GetWeight 获取实例权重
	GetWeight() int
}

// WeightRoundRobinBalancer 加权轮询负载均衡器接口
//   - 该负载均衡器会根据实例的权重选择实例
//
// marks.ConcurrencySafe
type WeightRoundRobinBalancer[I balancer.Instance] interface {
	balancer.Balancer[I]

	// Select 根据权重选择一个实例，如果没有实例则返回 I 的零值
	//  - 如果实例权重为 WeightRoundRobinBalancerMinWeight 则不会被选择
	//
	//  - marks.ConcurrencySafe
	Select() (selected I)

	// RegisterInstance 添加一个实例到负载均衡器的实例列表
	//  - 实例的权重将使用 GetWeight 方法获取，并且用作初始化权重，后续可以通过 SetWeight 方法修改
	//  - 最小权重为: WeightRoundRobinBalancerMinWeight
	//  - 最大权重为: WeightRoundRobinBalancerMaxWeight
	//  - 如果超出范围则使用默认值: WeightRoundRobinBalancerDefaultWeight
	//
	//  - marks.ConcurrencySafe
	RegisterInstance(instance I)

	// DeregisterInstance 根据实例 ID 移除一个实例，如果没有找到则不执行任何操作
	//
	//  - marks.ConcurrencySafe
	DeregisterInstance(id balancer.InstanceID)

	// Count 返回实例数量
	//
	//  - marks.ConcurrencySafe
	Count() int

	// SetWeight 设置实例权重
	//   - 最小权重为: WeightRoundRobinBalancerMinWeight
	//   - 最大权重为: WeightRoundRobinBalancerMaxWeight
	//   - 如果超出范围则使用默认值: WeightRoundRobinBalancerDefaultWeight
	//
	//  - marks.ConcurrencySafe
	SetWeight(id balancer.InstanceID, weight int)
}

type implOfWeightRoundRobinBalancer[I WeightRoundRobinBalancerInstance] struct {
	instances  []I
	curr       int
	rawWeights map[balancer.InstanceID]int
	weights    map[balancer.InstanceID]int
	rw         sync.RWMutex
}

func (w *implOfWeightRoundRobinBalancer[I]) Select() (selected I) {
	w.rw.RLock()
	defer w.rw.RUnlock()

	if len(w.instances) == 0 {
		return
	}

	var totalWeight int
	for _, instance := range w.instances {
		rawWeight := w.rawWeights[instance.GetID()]
		totalWeight += rawWeight
		w.weights[instance.GetID()] += rawWeight
	}

	var n, index int

	for i, instance := range w.instances {
		if n < w.weights[instance.GetID()] {
			selected = instance
			index = i
		}
	}

	w.weights[w.instances[index].GetID()] -= totalWeight
	return
}

func (w *implOfWeightRoundRobinBalancer[I]) RegisterInstance(instance I) {
	weight := instance.GetWeight()
	if weight < WeightRoundRobinBalancerMinWeight || weight > WeightRoundRobinBalancerMaxWeight {
		weight = WeightRoundRobinBalancerDefaultWeight
	}

	w.rw.Lock()
	defer w.rw.Unlock()

	for _, ins := range w.instances {
		if ins.GetID() == instance.GetID() {
			return
		}
	}

	w.rawWeights[instance.GetID()] = weight
	w.instances = append(w.instances, instance)
}

func (w *implOfWeightRoundRobinBalancer[I]) DeregisterInstance(id balancer.InstanceID) {
	w.rw.Lock()
	defer w.rw.Unlock()

	delete(w.weights, id)
	delete(w.rawWeights, id)
	for idx, instance := range w.instances {
		if instance.GetID() == id {
			w.instances = append(w.instances[:idx], w.instances[idx+1:]...)
			return
		}
	}
}

func (w *implOfWeightRoundRobinBalancer[I]) Count() int {
	w.rw.RLock()
	defer w.rw.RUnlock()

	return len(w.instances)
}

func (w *implOfWeightRoundRobinBalancer[I]) SetWeight(id balancer.InstanceID, weight int) {
	if weight < WeightRoundRobinBalancerMinWeight || weight > WeightRoundRobinBalancerMaxWeight {
		weight = WeightRoundRobinBalancerDefaultWeight
	}

	w.rw.Lock()
	defer w.rw.Unlock()

	w.rawWeights[id] = weight
}
