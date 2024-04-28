package navmesh

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
	"github.com/kercylan98/minotaur/utils/geometry/astar"
	"github.com/kercylan98/minotaur/utils/maths"
)

// NewNavMesh 创建一个新的导航网格，并返回一个指向该导航网格的指针。
//
// 参数：
//   - shapes: 形状切片，类型为 []geometry.Shape[V]，表示导航网格中的形状。
//   - meshShrinkAmount: 网格缩小量，类型为 V，表示导航网格的缩小量。
//
// 返回值：
//   - *NavMesh[V]: 指向创建的导航网格的指针。
//
// 注意事项：
//   - 导航网格的形状可以是任何几何形状。
//   - meshShrinkAmount 表示导航网格的缩小量，用于在形状之间创建链接时考虑形状的缩小效果。
//   - 函数内部使用了泛型类型参数 V，可以根据需要指定形状的坐标类型。
//   - 函数返回一个指向创建的导航网格的指针。
//
// 使用建议：
//   - 确保 NavMesh 计算精度的情况下，V 建议使用 float64 类型
func NewNavMesh[V generic.SignedNumber](shapes []geometry.Shape[V], meshShrinkAmount V) *NavMesh[V] {
	nm := &NavMesh[V]{
		meshShapes:       make([]*shape[V], len(shapes)),
		meshShrinkAmount: meshShrinkAmount,
	}
	for i, shape := range shapes {
		nm.meshShapes[i] = newShape(i, shape)
	}
	nm.generateLink()
	return nm
}

type NavMesh[V generic.SignedNumber] struct {
	meshShapes       []*shape[V]
	meshShrinkAmount V
}

// Neighbours 实现 astar.Graph 的接口，用于向 A* 算法提供相邻图形
func (slf *NavMesh[V]) Neighbours(node *shape[V]) []*shape[V] {
	return node.links
}

// Find 用于在 NavMesh 中查找离给定点最近的形状，并返回距离、找到的点和找到的形状。
//
// 参数：
//   - point: 给定的点，类型为 geometry.Point[V]，表示一个 V 维度的点坐标。
//   - maxDistance: 最大距离，类型为 V，表示查找的最大距离限制。
//
// 返回值：
//   - distance: 距离，类型为 V，表示离给定点最近的形状的距离。
//   - findPoint: 找到的点，类型为 geometry.Point[V]，表示离给定点最近的点坐标。
//   - findShape: 找到的形状，类型为 geometry.Shape[V]，表示离给定点最近的形状。
//
// 注意事项：
//   - 如果给定点在 NavMesh 中的某个形状内部或者在形状的边上，距离为 0，找到的形状为该形状，找到的点为给定点。
//   - 如果给定点不在任何形状内部或者形状的边上，将计算给定点到每个形状的距离，并找到最近的形状和对应的点。
//   - 距离的计算采用几何学中的投影点到形状的距离。
//   - 函数返回离给定点最近的形状的距离、找到的点和找到的形状。
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
		distance := geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(
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

// FindPath 函数用于在 NavMesh 中查找从起点到终点的路径，并返回路径上的点序列。
//
// 参数：
//   - start: 起点，类型为 geometry.Point[V]，表示路径的起始点。
//   - end: 终点，类型为 geometry.Point[V]，表示路径的终点。
//
// 返回值：
//   - result: 路径上的点序列，类型为 []geometry.Point[V]。
//
// 注意事项：
//   - 函数首先根据起点和终点的位置，找到离它们最近的形状作为起点形状和终点形状。
//   - 如果起点或终点不在任何形状内部，且 NavMesh 的 meshShrinkAmount 大于0，则会考虑缩小的形状。
//   - 使用 A* 算法在 NavMesh 上搜索从起点形状到终点形状的最短路径。
//   - 使用漏斗算法对路径进行优化，以得到最终的路径点序列。
func (slf *NavMesh[V]) FindPath(start, end geometry.Point[V]) (result []geometry.Point[V]) {
	var startShape, endShape *shape[V]
	var startDistance, endDistance = V(-1), V(-1)

	for _, meshShape := range slf.meshShapes {
		br := meshShape.BoundingRadius()

		distance := geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(meshShape.Centroid(), start))
		if (distance <= startDistance || startDistance == V(-1)) && distance <= br && meshShape.Contains(start) {
			startShape = meshShape
			startDistance = distance
		}

		distance = geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(meshShape.Centroid(), end))
		if (distance <= endDistance || endDistance == V(-1)) && distance <= br && meshShape.Contains(end) {
			endShape = meshShape
			endDistance = distance
		}
	}

	if endShape == nil && slf.meshShrinkAmount > V(0) {
		for _, meshShape := range slf.meshShapes {
			br := meshShape.BoundingRadius() + slf.meshShrinkAmount
			distance := geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(meshShape.Centroid(), end))
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
			distance := geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(meshShape.Centroid(), start))
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

	path := astar.Find[*shape[V], V](slf, startShape, endShape, func(a, b *shape[V]) V {
		return geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(a.centroid, b.centroid))
	}, func(a, b *shape[V]) V {
		return geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(a.centroid, b.centroid))
	})

	if len(path) == 0 {
		return
	}

	path = append([]*shape[V]{startShape}, path...)

	funnel := new(funnel[V])
	funnel.pushSingle(start)
	for i := 0; i < len(path)-1; i++ {
		current := path[i]
		next := path[i+1]
		if current.id == next.id {
			continue
		}

		var portal geometry.LineSegment[V]
		var find bool
		for i := 0; i < len(current.links); i++ {
			if current.links[i].id == next.id {
				portal = current.portals[i]
				find = true
			}
		}
		if !find {
			panic("not found portal")
		}

		funnel.push(portal.GetStart(), portal.GetEnd())
	}
	funnel.pushSingle(end)
	funnel.stringPull()

	var lastPoint geometry.Point[V]
	for i, point := range funnel.path {
		var np = point.Copy()
		if i == 0 || !np.Equal(lastPoint) {
			result = append(result, np)
		}
		lastPoint = np
	}
	return result
}

