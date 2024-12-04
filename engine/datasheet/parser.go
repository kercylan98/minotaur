package datasheet

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/parser"
	"regexp"
	"strings"
)

var (
	ErrInvalidStruct      = errors.New("invalid struct")
	ErrInvalidStructField = errors.New("invalid struct field")
	ErrInvalidArray       = errors.New("invalid array")
	ErrInvalidBasicType   = errors.New("invalid basic type")
	ErrInvalidMap         = errors.New("invalid map")
)

var (
	// nameRegexp 匹配名称的正则表达式
	//  - [a-zA-Z] 首字母必须为字母
	//  - [a-zA-Z0-9_]* 后续字符必须为字母、数字或下划线
	nameRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
)

func NewParser() *Parser {
	p := &Parser{
		parser: parser.New[Type](
			parser.SymbolLeftBrace, parser.SymbolRightBrace, parser.SymbolDot, parser.SymbolComma, parser.SymbolPipe,
			parser.SymbolColon, parser.SymbolLeftBracket, parser.SymbolRightBracket,
		),
	}

	return p
}

type Parser struct {
	parser  *parser.Parser[Type]
	prefabs map[string][]*Struct // 预制件依赖
}

func (s *Parser) GetPrefabDeps() map[string][]*Struct {
	return s.prefabs
}

func (s *Parser) parse(p parser.FunctionalHandler[Type]) parser.Handler[Type] {
	return parser.FunctionalHandler[Type](func(tokens *parser.Tokens[Type]) (Type, error) {
		s.prefabs = nil

		t, err := p.Handle(tokens)
		if err != nil {
			return nil, err
		}

		return t, nil
	})
}

func (s *Parser) ParseStruct(input string) (Type, error) {
	tokenized := s.parser.Tokenize(input)
	return tokenized.Parse(s.parse(s.parseStruct))
}

func (s *Parser) ParseSliceOrArray(input string) (Type, error) {
	tokenized := s.parser.Tokenize(input)
	return tokenized.Parse(s.parse(s.parseSliceOrArray))
}

func (s *Parser) ParseBasicType(input string) (Type, error) {
	tokenized := s.parser.Tokenize(input)
	return tokenized.Parse(s.parse(s.parseBasicType))
}

func (s *Parser) ParseMap(input string) (Type, error) {
	tokenized := s.parser.Tokenize(input)
	return tokenized.Parse(s.parse(s.parseMap))
}

func (s *Parser) parseStruct(tokens *parser.Tokens[Type]) (Type, error) {
	token := tokens.Consume()
	value := &Struct{}

	// 结构体头解析
	if !token.EqualSymbol(parser.SymbolLeftBrace) {
		return nil, fmt.Errorf("%w, expect '{', but got %s", ErrInvalidStruct, token)
	}

	// 名称解析
	name := tokens.Consume()
	if !nameRegexp.MatchString(name.String()) {
		return nil, fmt.Errorf("%w, got: %s", ErrInvalidStruct, name)
	}

	// 解析是结构体名称还是字段
	token = tokens.Consume()
	if token.EqualSymbol(parser.SymbolPipe) {
		// 结构体名称
	} else if token.EqualSymbol(parser.SymbolColon) {
		// 匿名结构体读取到字段名称，回退消费，后续开始字段读取
		value.Anonymity = true
		name = ""
		tokens.Undo() // 回退 ':'
		tokens.Undo() // 回退名称
	} else {
		return nil, fmt.Errorf("%w, expect '|' or '.', but got %s", ErrInvalidStruct, token)
	}

	// 字段解析
	for {
		fieldName := tokens.Consume()
		if !nameRegexp.MatchString(fieldName.String()) {
			return nil, fmt.Errorf("%w, got: %s", ErrInvalidStructField, fieldName)
		}

		// 消费 ':'
		if !tokens.Consume().EqualSymbol(parser.SymbolColon) {
			return nil, fmt.Errorf("%w, expect ':', but got %s", ErrInvalidStructField, fieldName)
		}

		fieldTypeToken := tokens.Consume()
		if fieldTypeToken.Empty() {
			return nil, fmt.Errorf("%w, expect %s field type, but got empty", ErrInvalidStructField, fieldTypeToken)
		}

		// 条件解析，消费第一个类型符号
		var fieldType any
		var err error
		switch {
		case fieldTypeToken == "map":
			tokens.Undo()
			fieldType, err = s.parseMap(tokens)
			if err != nil {
				return nil, err
			}
			value.Fields = append(value.Fields, &Field{
				Owner:     value,
				FieldName: fieldName.String(),
				FieldType: fieldType.(*Map),
			})
		case fieldTypeToken.EqualSymbol(parser.SymbolLeftBracket):
			tokens.Undo()
			fieldType, err = s.parseSliceOrArray(tokens)
			if err != nil {
				return nil, err
			}
			var t Type
			switch v := fieldType.(type) {
			case *Array:
				t = v
			case *Slice:
				t = v
			}
			value.Fields = append(value.Fields, &Field{
				Owner:     value,
				FieldName: fieldName.String(),
				FieldType: t,
			})
		case fieldTypeToken.EqualSymbol(parser.SymbolLeftBrace):
			// 回退类型消费
			tokens.Undo()
			fieldType, err = s.parseStruct(tokens)
			if err != nil {
				return nil, err
			}
			value.Fields = append(value.Fields, &Field{
				Owner:     value,
				FieldName: fieldName.String(),
				FieldType: fieldType.(*Struct),
			})
		default:
			// 回退类型消费
			tokens.Undo()
			// 尝试解析基础类型
			fieldType, err = s.parseBasicType(tokens)
			if err != nil {
				return nil, err
			}
			var t Type
			switch v := fieldType.(type) {
			case *BasicType:
				t = v
			case *Struct: // 预制体
				t = v
			}
			value.Fields = append(value.Fields, &Field{
				Owner:     value,
				FieldName: fieldName.String(),
				FieldType: t,
			})
		}

		// 检查是否还有 ',', 如果有则继续解析字段
		token = tokens.Consume()
		if token.EqualSymbol(parser.SymbolRightBrace) {
			break
		} else if !token.EqualSymbol(parser.SymbolComma) {
			return nil, fmt.Errorf("%w, expect ',' or '}', but got %s", ErrInvalidStructField, token.String())
		}
	}

	value.StructName = name.String()
	return value, nil
}

