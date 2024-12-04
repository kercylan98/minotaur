package parser

type Handler[T any] interface {
	Handle(tokens *Tokens[T]) (T, error)
}

type FunctionalHandler[T any] func(tokens *Tokens[T]) (T, error)

func (f FunctionalHandler[T]) Handle(tokens *Tokens[T]) (T, error) {
	return f(tokens)
}
