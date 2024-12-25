package cluster

import "github.com/kercylan98/minotaur/engine/prc"

type SeedProvider interface {
	// Provide 获取种子节点
	Provide() ([]prc.PhysicalAddress, error)
}

type FunctionalSeedProvider func() ([]prc.PhysicalAddress, error)

func (f FunctionalSeedProvider) Provide() ([]prc.PhysicalAddress, error) {
	return f()
}
