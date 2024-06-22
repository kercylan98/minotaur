package core

type Queue interface {
	Enqueue(message Message)
	Dequeue() Message
}
