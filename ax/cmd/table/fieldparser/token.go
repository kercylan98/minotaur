package fieldparser

type tokenType int

const (
	TokenEof tokenType = iota
	TokenLbrace
	TokenRbrace
	TokenLbracket
	TokenRbracket
	TokenColon
	TokenComma
	TokenIdent
)

type token struct {
	Type  tokenType
	Value string
}
