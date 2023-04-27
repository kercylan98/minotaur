package notify

type Sender interface {
	Push(notify Notify) error
}
