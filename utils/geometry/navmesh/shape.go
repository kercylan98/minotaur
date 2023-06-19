package navmesh

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
)

func newShape[V generic.SignedNumber](s geometry.Shape[V]) *shape[V] {
	return &shape[V]{
		Shape: s,
	}
}

type shape[V generic.SignedNumber] struct {
	geometry.Shape[V]
	links   []*shape[V]
	portals []geometry.Line[V]
}