func (s *Parser) parseSliceOrArray(tokens *parser.Tokens[Type]) (Type, error) {
	token := tokens.Consume()
	if !token.EqualSymbol(parser.SymbolLeftBracket) {
		return nil, fmt.Errorf("%w, expect '[', but got %s", ErrInvalidArray, token)
	}

	next := tokens.Consume()
	if next.EqualSymbol(parser.SymbolRightBracket) {
		// 条件解析
		token := tokens.Consume()
		switch {
		case token == "map":
			tokens.Undo()
			value, err := s.parseMap(tokens)
			if err != nil {
				return nil, err
			}
			return &Slice{
				ValueType: value.(*Map),
			}, nil
		case token.EqualSymbol(parser.SymbolLeftBrace):
			tokens.Undo()
			// 匿名结构体
			value, err := s.parseStruct(tokens)
			if err != nil {
				return nil, err
			}
			return &Slice{
				ValueType: value.(*Struct),
			}, nil
		case token.EqualSymbol(parser.SymbolLeftBracket):
			// 数组或切片
			tokens.Undo()
			value, err := s.parseSliceOrArray(tokens)
			if err != nil {
				return nil, err
			}
			switch v := value.(type) {
			case *Array:
				return &Slice{ValueType: v}, nil
			case *Slice:
				return &Slice{ValueType: v.ValueType}, nil
			default:
				return nil, fmt.Errorf("%w, expect ']', but got %s", ErrInvalidArray, tokens.Peek().String())
			}
		default:
			tokens.Undo()
			value, err := s.parseBasicType(tokens)
			if err != nil {
				return nil, err
			}
			var t Type
			switch v := value.(type) {
			case *BasicType:
				t = v
			case *Struct: // 预制体
				t = v
			}
			return &Slice{
				ValueType: t,
			}, nil
		}

	} else {
		// 数组
		var size = -1
		if _, err := fmt.Sscanf(next.String(), "%d", &size); err != nil {
			return nil, fmt.Errorf("%w, expect array size, but got %s", ErrInvalidStructField, next.String())
		}
		if token := tokens.Consume(); !token.EqualSymbol(parser.SymbolRightBracket) {
			return nil, fmt.Errorf("%w, expect ']', but got %s", ErrInvalidArray, token.String())
		}

		// 条件解析
		token := tokens.Consume()
		switch {
		case token == "map":
			tokens.Undo()
			value, err := s.parseMap(tokens)
			if err != nil {
				return nil, err
			}
			return &Array{
				Length:    size,
				ValueType: value.(*Map),
			}, nil
		case token.EqualSymbol(parser.SymbolLeftBrace):
			tokens.Undo()
			// 匿名结构体
			value, err := s.parseStruct(tokens)
			if err != nil {
				return nil, err
			}
			return &Array{
				Length:    size,
				ValueType: value.(*Struct),
			}, nil
		case token.EqualSymbol(parser.SymbolLeftBracket):
			// 数组或切片
			tokens.Undo()
			value, err := s.parseSliceOrArray(tokens)
			if err != nil {
				return nil, err
			}
			switch v := value.(type) {
			case *Array:
				return &Array{
					Length:    size,
					ValueType: v,
				}, nil
			case *Slice:
				return &Array{
					Length:    size,
					ValueType: v.ValueType,
				}, nil
			default:
				return nil, fmt.Errorf("%w, expect ']', but got %s", ErrInvalidArray, tokens.Peek().String())
			}
		default:
			tokens.Undo()
			value, err := s.parseBasicType(tokens)
			if err != nil {
				return nil, err
			}
			var t Type
			switch v := value.(type) {
			case *BasicType:
				t = v
			case *Struct: // 预制体
				t = v
			}
			return &Array{
				Length:    size,
				ValueType: t,
			}, nil
		}
	}
}

