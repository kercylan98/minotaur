package notify

// Notify 通用通知接口定义
type Notify interface {
	// Format 格式化通知内容
	Format() (string, error)
}
