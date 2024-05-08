package navmesh

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
)

func newShape[V generic.SignedNumber](id int, s geometry.Shape[V]) *shape[V] {
	return &shape[V]{
		id:             id,
		Shape:          s,
		centroid:       geometry.CalcRectangleCentroid(s),
		boundingRadius: geometry.CalcBoundingRadius(s),
		edges:          s.Edges(),
	}
}

type shape[V generic.SignedNumber] struct {
	id int
	geometry.Shape[V]
	links          []*shape[V]
	portals        []geometry.LineSegment[V]
	boundingRadius V
	centroid       geometry.Point[V]
	edges          []geometry.LineSegment[V]

	weight V
	x, y   V
}

func (slf *shape[V]) Edges() []geometry.LineSegment[V] {
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
	return geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(slf.Centroid(), point))
}
