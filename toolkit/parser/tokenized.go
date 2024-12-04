package parser

type Tokenized[T any] struct {
	tokens      []Token
	symbolTable map[Symbol]bool
}

func (t *Tokenized[T]) Parse(handler Handler[T]) (T, error) {
	tokens := &Tokens[T]{
		tokenized:   t,
		handler:     handler,
		symbolTable: t.symbolTable,
		pos:         0,
	}
	return handler.Handle(tokens)
}
