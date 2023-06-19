package navmesh

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
)

func newShape[V generic.SignedNumber](s geometry.Shape[V]) *shape[V] {
	return &shape[V]{
		Shape:          s,
		centroid:       geometry.CalcRectangleCentroid(s),
		boundingRadius: geometry.CalcBoundingRadius(s),
		edges:          s.Edges(),
	}
}

type shape[V generic.SignedNumber] struct {
	geometry.Shape[V]
	links          []*shape[V]
	portals        []geometry.Line[V]
	boundingRadius V
	centroid       geometry.Point[V]
	edges          []geometry.Line[V]

	weight V
	x, y   V
}

func (slf *shape[V]) Edges() []geometry.Line[V] {
	return slf.edges
}

func (slf *shape[V]) BoundingRadius() V {
	return slf.boundingRadius
}

func (slf *shape[V]) Centroid() geometry.Point[V] {
	return slf.centroid
}

func (slf *shape[V]) IsWall() bool {
	return slf.weight == 0
}

func (slf *shape[V]) GetCost(point geometry.Point[V]) V {
	return geometry.CalcDistance(geometry.DoublePointToCoordinate(slf.Centroid(), point))
}
