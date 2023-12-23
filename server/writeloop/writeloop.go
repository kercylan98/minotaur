package writeloop

type WriteLoop[Message any] interface {
	Put(message Message)
	Close()
}
