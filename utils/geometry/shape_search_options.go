package geometry

type shapeSearchOptions struct {
	lowerLimit          int
	upperLimit          int
	sort                int
	deduplication       bool
	directionCountLower map[Direction]int
	directionCountUpper map[Direction]int
	directionCount      int
	oppositionDirection Direction
	ra                  bool
}

// ShapeSearchOption 图形搜索可选项，用于 Shape.ShapeSearch 搜索支持
type ShapeSearchOption func(options *shapeSearchOptions)

// WithShapeSearchRightAngle 通过直角的方式进行搜索
func WithShapeSearchRightAngle() ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		options.ra = true
	}
}

// WithShapeSearchOppositionDirection 通过限制对立方向的方式搜索
//   - 对立方向例如上不能与下共存
func WithShapeSearchOppositionDirection(direction Direction) ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		options.oppositionDirection = direction
	}
}

// WithShapeSearchDirectionCount 通过限制方向数量的方式搜索
func WithShapeSearchDirectionCount(count int) ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		options.directionCount = count
	}
}

// WithShapeSearchDirectionCountLowerLimit 通过限制特定方向数量下限的方式搜索
func WithShapeSearchDirectionCountLowerLimit(direction Direction, count int) ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		if options.directionCountLower == nil {
			options.directionCountLower = map[Direction]int{}
		}
		options.directionCountLower[direction] = count
	}
}

// WithShapeSearchDirectionCountUpperLimit 通过限制特定方向数量上限的方式搜索
func WithShapeSearchDirectionCountUpperLimit(direction Direction, count int) ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		options.directionCountUpper[direction] = count
	}
}

// WithShapeSearchDeduplication 通过去重的方式进行搜索
//   - 去重方式中每个点仅会被使用一次
func WithShapeSearchDeduplication() ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		options.deduplication = true
	}
}

// WithShapeSearchPointCountLowerLimit 通过限制图形构成的最小点数进行搜索
//   - 当搜索到的图形的点数量低于 lowerLimit 时，将被忽略
func WithShapeSearchPointCountLowerLimit(lowerLimit int) ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		options.lowerLimit = lowerLimit
	}
}

// WithShapeSearchPointCountUpperLimit 通过限制图形构成的最大点数进行搜索
//   - 当搜索到的图形的点数量大于 upperLimit 时，将被忽略
func WithShapeSearchPointCountUpperLimit(upperLimit int) ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		options.upperLimit = upperLimit
	}
}

// WithShapeSearchAsc 通过升序的方式进行搜索
func WithShapeSearchAsc() ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		options.sort = 1
	}
}

// WithShapeSearchDesc 通过降序的方式进行搜索
func WithShapeSearchDesc() ShapeSearchOption {
	return func(options *shapeSearchOptions) {
		options.sort = -1
	}
}
