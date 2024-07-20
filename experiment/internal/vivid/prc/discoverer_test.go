package prc_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/prc"
	"testing"
)

func TestDiscoverer(t *testing.T) {
	rc1 := prc.NewResourceController(prc.FunctionalResourceControllerConfigurator(func(config *prc.ResourceControllerConfiguration) {
		config.WithPhysicalAddress("127.0.0.1:8080")
	}))

	rc2 := prc.NewResourceController(prc.FunctionalResourceControllerConfigurator(func(config *prc.ResourceControllerConfiguration) {
		config.WithPhysicalAddress("127.0.0.1:8081")
	}))

	d1 := prc.NewDiscoverer(rc1, "127.0.0.1", 18080, prc.FunctionalDiscovererConfigurator(func(configuration *prc.DiscovererConfiguration) {
		configuration.WithJoinNodes("127.0.0.1:18080")
	}))

	d2 := prc.NewDiscoverer(rc2, "127.0.0.1", 18081, prc.FunctionalDiscovererConfigurator(func(configuration *prc.DiscovererConfiguration) {
		configuration.WithJoinNodes("127.0.0.1:18080", "127.0.0.1:18081")
	}))

	if err := d1.Discover(); err != nil {
		panic(err)
	}

	if err := d2.Discover(); err != nil {
		panic(err)
	}

	d1.Leave()
	d2.Leave()
}
