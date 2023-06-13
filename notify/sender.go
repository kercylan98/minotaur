package notify

// Sender 通知发送器接口声明
type Sender interface {
	// Push 推送通知
	Push(notify Notify) error
}