func (s *Parser) parseBasicType(tokens *parser.Tokens[Type]) (Type, error) {
	token := tokens.Consume()
	tokenLower := strings.ToLower(token.String())
	switch tokenLower {
	case "boolean", "bool", "int", "int8", "int16", "int32", "int64", "uint",
		"uint8", "uint16", "uint32", "uint64", "float", "double", "byte", "time",
		"float32", "float64", "short", "long", "string", "bigint", "bigfloat":
		return &BasicType{TypeName: tokenLower}, nil
	default:
		// 预制件依赖
		if s.prefabs == nil {
			s.prefabs = make(map[string][]*Struct)
		}
		v := &Struct{
			StructName: token.String(),
		}
		s.prefabs[v.StructName] = append(s.prefabs[v.StructName], v)
		return v, nil
	}
}

func (s *Parser) parseMap(tokens *parser.Tokens[Type]) (Type, error) {
	token := tokens.Consume()
	if token != "map" {
		return nil, fmt.Errorf("%w, expect 'map', but got %s", ErrInvalidMap, token)
	}

	token = tokens.Consume()
	if !token.EqualSymbol(parser.SymbolLeftBracket) {
		return nil, fmt.Errorf("%w, expect '[', but got %s", ErrInvalidMap, token.String())
	}

	keyType, err := s.parseBasicType(tokens)
	switch v := keyType.(type) {
	case *Struct: // 预制体
		return nil, fmt.Errorf("%w, key expect basic type, got %s", ErrInvalidMap, v.StructName)
	}
	if err != nil {
		return nil, fmt.Errorf("%w, key expect basic type, %w", ErrInvalidMap, err)
	}
	if token := tokens.Consume(); !token.EqualSymbol(parser.SymbolRightBracket) {
		return nil, fmt.Errorf("%w, expect ']', but got %s", ErrInvalidMap, token.String())
	}

	// 条件解析
	token = tokens.Consume()
	switch {
	case token == "map":
		tokens.Undo()
		value, err := s.parseMap(tokens)
		if err != nil {
			return nil, err
		}
		return &Map{
			KeyType:   keyType.(*BasicType),
			ValueType: value.(*Map),
		}, nil
	case token.EqualSymbol(parser.SymbolLeftBrace):
		tokens.Undo()
		// 匿名结构体
		value, err := s.parseStruct(tokens)
		if err != nil {
			return nil, err
		}
		return &Map{
			KeyType:   keyType.(*BasicType),
			ValueType: value.(*Struct),
		}, nil
	case token.EqualSymbol(parser.SymbolLeftBracket):
		// 数组或切片
		tokens.Undo()
		value, err := s.parseSliceOrArray(tokens)
		if err != nil {
			return nil, err
		}
		switch v := value.(type) {
		case *Array:
			return &Map{
				KeyType:   keyType.(*BasicType),
				ValueType: v,
			}, nil
		case *Slice:
			return &Map{
				KeyType:   keyType.(*BasicType),
				ValueType: v.ValueType,
			}, nil
		default:
			return nil, fmt.Errorf("%w, expect array or slice, but got %s", ErrInvalidMap, token.String())
		}
	default:
		tokens.Undo()
		value, err := s.parseBasicType(tokens)
		if err != nil {
			return nil, fmt.Errorf("%w, value expect basic type, %w", ErrInvalidMap, err)
		}
		var t Type
		switch v := value.(type) {
		case *BasicType:
			t = v
		case *Struct: // 预制体
			t = v
		}
		return &Map{
			KeyType:   keyType.(*BasicType),
			ValueType: t,
		}, nil
	}
}
