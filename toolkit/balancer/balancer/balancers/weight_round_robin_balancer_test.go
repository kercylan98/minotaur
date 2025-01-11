package balancers_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/balancer/balancer/balancers"
	"testing"
)

func TestNewWeightRoundRobinBalancerBuilder(t *testing.T) {
	builder := balancers.NewWeightRoundRobinBalancerBuilder[*TestInstance]()
	if builder == nil {
		t.Fatal("NewRoundRobinBalancerBuilder failed")
	}
}

func TestImplOfWeightRoundRobinBalancerBuilder_Build(t *testing.T) {
	builder := balancers.NewWeightRoundRobinBalancerBuilder[*TestInstance]()
	balancer := builder.Build()
	if balancer == nil {
		t.Fatal("Build failed")
	}
}

func TestImplOfWeightRoundRobinBalancerBuilder_InstancesOf(t *testing.T) {
	builder := balancers.NewWeightRoundRobinBalancerBuilder[*TestInstance]()
	instances := NewTestInstances("1", "2", "3")
	balancer := builder.InstancesOf(instances...)
	if balancer == nil {
		t.Fatal("InstancesOf failed")
	}
}

func TestImplOfWeightRoundRobinBalancer_Select(t *testing.T) {
	builder := balancers.NewWeightRoundRobinBalancerBuilder[*TestInstance]()
	instances := []*TestInstance{
		NewWeightTestInstance("1", balancers.WeightRoundRobinBalancerDefaultWeight),
		NewWeightTestInstance("2", balancers.WeightRoundRobinBalancerDefaultWeight),
		NewWeightTestInstance("3", balancers.WeightRoundRobinBalancerDefaultWeight),
	}
	balancer := builder.InstancesOf(instances...)
	if balancer == nil {
		t.Fatal("InstancesOf failed")
	}
	selected := balancer.Select()
	if selected == nil {
		t.Fatal("Select failed")
	}

	t.Log("selected", selected.GetID())
}

func TestImplOfWeightRoundRobinBalancer_RegisterInstance(t *testing.T) {
	builder := balancers.NewWeightRoundRobinBalancerBuilder[*TestInstance]()
	instances := NewTestInstances("1", "2", "3")
	balancer := builder.InstancesOf(instances...)
	if balancer == nil {
		t.Fatal("InstancesOf failed")
	}
	balancer.RegisterInstance(NewTestInstance("4"))

	if balancer.Count() != 4 {
		t.Fatal("RegisterInstance failed")
	}
}

func TestImplOfWeightRoundRobinBalancer_DeregisterInstance(t *testing.T) {
	builder := balancers.NewWeightRoundRobinBalancerBuilder[*TestInstance]()
	instances := NewTestInstances("1", "2", "3")
	balancer := builder.InstancesOf(instances...)
	if balancer == nil {
		t.Fatal("InstancesOf failed")
	}
	balancer.DeregisterInstance("1")

	if balancer.Count() != 2 {
		t.Fatal("DeregisterInstance failed")
	}
}

func TestImplOfWeightRoundRobinBalancer_SetWeight(t *testing.T) {
	builder := balancers.NewWeightRoundRobinBalancerBuilder[*TestInstance]()
	instances := []*TestInstance{
		NewWeightTestInstance("1", 0),
		NewWeightTestInstance("2", 2),
		NewWeightTestInstance("3", 3),
	}
	balancer := builder.InstancesOf(instances...)
	if balancer == nil {
		t.Fatal("InstancesOf failed")
	}

	for i := 0; i < 100; i++ {
		selected := balancer.Select()
		if selected == nil || selected.GetID() == "1" {
			t.Fatal("Select target error")
		}
	}

	balancer.SetWeight("1", 100)
	balancer.SetWeight("2", 0)
	balancer.SetWeight("3", 0)

	for i := 0; i < 100; i++ {
		selected := balancer.Select()
		if selected == nil || selected.GetID() != "1" {
			t.Fatal("Select target error")
		}
	}

}

func TestImplOfWeightRoundRobinBalancer_Count(t *testing.T) {
	builder := balancers.NewWeightRoundRobinBalancerBuilder[*TestInstance]()
	instances := NewTestInstances("1", "2", "3")
	balancer := builder.InstancesOf(instances...)
	if balancer == nil {
		t.Fatal("InstancesOf failed")
	}

	if balancer.Count() != 3 {
		t.Fatal("Count failed")
	}
}

func TestImplOfWeightRoundRobinBalancer_WeightSelect(t *testing.T) {
	var instances []*TestInstance
	for i := 0; i < 9; i++ {
		instances = append(instances, NewWeightTestInstance(fmt.Sprint(i+1), balancers.WeightRoundRobinBalancerDefaultWeight))
	}
	instances = append(instances, NewWeightTestInstance("9", balancers.WeightRoundRobinBalancerMaxWeight))

	builder := balancers.NewWeightRoundRobinBalancerBuilder[*TestInstance]()
	balancer := builder.InstancesOf(instances...)

	for i := 0; i < 100; i++ {
		selected := balancer.Select()
		if selected == nil {
			t.Fatal("WeightSelect failed")
		}

		t.Log("selected", i+1, selected.GetID())
	}
}
