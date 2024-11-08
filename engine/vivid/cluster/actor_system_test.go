package cluster_test

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid/cluster"
	"testing"
	"time"
)

func TestActorSystem(t *testing.T) {
	system1 := cluster.NewActorSystem("127.0.0.1:6666", []prc.PhysicalAddress{"127.0.0.1:6666"})
	system2 := cluster.NewActorSystem("127.0.0.1:6667", []prc.PhysicalAddress{"127.0.0.1:6666"})
	system3 := cluster.NewActorSystem("127.0.0.1:6668", []prc.PhysicalAddress{"127.0.0.1:6666"})

	_ = system1
	_ = system2
	_ = system3

	time.Sleep(time.Second * 3)

	system3.Shutdown(true)

	time.Sleep(time.Hour)
}
