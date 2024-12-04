package parser_test

import (
	"errors"
	"github.com/kercylan98/minotaur/toolkit/parser"
	"testing"
)

func TestParser(t *testing.T) {
	p := parser.New(parser.SymbolAt, parser.SymbolLeftParen, parser.SymbolRightParen)
	var tests = []struct {
		Annotation string
	}{
		{Annotation: "@Annotation()"},
		{Annotation: "@Annotation(a,b,c)"},
	}

	for _, test := range tests {
		t.Run(test.Annotation, func(t *testing.T) {
			tokenized := p.Tokenize(test.Annotation)
			err := tokenized.Parse(parser.FunctionalHandler(func(tokens *parser.Tokens) error {
				token := tokens.Consume()

				if token.EqualSymbol(parser.SymbolAt) {
					name := tokens.Consume()
					if len(name) == 0 {
						return errors.New("name is empty")
					}

					if !tokens.Consume().EqualSymbol(parser.SymbolLeftParen) {
						return errors.New("expect (")
					}
					var params string
					token = tokens.Consume()
					if !token.EqualSymbol(parser.SymbolRightParen) {
						params = token.String()
						token = tokens.Consume()
					}
					if !token.EqualSymbol(parser.SymbolRightParen) {
						return errors.New("expect )")
					}

					t.Log(name, params)
					return nil
				}

				return errors.New("expect @")
			}))
			if err != nil {
				panic(err)
			}
		})
	}
}
