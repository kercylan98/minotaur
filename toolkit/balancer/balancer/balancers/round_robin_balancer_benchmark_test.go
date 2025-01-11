package balancers_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/balancer/balancer/balancers"
	"testing"
)

func BenchmarkImplOfRoundRobinBalancer_Select(b *testing.B) {
	tests := []struct {
		name          string
		instanceCount int
	}{
		{"InstanceCount=1000", 1000},
		{"InstanceCount=10000", 10000},
		{"InstanceCount=100000", 100000},
		{"InstanceCount=1000000", 1000000},
	}

	for _, tt := range tests {
		builder := balancers.NewRoundRobinBalancerBuilder[*TestInstance]()
		balancer := builder.Build()
		for i := 0; i < tt.instanceCount; i++ {
			balancer.RegisterInstance(NewTestInstance(fmt.Sprint(i)))
		}
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				balancer.Select()
			}
		})
	}
}
