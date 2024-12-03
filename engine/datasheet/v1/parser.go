package datasheetv1

import (
	"fmt"
	"strings"
	"unicode"
)

type Type struct {
	Name  string
	Value interface{}
}

type Array struct {
	Size int
	Elem *Type
}

type Map struct {
	Key   *Type
	Value *Type
}

type Struct struct {
	Fields map[string]*Type
}

type Parser struct {
	input  string
	pos    int
	tokens []string
}

func newParser(input string) *Parser {
	p := &Parser{input: input}
	p.tokenize()
	return p
}

func (p *Parser) tokenize() {
	var tokens []string
	var current strings.Builder

	for p.pos < len(p.input) {
		ch := p.input[p.pos]
		switch {
		case unicode.IsSpace(rune(ch)):
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
		case ch == '{', ch == '}', ch == '[', ch == ']', ch == ':', ch == ',', ch == '=':
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(ch))
		default:
			current.WriteByte(ch)
		}
		p.pos++
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	p.tokens = tokens
	p.pos = 0
}

func (p *Parser) consume() string {
	if p.pos >= len(p.tokens) {
		return ""
	}
	token := p.tokens[p.pos]
	p.pos++
	return token
}

func (p *Parser) peek() string {
	if p.pos >= len(p.tokens) {
		return ""
	}
	return p.tokens[p.pos]
}

func (p *Parser) parseType() (*Type, error) {
	token := p.consume()

	toLower := strings.ToLower(token)
	switch toLower {
	case "boolean", "bool", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64",
		"float", "double", "byte", "time", "float32", "float64", "short", "long", "string", "bigint", "bigfloat":
		return &Type{Name: toLower}, nil
	}

	if token == "[" {
		sizeToken := p.consume()
		var size = -1
		var name = "slice"
		if sizeToken != "]" {
			name = "array"
			if _, err := fmt.Sscanf(sizeToken, "%d", &size); err != nil {
				panic(fmt.Sprintf("invalid array size: %s", sizeToken))
			}
		} else {
			p.pos--
		}
		p.consume()

		elemType, err := p.parseType()
		if err != nil {
			return nil, err
		}
		return &Type{Name: name, Value: &Array{Size: size, Elem: elemType}}, nil
	}

	if token == "map" {
		p.consume()
		keyType, err := p.parseType()
		if err != nil {
			return nil, err
		}
		p.consume()
		valType, err := p.parseType()
		if err != nil {
			return nil, err
		}
		return &Type{Name: "map", Value: &Map{Key: keyType, Value: valType}}, nil
	}

	if token == "{" {
		fields := make(map[string]*Type)
		for {
			fieldName := p.consume()
			if fieldName == "}" {
				break
			}
			p.consume()
			fieldType, err := p.parseType()
			if err != nil {
				return nil, err
			}
			fields[fieldName] = fieldType
			if p.peek() == "," {
				p.consume() // Consume the comma
			}
		}
		return &Type{Name: "struct", Value: &Struct{Fields: fields}}, nil
	}

	ch := token[0]
	switch {
	case ch == '{', ch == '}', ch == '[', ch == ']', ch == ':', ch == ',', ch == '=':
		return nil, nil
	}
	return nil, fmt.Errorf("unknown type: %s", token)
}

func (p *Parser) parse() (*Type, error) {
	return p.parseType()
}
