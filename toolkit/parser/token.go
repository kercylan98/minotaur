package parser

type Token string

func (t Token) EqualSymbol(symbol Symbol) bool {
	return t == Token(symbol)
}

func (t Token) String() string {
	return string(t)
}
