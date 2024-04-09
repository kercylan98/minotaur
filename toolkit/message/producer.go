package message

type Producer[T comparable] interface {
	GetTopic() T
}
