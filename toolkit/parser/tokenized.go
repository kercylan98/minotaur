package parser

type Tokenized struct {
	tokens []Token
}

func (t *Tokenized) Parse(handler Handler) error {
	tokens := &Tokens{
		tokenized: t,
		handler:   handler,
		pos:       0,
	}
	return handler.Handle(tokens)
}
