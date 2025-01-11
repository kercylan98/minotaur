package balancers_test

import (
	"github.com/kercylan98/minotaur/toolkit/balancer/balancer/balancers"
	"testing"
)

func TestNewRoundRobinBalancerBuilder(t *testing.T) {
	builder := balancers.NewRoundRobinBalancerBuilder[*TestInstance]()
	if builder == nil {
		t.Fatal("NewRoundRobinBalancerBuilder failed")
	}
}

func TestImplOfRoundRobinBalancerBuilder_Build(t *testing.T) {
	builder := balancers.NewRoundRobinBalancerBuilder[*TestInstance]()
	balancer := builder.Build()
	if balancer == nil {
		t.Fatal("Build failed")
	}
}

func TestImplOfRoundRobinBalancerBuilder_InstancesOf(t *testing.T) {
	builder := balancers.NewRoundRobinBalancerBuilder[*TestInstance]()
	instances := NewTestInstances("1", "2", "3")
	balancer := builder.InstancesOf(instances...)
	if balancer == nil {
		t.Fatal("InstancesOf failed")
	}
}

func TestImplOfRoundRobinBalancer_Select(t *testing.T) {
	builder := balancers.NewRoundRobinBalancerBuilder[*TestInstance]()
	instances := NewTestInstances("1", "2", "3")
	balancer := builder.InstancesOf(instances...)
	if balancer == nil {
		t.Fatal("InstancesOf failed")
	}
	selected := balancer.Select()
	if selected == nil {
		t.Fatal("Select failed")
	}

	if selected.GetID() != "1" {
		t.Fatal("Select target error")
	}
}

func TestImplOfRoundRobinBalancer_RegisterInstance(t *testing.T) {
	builder := balancers.NewRoundRobinBalancerBuilder[*TestInstance]()
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

func TestImplOfRoundRobinBalancer_DeregisterInstance(t *testing.T) {
	builder := balancers.NewRoundRobinBalancerBuilder[*TestInstance]()
	instances := NewTestInstances("1", "2", "3")
	balancer := builder.InstancesOf(instances...)
	if balancer == nil {
		t.Fatal("InstancesOf failed")
	}
	balancer.DeregisterInstance("2")
	balancer.DeregisterInstance("0")

	if balancer.Count() != 2 {
		t.Fatal("DeregisterInstance failed")
	}
}
