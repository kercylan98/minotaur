package navmesh

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
	"github.com/kercylan98/minotaur/utils/maths"
)

func NewNavMesh[V generic.SignedNumber](shapes []geometry.Shape[V], meshShrinkAmount V) *NavMesh[V] {
	nm := &NavMesh[V]{
		meshShapes:       make([]*shape[V], len(shapes)),
		meshShrinkAmount: meshShrinkAmount,
	}
	for i, shape := range shapes {
		nm.meshShapes[i] = newShape(shape)
	}
	nm.generateLink()
	return nm
}

type NavMesh[V generic.SignedNumber] struct {
	meshShapes       []*shape[V]
	meshShrinkAmount V
}

// Find 在网格中找到与给定点最近的点。如果该点已经在网格中，这将为您提供该点。如果点在网格之外，这将尝试将此点投影到网格中（直到给定的 maxDistance）
func (slf *NavMesh[V]) Find(point geometry.Point[V], maxDistance V) (distance V, findPoint geometry.Point[V], findShape geometry.Shape[V]) {
	var minDistance = maxDistance
	var closest *shape[V]
	var pointOnClosest geometry.Point[V]
	for _, meshShape := range slf.meshShapes {
		if meshShape.Contains(point) || geometry.IsPointOnEdge(meshShape.Edges(), point) {
			minDistance = 0
			closest = meshShape
			pointOnClosest = point
			break
		}
		br := geometry.CalcBoundingRadius(meshShape.Shape)
		distance := geometry.CalcDistance(geometry.DoublePointToCoordinate(
			geometry.CalcRectangleCentroid(meshShape.Shape),
			point,
		))
		if distance-br < minDistance {
			point, distance := geometry.ProjectionPointToShape(point, meshShape.Shape)
			if distance < minDistance {
				minDistance = distance
				closest = meshShape
				pointOnClosest = point
			}
		}
	}
	return minDistance, pointOnClosest, closest.Shape
}

// FindPath 使用此导航网格查找从起点到终点的路径。
func (slf *NavMesh[V]) FindPath(start, end geometry.Point[V]) (result []geometry.Point[V]) {
	var startShape, endShape *shape[V]
	var startDistance, endDistance = V(-1), V(-1)

	for _, meshShape := range slf.meshShapes {
		br := meshShape.BoundingRadius()

		distance := geometry.CalcDistance(geometry.DoublePointToCoordinate(meshShape.Centroid(), start))
		if (distance <= startDistance || startDistance == V(-1)) && distance <= br && meshShape.Contains(start) {
			startShape = meshShape
			startDistance = distance
		}

		distance = geometry.CalcDistance(geometry.DoublePointToCoordinate(meshShape.Centroid(), end))
		if (distance <= endDistance || endDistance == V(-1)) && distance <= br && meshShape.Contains(end) {
			endShape = meshShape
			endDistance = distance
		}
	}

	if endShape == nil && slf.meshShrinkAmount > V(0) {
		for _, meshShape := range slf.meshShapes {
			br := meshShape.BoundingRadius() + slf.meshShrinkAmount
			distance := geometry.CalcDistance(geometry.DoublePointToCoordinate(meshShape.Centroid(), end))
			if distance <= br {
				_, projectionDistance := geometry.ProjectionPointToShape(end, meshShape.Shape)
				if projectionDistance <= slf.meshShrinkAmount && projectionDistance < endDistance {
					endShape = meshShape
					endDistance = projectionDistance
				}
			}
		}
	}

	if endShape == nil {
		return
	}

	if startShape == nil && slf.meshShrinkAmount > 0 {
		for _, meshShape := range slf.meshShapes {
			br := meshShape.BoundingRadius() + slf.meshShrinkAmount
			distance := geometry.CalcDistance(geometry.DoublePointToCoordinate(meshShape.Centroid(), start))
			if distance <= br {
				_, projectionDistance := geometry.ProjectionPointToShape(start, meshShape.Shape)
				if projectionDistance <= slf.meshShrinkAmount && projectionDistance < startDistance {
					startShape = meshShape
					startDistance = projectionDistance
				}
			}
		}
	}

	if startShape == nil {
		return
	}

	if startShape == endShape {
		return append(result, start, end)
	}

}

func (slf *NavMesh[V]) aStar() {

}

func (slf *NavMesh[V]) generateLink() {
	refer := len(slf.meshShapes)
	for i := 0; i < refer; i++ {
		shapePkg := slf.meshShapes[i]
		shapeCentroid := shapePkg.Centroid()
		shapeBoundingRadius := shapePkg.BoundingRadius()
		shapeEdges := shapePkg.Edges()
		for t := i + 1; t < len(slf.meshShapes); t++ {
			targetShapePkg := slf.meshShapes[t]
			targetShapeCentroid := targetShapePkg.Centroid()
			targetShapeBoundingRadius := targetShapePkg.BoundingRadius()
			centroidDistance := geometry.CalcDistance(geometry.DoublePointToCoordinate(shapeCentroid, targetShapeCentroid))
			if centroidDistance > shapeBoundingRadius+targetShapeBoundingRadius {
				continue
			}

			for _, shapeEdge := range shapeEdges {
				for _, targetEdge := range targetShapePkg.Edges() {
					if !geometry.CalcLineIsCollinear(shapeEdge, targetEdge, V(maths.GetDefaultTolerance())) {
						continue
					}

					var overlapLine, overlap = geometry.CalcLineIsOverlap(shapeEdge, targetEdge)
					if !overlap {
						continue
					}

					shapePkg.links = append(shapePkg.links, targetShapePkg)
					targetShapePkg.links = append(targetShapePkg.links, shapePkg)

					edgeAngle := geometry.CalcAngle(geometry.DoublePointToCoordinate(shapeCentroid, shapeEdge.GetStart()))
					a1 := geometry.CalcAngle(geometry.DoublePointToCoordinate(shapeCentroid, overlapLine.GetStart()))
					a2 := geometry.CalcAngle(geometry.DoublePointToCoordinate(shapeCentroid, overlapLine.GetEnd()))
					a3 := geometry.CalcAngleDifference(edgeAngle, a1)
					a4 := geometry.CalcAngleDifference(edgeAngle, a2)
					if a3 < a4 {
						shapePkg.portals = append(shapePkg.portals, geometry.NewLine(overlapLine.GetStart(), overlapLine.GetEnd()))
					} else {
						shapePkg.portals = append(shapePkg.portals, geometry.NewLine(overlapLine.GetEnd(), overlapLine.GetStart()))
					}

					edgeAngle = geometry.CalcAngle(geometry.DoublePointToCoordinate(targetShapeCentroid, targetEdge.GetStart()))
					a1 = geometry.CalcAngle(geometry.DoublePointToCoordinate(targetShapeCentroid, overlapLine.GetStart()))
					a2 = geometry.CalcAngle(geometry.DoublePointToCoordinate(targetShapeCentroid, overlapLine.GetEnd()))
					a3 = geometry.CalcAngleDifference(edgeAngle, a1)
					a4 = geometry.CalcAngleDifference(edgeAngle, a2)
					if a3 < a4 {
						targetShapePkg.portals = append(targetShapePkg.portals, geometry.NewLine(overlapLine.GetStart(), overlapLine.GetEnd()))
					} else {
						targetShapePkg.portals = append(targetShapePkg.portals, geometry.NewLine(overlapLine.GetEnd(), overlapLine.GetStart()))
					}
				}
			}

		}
	}
}
