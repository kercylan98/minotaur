package survey

// Option 选项
type Option func(logger *logger)

// WithLayout 设置日志文件名的时间戳格式
//   - 默认为 time.DateOnly
func WithLayout(layout string) Option {
	return func(logger *logger) {
		logger.layout = layout
		logger.layoutLen = len(layout)
	}
}
