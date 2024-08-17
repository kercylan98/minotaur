package fieldparser

import (
	"fmt"
	"unicode"
)

type lexer struct {
	input string
	pos   int
}

func newLexer(input string) *lexer {
	return &lexer{input: input}
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	ch := rune(l.input[l.pos])
	l.pos++
	return ch
}

func (l *lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return rune(l.input[l.pos])
}

func (l *lexer) lex() token {
	for unicode.IsSpace(l.peek()) {
		l.next()
	}
	switch ch := l.next(); ch {
	case 0:
		return token{Type: TokenEof}
	case '{':
		return token{Type: TokenLbrace, Value: "{"}
	case '}':
		return token{Type: TokenRbrace, Value: "}"}
	case '[':
		return token{Type: TokenLbracket, Value: "["}
	case ']':
		return token{Type: TokenRbracket, Value: "]"}
	case ':':
		return token{Type: TokenColon, Value: ":"}
	case ',':
		return token{Type: TokenComma, Value: ","}
	default:
		if unicode.IsLetter(ch) {
			start := l.pos - 1
			for unicode.IsLetter(l.peek()) {
				l.next()
			}
			return token{Type: TokenIdent, Value: l.input[start:l.pos]}
		}
		panic(fmt.Errorf("%s unexpected character: %q", l.input, ch))
	}
}
