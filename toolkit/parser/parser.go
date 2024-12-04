package parser

import (
	"errors"
	"strings"
	"unicode"
)

func New(symbols ...Symbol) *Parser {
	parser := &Parser{
		symbolTable: make(map[Symbol]bool),
	}
	for _, symbol := range symbols {
		if unicode.IsSpace(rune(symbol)) {
			panic(errors.New("symbol cannot be a space"))
		}
		parser.symbolTable[symbol] = true
	}
	return parser
}

type Parser struct {
	delimiter   Symbol          // 分隔符
	symbolTable map[Symbol]bool // 符号表
}

func (p *Parser) Tokenize(input string) *Tokenized {
	var pos int
	var tokens []Token
	var current strings.Builder

	for pos < len(input) {
		char := input[pos]
		switch {
		case unicode.IsSpace(rune(char)):
			// 始终将空格作为分隔标记，并将当前缓冲区中的内容添加到 tokens 中
			if current.Len() > 0 {
				tokens = append(tokens, Token(current.String()))
				current.Reset()
			}
		case p.symbolTable[char]:
			// 满足符号表条件，将当前缓冲区中的内容添加到 tokens 中，并将当前字符添加到 tokens 中
			if current.Len() > 0 {
				tokens = append(tokens, Token(current.String()))
				current.Reset()
			}
			tokens = append(tokens, Token(char))
		default:
			// 否则，将当前字符添加到当前缓冲区中
			current.WriteByte(char)
		}
		pos++
	}

	if current.Len() > 0 {
		tokens = append(tokens, Token(current.String()))
	}

	return &Tokenized{tokens: tokens}
}
