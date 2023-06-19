package navmesh

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
	"github.com/kercylan98/minotaur/utils/maths"
)

type NavMesh[V generic.SignedNumber] struct {
	meshShapes []*shape[V]
}

func (slf *NavMesh[V]) generateLink() {
	refer := len(slf.meshShapes)
	for i := 0; i < refer; i++ {
		shapePkg := slf.meshShapes[i]
		shape := shapePkg.Shape
		shapeCentroid := geometry.CalcRectangleCentroid(shape)
		shapeBoundingRadius := geometry.CalcBoundingRadiusWithCentroid(shape, shapeCentroid)
		shapeEdges := shape.Edges()
		for t := i + 1; t < len(slf.meshShapes); t++ {
			targetShapePkg := slf.meshShapes[t]
			targetShape := targetShapePkg.Shape
			targetShapeCentroid := geometry.CalcRectangleCentroid(targetShape)
			targetShapeBoundingRadius := geometry.CalcBoundingRadiusWithCentroid(shape, targetShapeCentroid)
			centroidDistance := geometry.CalcDistance(geometry.DoublePointToCoordinate(shapeCentroid, targetShapeCentroid))
			if centroidDistance > shapeBoundingRadius+targetShapeBoundingRadius {
				continue
			}

			for _, shapeEdge := range shapeEdges {
				for _, targetEdge := range targetShape.Edges() {
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