func (slf *NavMesh[V]) generateLink() {
	for i := 0; i < len(slf.meshShapes); i++ {
		shapePkg := slf.meshShapes[i]
		shapeCentroid := shapePkg.Centroid()
		shapeBoundingRadius := shapePkg.BoundingRadius()
		shapeEdges := shapePkg.Edges()
		for t := i + 1; t < len(slf.meshShapes); t++ {
			targetShapePkg := slf.meshShapes[t]
			targetShapeCentroid := targetShapePkg.Centroid()
			targetShapeBoundingRadius := targetShapePkg.BoundingRadius()
			centroidDistance := geometry.CalcDistanceWithCoordinate(geometry.DoublePointToCoordinate(shapeCentroid, targetShapeCentroid))
			if centroidDistance > shapeBoundingRadius+targetShapeBoundingRadius {
				continue
			}

			for _, shapeEdge := range shapeEdges {
				for _, targetEdge := range targetShapePkg.Edges() {
					if !geometry.CalcLineSegmentIsCollinear(shapeEdge, targetEdge, V(maths.GetDefaultTolerance())) {
						continue
					}

					var overlapLine, overlap = geometry.CalcLineSegmentIsOverlap(shapeEdge, targetEdge)
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
						shapePkg.portals = append(shapePkg.portals, geometry.NewLineSegment(overlapLine.GetStart(), overlapLine.GetEnd()))
					} else {
						shapePkg.portals = append(shapePkg.portals, geometry.NewLineSegment(overlapLine.GetEnd(), overlapLine.GetStart()))
					}

					edgeAngle = geometry.CalcAngle(geometry.DoublePointToCoordinate(targetShapeCentroid, targetEdge.GetStart()))
					a1 = geometry.CalcAngle(geometry.DoublePointToCoordinate(targetShapeCentroid, overlapLine.GetStart()))
					a2 = geometry.CalcAngle(geometry.DoublePointToCoordinate(targetShapeCentroid, overlapLine.GetEnd()))
					a3 = geometry.CalcAngleDifference(edgeAngle, a1)
					a4 = geometry.CalcAngleDifference(edgeAngle, a2)
					if a3 < a4 {
						targetShapePkg.portals = append(targetShapePkg.portals, geometry.NewLineSegment(overlapLine.GetStart(), overlapLine.GetEnd()))
					} else {
						targetShapePkg.portals = append(targetShapePkg.portals, geometry.NewLineSegment(overlapLine.GetEnd(), overlapLine.GetStart()))
					}
				}
			}

		}
	}
}
