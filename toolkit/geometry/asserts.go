package geometry

// AssertPolygonValid 断言检查多边形是否有效
func AssertPolygonValid(polygons ...Polygon) {
	for _, polygon := range polygons {
		AssertPointValid(polygon...)
		if len(polygon) < 3 {
			panic("polygon must have at least 3 points")
		}
	}
}

// AssertLineSegmentValid 断言检查线段是否有效
func AssertLineSegmentValid(lineSegments ...LineSegment) {
	for _, lineSegment := range lineSegments {
		AssertPointValid(lineSegment...)
		if len(lineSegment) < 2 {
			panic("line segment must have at least 2 points")
		}
	}
}

// AssertPointValid 断言检查点是否有效
func AssertPointValid(points ...Point) {
	for _, point := range points {
		if len(point) != 2 {
			panic("point must have 2 coordinates")
		}
	}
}

// AssertVector2Valid 断言检查二维向量是否有效
func AssertVector2Valid(vectors ...Vector2) {
	for _, vector := range vectors {
		if len(vector) != 2 {
			panic("vector must have 2 coordinates")
		}
	}
}

// AssertVector3Valid 断言检查三维向量是否有效
func AssertVector3Valid(vectors ...Vector3) {
	for _, vector := range vectors {
		if len(vector) != 3 {
			panic("vector must have 3 coordinates")
		}
	}
}
