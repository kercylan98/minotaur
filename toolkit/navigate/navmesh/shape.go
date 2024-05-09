package navmesh

import (
	"github.com/kercylan98/minotaur/toolkit/geometry"
)

func newShape(id int, s geometry.Polygon) *shape {
	return &shape{
		id:             id,
		Polygon:        s,
		centroid:       geometry.CalcRectangleVerticesCentroid(s),
		boundingRadius: s.CircumscribedCircleRadiusWithVerticesCentroid(),
		edges:          s.GetEdges(),
	}
}

type shape struct {
	id int
	geometry.Polygon
	links          []*shape
	portals        []geometry.LineSegment
	boundingRadius float64
	centroid       geometry.Point
	edges          []geometry.LineSegment

	weight float64
	x, y   float64
}

func (slf *shape) Id() int {
	return slf.id
}

func (slf *shape) Edges() []geometry.LineSegment {
	return slf.edges
}

func (slf *shape) BoundingRadius() float64 {
	return slf.boundingRadius
}

func (slf *shape) Centroid() geometry.Point {
	return slf.centroid
}

func (slf *shape) IsWall() bool {
	return slf.weight == 0
}

func (slf *shape) GetCost(point geometry.Point) float64 {
	return slf.Centroid().Distance2D(point)
}
