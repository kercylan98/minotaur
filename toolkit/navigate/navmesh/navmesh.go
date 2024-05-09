package navmesh

import (
	"github.com/kercylan98/minotaur/toolkit/geometry"
	"github.com/kercylan98/minotaur/toolkit/navigate/astar"
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
func NewNavMesh(shapes []geometry.Polygon, meshShrinkAmount float64) *NavMesh {
	nm := &NavMesh{
		meshShapes:       make([]*shape, len(shapes)),
		meshShrinkAmount: meshShrinkAmount,
	}
	for i, shape := range shapes {
		nm.meshShapes[i] = newShape(i, shape)
	}
	nm.generateLink()
	return nm
}

type NavMesh struct {
	meshShapes       []*shape
	meshShrinkAmount float64
}

// GetNodeId 实现 astar.Graph 的接口，用于返回给定形状的唯一标识。
func (m *NavMesh) GetNodeId(node *shape) int {
	return node.id
}

// GetNeighbours 实现 astar.Graph 的接口，用于向 A* 算法提供相邻图形
func (m *NavMesh) GetNeighbours(node *shape) []*shape {
	return node.links
}

// Find 用于在 NavMesh 中查找离给定点最近的形状，并返回距离、找到的点和找到的形状。
//
// 参数：
//   - point: 给定的点，类型为 geometry.Point，表示一个 V 维度的点坐标。
//   - maxDistance: 最大距离，类型为 V，表示查找的最大距离限制。
//
// 返回值：当 distance 为 -1 时表示未找到最近的形状，否则返回距离、找到的点和找到的形状。
//   - distance: 距离，类型为 float64，表示离给定点最近的形状的距离。
//   - findPoint: 找到的点，类型为 geometry.Point，表示离给定点最近的点坐标。
//   - findPolygon: 找到的多边形，类型为 geometry.Polygon，表示离给定点最近的形状。
//
// 注意事项：
//   - 如果给定点在 NavMesh 中的某个形状内部或者在形状的边上，距离为 0，找到的形状为该形状，找到的点为给定点。
//   - 如果给定点不在任何形状内部或者形状的边上，将计算给定点到每个形状的距离，并找到最近的形状和对应的点。
//   - 距离的计算采用几何学中的投影点到形状的距离。
//   - 函数返回离给定点最近的形状的距离、找到的点和找到的形状。
func (m *NavMesh) Find(point geometry.Point, maxDistance float64) (distance float64, findPoint geometry.Point, findPolygon geometry.Polygon) {
	var minDistance = maxDistance
	var closest *shape
	var pointOnClosest geometry.Point
	for _, meshShape := range m.meshShapes {
		if meshShape.IsPointInside(point) || meshShape.IsPointOnEdge(point) {
			minDistance = 0
			closest = meshShape
			pointOnClosest = point
			break
		}
		br := meshShape.CircumscribedCircleRadius()
		distance := meshShape.VerticesCentroid().Distance2D(point)
		if distance-br < minDistance {
			point, distance := geometry.CalcPolygonPointProjection(meshShape.Polygon, point)
			if distance < minDistance {
				minDistance = distance
				closest = meshShape
				pointOnClosest = point
			}
		}
	}
	if closest == nil {
		return -1, nil, nil
	}
	return minDistance, pointOnClosest, closest.Polygon
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
func (m *NavMesh) FindPath(start, end geometry.Point) (result []geometry.Point) {
	var startShape, endShape *shape
	var startDistance, endDistance = -1.0, -1.0

	for _, meshShape := range m.meshShapes {
		br := meshShape.boundingRadius

		distance := meshShape.centroid.Distance2D(start)
		if (distance <= startDistance || startDistance == -1) && distance <= br && meshShape.IsPointInside(start) {
			startShape = meshShape
			startDistance = distance
		}

		distance = meshShape.centroid.Distance2D(end)
		if (distance <= endDistance || endDistance == -1) && distance <= br && meshShape.IsPointInside(end) {
			endShape = meshShape
			endDistance = distance
		}
	}

	if endShape == nil && m.meshShrinkAmount > 0 {
		for _, meshShape := range m.meshShapes {
			br := meshShape.boundingRadius + m.meshShrinkAmount
			distance := meshShape.centroid.Distance2D(end)
			if distance <= br {
				_, projectionDistance := geometry.CalcPolygonPointProjection(meshShape.Polygon, end)
				if projectionDistance <= m.meshShrinkAmount && projectionDistance < endDistance {
					endShape = meshShape
					endDistance = projectionDistance
				}
			}
		}
	}

	if endShape == nil {
		return
	}

	if startShape == nil && m.meshShrinkAmount > 0 {
		for _, meshShape := range m.meshShapes {
			br := meshShape.BoundingRadius() + m.meshShrinkAmount
			distance := meshShape.centroid.Distance2D(start)
			if distance <= br {
				_, projectionDistance := geometry.CalcPolygonPointProjection(meshShape.Polygon, start)
				if projectionDistance <= m.meshShrinkAmount && projectionDistance < startDistance {
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

	path := astar.Find[int, *shape](m, startShape, endShape,
		func(a, b *shape) float64 {
			return a.VerticesCentroid().Distance2D(b.VerticesCentroid())
		},
		func(a, b *shape) float64 {
			return a.VerticesCentroid().Distance2D(b.VerticesCentroid())
		},
	)

	if len(path) == 0 {
		return
	}

	path = append([]*shape{startShape}, path...)

	f := new(funnel)
	f.pushSingle(start)
	for i := 0; i < len(path)-1; i++ {
		current := path[i]
		next := path[i+1]
		if current.id == next.id {
			continue
		}

		var portal geometry.LineSegment
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

		f.push(portal[0], portal[1])
	}
	f.pushSingle(end)
	f.stringPull()

	var lastPoint geometry.Point
	for i, point := range f.path {
		var np = point.Clone()
		if i == 0 || !np.Equal(lastPoint) {
			result = append(result, np)
		}
		lastPoint = np
	}
	return result
}

func (m *NavMesh) generateLink() {
	for i := 0; i < len(m.meshShapes); i++ {
		shapePkg := m.meshShapes[i]
		shapeCentroid := shapePkg.centroid
		shapeBoundingRadius := shapePkg.boundingRadius
		shapeEdges := shapePkg.Edges()
		for t := i + 1; t < len(m.meshShapes); t++ {
			targetShapePkg := m.meshShapes[t]
			targetShapeCentroid := targetShapePkg.centroid
			targetShapeBoundingRadius := targetShapePkg.boundingRadius
			centroidDistance := shapeCentroid.Distance2D(targetShapeCentroid)
			if centroidDistance > shapeBoundingRadius+targetShapeBoundingRadius {
				continue
			}

			for _, shapeEdge := range shapeEdges {
				for _, targetEdge := range targetShapePkg.Edges() {
					if !geometry.CalcLineSegmentCollinearWithEpsilon(shapeEdge, targetEdge, 1e-4) {
						continue
					}

					var overlapLine, overlap = geometry.CalcLineSegmentOverlap(shapeEdge, targetEdge)
					if !overlap {
						continue
					}

					shapePkg.links = append(shapePkg.links, targetShapePkg)
					targetShapePkg.links = append(targetShapePkg.links, shapePkg)

					edgeAngle := shapeCentroid.PolarAngle(shapeEdge[0])
					a1 := shapeCentroid.PolarAngle(overlapLine[0])
					a2 := shapeCentroid.PolarAngle(overlapLine[1])
					a3 := geometry.CalcAngleDifference(edgeAngle, a1)
					a4 := geometry.CalcAngleDifference(edgeAngle, a2)
					if a3 < a4 {
						shapePkg.portals = append(shapePkg.portals, geometry.NewLineSegment(overlapLine[0], overlapLine[1]))
					} else {
						shapePkg.portals = append(shapePkg.portals, geometry.NewLineSegment(overlapLine[1], overlapLine[0]))
					}

					edgeAngle = targetShapeCentroid.PolarAngle(targetEdge[0])
					a1 = targetShapeCentroid.PolarAngle(overlapLine[0])
					a2 = targetShapeCentroid.PolarAngle(overlapLine[1])
					a3 = geometry.CalcAngleDifference(edgeAngle, a1)
					a4 = geometry.CalcAngleDifference(edgeAngle, a2)
					if a3 < a4 {
						targetShapePkg.portals = append(targetShapePkg.portals, geometry.NewLineSegment(overlapLine[0], overlapLine[1]))
					} else {
						targetShapePkg.portals = append(targetShapePkg.portals, geometry.NewLineSegment(overlapLine[1], overlapLine[0]))
					}
				}
			}

		}
	}
}
