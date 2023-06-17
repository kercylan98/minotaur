package geometry

type shapeSearchOptions struct {
	lowerLimit    int
	upperLimit    int
	sort          int
	deduplication bool
}

// ShapeSearchOption 图形搜索可选项，用于 Shape.ShapeSearch 搜索支持
type ShapeSearchOption func(options *shapeSearchOptions)

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
