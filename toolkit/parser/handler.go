package parser

type Handler interface {
	Handle(tokens *Tokens) error
}

type FunctionalHandler func(tokens *Tokens) error

func (f FunctionalHandler) Handle(tokens *Tokens) error {
	return f(tokens)
}
