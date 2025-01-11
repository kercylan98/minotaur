package balancers_test

import (
	"github.com/kercylan98/minotaur/toolkit/balancer/balancer"
	"github.com/kercylan98/minotaur/toolkit/balancer/balancer/balancers"
)

var _ balancer.Instance = (*TestInstance)(nil)
var _ balancers.WeightRoundRobinBalancerInstance = (*TestInstance)(nil)

type TestInstance struct {
	ID     balancer.InstanceID
	weight int
}

func (t *TestInstance) GetWeight() int {
	return t.weight
}

func (t *TestInstance) GetID() balancer.InstanceID {
	return t.ID
}

func NewTestInstance(id balancer.InstanceID) *TestInstance {
	return &TestInstance{
		ID: id,
	}
}

func NewTestInstances(ids ...balancer.InstanceID) []*TestInstance {
	instances := make([]*TestInstance, 0, len(ids))
	for _, id := range ids {
		instances = append(instances, NewTestInstance(id))
	}
	return instances
}

func NewWeightTestInstance(id balancer.InstanceID, weight int) *TestInstance {
	return &TestInstance{
		ID:     id,
		weight: weight,
	}
}
