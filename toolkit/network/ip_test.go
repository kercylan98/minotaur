package network_test

import (
	"github.com/kercylan98/minotaur/toolkit/network"
	"testing"
)

func TestIsSameLocalAddress(t *testing.T) {
	t.Log(network.IsSameLocalAddress(":8080", ":8080"))
}
