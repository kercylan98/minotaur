package fieldparser

import "github.com/kercylan98/minotaur/ax/cmd/table"

func New() table.TypeParser {
	return &parser{}
}

type parser struct {
	lexer *lexer
	token token

	types []table.Type // 索引越靠后越外层
}

func (*parser) Parse(input string) []table.Type {
	p := &parser{lexer: newLexer(input)}
	p.nextToken()
	switch t := p.parseType().(type) {
	case *table.StructType:
	default:
		p.types = append(p.types, t)
	}
	return p.types
}

func (p *parser) nextToken() {
	p.token = p.lexer.lex()
}

func (p *parser) parseType() table.Type {
	switch p.token.Type {
	case TokenIdent:
		ident := p.token.Value
		p.nextToken()
		return &table.BasicType{
			Type: ident,
		}
	case TokenLbracket:
		p.nextToken()
		if p.token.Type != TokenRbracket {
			panic("expected ']'")
		}
		p.nextToken()
		elemType := p.parseType()
		return &table.ArrayType{
			ElementType: elemType,
		}
	case TokenLbrace:
		return p.parseStruct()
	default:
		panic("unexpected token")
	}
}

func (p *parser) parseStruct() *table.StructType {
	p.nextToken() // consume '{'
	var fields []*table.StructField
	for p.token.Type != TokenRbrace {
		if p.token.Type != TokenIdent {
			panic("expected field Type")
		}
		fieldName := p.token.Value
		p.nextToken()
		if p.token.Type != TokenColon {
			panic("expected ':'")
		}
		p.nextToken() // consume ':'
		fieldType := p.parseType()
		if structType, ok := fieldType.(*table.StructType); ok {
			structType.Name = fieldName
			fieldType = structType
		}
		fields = append(fields, &table.StructField{
			Name: fieldName,
			Type: fieldType,
		})
		if p.token.Type == TokenComma {
			p.nextToken() // consume ','
		}
	}
	p.nextToken() // consume '}'

	structType := &table.StructType{
		Fields: fields,
	}
	p.types = append(p.types, structType)
	return structType
}
