package notify

// Notify 通用通知接口定义
type Notify interface {
	// GetTitle 获取通知标题
	GetTitle() string
	// GetContent 获取通知内容
	GetContent() (string, error)
}
