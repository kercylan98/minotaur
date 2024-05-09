package geometry

import "math"

// NewPolygon 创建一个多边形
func NewPolygon(points ...Point) Polygon {
	AssertPolygonValid(points)
	return points
}

// Polygon 由至少三个点组成的多边形
type Polygon []Point

// GetEdges 获取多边形的边
func (p Polygon) GetEdges() []LineSegment {
	AssertPolygonValid(p)
	var edges []LineSegment
	for i := 0; i < len(p); i++ {
		edges = append(edges, NewLineSegment(p[i], p[(i+1)%len(p)]))
	}
	return edges
}

// IsPointInside 判断点是否在多边形内
func (p Polygon) IsPointInside(point Point) bool {
	x, y := point.GetXY()
	inside := false
	for i, j := 0, len(p)-1; i < len(p); i, j = i+1, i {
		ix := p[i].GetX()
		iy := p[i].GetY()
		jx := p[j].GetX()
		jy := p[j].GetY()

		if ((iy <= y && y < jy) || (jy <= y && y < iy)) && x < ((jx-ix)*(y-iy))/(jy-iy)+ix {
			inside = !inside
		}
	}
	return inside
}

// IsPointOnEdge 判断点是否在多边形的边上
func (p Polygon) IsPointOnEdge(point Point) bool {
	AssertPolygonValid(p)
	for _, edge := range p.GetEdges() {
		if edge.IsPointOnSegment(point) {
			return true
		}
	}
	return false
}

// CircumscribedCircleCenter 计算多边形外接圆的圆心
func (p Polygon) CircumscribedCircleCenter() Point {
	return CalcPolygonCircumscribedCircleCenter(p)
}

// CircumscribedCircleRadius 计算多边形外接圆的半径
func (p Polygon) CircumscribedCircleRadius() float64 {
	return CalcPolygonCircumscribedCircleRadius(p)
}

// CircumscribedCircleRadiusWithVerticesCentroid 基于多边形的顶点平均值计算的质心计算多边形外接圆的半径
func (p Polygon) CircumscribedCircleRadiusWithVerticesCentroid() float64 {
	var boundingRadius float64
	var centroid = CalcRectangleVerticesCentroid(p)
	for _, point := range p {
		distance := centroid.Distance2D(point)
		if distance > boundingRadius {
			boundingRadius = distance
		}
	}
	return boundingRadius
}

// VerticesCentroid 基于多边形的顶点的平均值计算质心
func (p Polygon) VerticesCentroid() Point {
	return CalcPolygonVerticesCentroid(p)
}

// Centroid 计算多边形质心
func (p Polygon) Centroid() Point {
	return CalcPolygonCentroid(p)
}

// CalcRectangleVerticesCentroid 基于矩形的顶点的平均值计算质心
func CalcRectangleVerticesCentroid(rectangle Polygon) Point {
	var x, y float64
	length := float64(len(rectangle))
	for _, point := range rectangle {
		x += point.GetX()
		y += point.GetY()
	}
	x /= length
	y /= length
	return NewPoint(x, x)
}

// CalcPolygonVerticesCentroid 基于多边形的顶点的平均值计算质心
func CalcPolygonVerticesCentroid(polygon Polygon) Point {
	var centroid = NewPoint(0, 0)
	for _, point := range polygon {
		centroid = centroid.Add(point)
	}
	centroid = centroid.Div(float64(len(polygon)))
	return centroid
}

// CalcPolygonCentroid 计算多边形质心
func CalcPolygonCentroid(polygon Polygon) Point {
	var area float64
	var centroid = NewPoint(0, 0)
	for i := 0; i < len(polygon); i++ {
		j := (i + 1) % len(polygon)
		area += polygon[i][0]*polygon[j][1] - polygon[j][0]*polygon[i][1]
		centroid = centroid.Add(polygon[i].Add(polygon[j]).Mul(polygon[i][0]*polygon[j][1] - polygon[j][0]*polygon[i][1]))
	}
	area /= 2
	centroid = centroid.Div(6 * area)
	return centroid
}

// CalcPolygonCircumscribedCircleCenter 计算多边形外接圆的圆心
func CalcPolygonCircumscribedCircleCenter(polygon Polygon) Point {
	// 计算多边形的中心
	var center = NewPoint(0, 0)
	for _, point := range polygon {
		center = center.Add(point)
	}
	center = center.Div(float64(len(polygon)))

	// 计算多边形的外接圆半径
	var radius float64
	for _, point := range polygon {
		radius = math.Max(radius, point.Distance2D(center))
	}

	return center
}

// CalcPolygonCircumscribedCircleRadius 计算多边形外接圆的半径
func CalcPolygonCircumscribedCircleRadius(polygon Polygon) float64 {
	center := CalcPolygonCircumscribedCircleCenter(polygon)
	var radius float64
	for _, point := range polygon {
		radius = math.Max(radius, point.Distance2D(center))
	}
	return radius
}

// CalcPolygonPointProjection 给定一个点和一个多边形，计算多边形边界上与该点距离最短的点，并返回投影点和距离
func CalcPolygonPointProjection(polygon Polygon, point Point) (projection Point, distance float64) {
	var closestProjection Point
	var hasClosestProjection bool
	var closestDistance float64
	for _, edge := range polygon.GetEdges() {
		projectedPoint := edge.ClosestPoint(point)
		distance := point.Distance2D(projectedPoint)
		if !hasClosestProjection || distance < closestDistance {
			closestDistance = distance
			closestProjection = projectedPoint
			hasClosestProjection = true
		}
	}

	return closestProjection, closestDistance
}
