package parser

type Token string

func (t Token) EqualSymbol(symbol Symbol) bool {
	ts := Token(symbol)
	return t == ts
}

func (t Token) String() string {
	return string(t)
}

func (t Token) Empty() bool {
	return t.String() == ""
}
