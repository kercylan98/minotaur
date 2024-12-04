package datasheet

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/parser"
	"github.com/xuri/excelize/v2"
	"strings"
)

// LoadDatasheets 从文件路径加载数据表，这些数据表将被聚集为一个数据表集合
func LoadDatasheets(filePaths ...string) (*Set, error) {
	set := newSet()
	files := make([]*excelize.File, 0, len(filePaths))
	prefabSheets := make(map[string][]*excelize.File)

	// 打开文件
	for _, path := range filePaths {
		file, err := excelize.OpenFile(path)
		if err != nil {
			return nil, fmt.Errorf("open file failed: %w", err)
		}
		files = append(files, file)

		var dType DType
		for _, sheetName := range file.GetSheetList() {
			dType, err = file.GetCellValue(sheetName, "B1")
			if err != nil {
				return nil, fmt.Errorf("get datasheet [%s] %s type failed: %w", sheetName, path, err)
			}
			switch dType {
			case DTypePrefab:
				prefabSheets[sheetName] = append(prefabSheets[sheetName], file)
			case DTypeIndex, DTypeStandard:
				set.Datasheets[dType] = append(set.Datasheets[dType], &Datasheet{
					Description: sheetName,
				})
			default:
				return nil, fmt.Errorf("unsupport datasheet type %s, [%s] %s", dType, sheetName, path)
			}

		}
	}

	// 解析预制体
	if err := parsePrefab(set, prefabSheets); err != nil {
		return nil, err
	}

	// 清理文件
	for _, file := range files {
		if err := file.Close(); err != nil {
			return nil, fmt.Errorf("close datasheet file %s failed, err: %w", file.Path, err)
		}
	}

	return set, nil
}

func parsePrefab(set *Set, sheets map[string][]*excelize.File) error {
	prefabs := make(map[string]*Prefab)
	for sheetName, files := range sheets {
		for _, file := range files {
			if err := parsePrefabSheet(prefabs, sheetName, file); err != nil {
				return err
			}
		}
	}

	// 预制体依赖处理
	var prefabDepHandler func(t Type) error
	prefabDepHandler = func(t Type) error {
		var exist bool
		switch value := t.(type) {
		case *Struct:
			for _, field := range value.Fields {
				if field.Type.isTodoPrefab() {
					exceptName := field.Type.(*TodoPrefab).Name
					if field.Type, exist = prefabs[exceptName]; !exist {
						return fmt.Errorf("prefab %s not existed", exceptName)
					}
				}
				if err := prefabDepHandler(field.Type); err != nil {
					return err
				}
			}
		case *Array:
			if value.Type.isTodoPrefab() {
				exceptName := value.Type.(*TodoPrefab).Name
				if value.Type, exist = prefabs[exceptName]; !exist {
					return fmt.Errorf("prefab %s not existed", exceptName)
				}
				if err := prefabDepHandler(value.Type); err != nil {
					return err
				}
			}
		case *Slice:
			if value.Type.isTodoPrefab() {
				exceptName := value.Type.(*TodoPrefab).Name
				if value.Type, exist = prefabs[exceptName]; !exist {
					return fmt.Errorf("prefab %s not existed", exceptName)
				}
				if err := prefabDepHandler(value.Type); err != nil {
					return err
				}
			}
		case *Map:
			if value.ValueType.isTodoPrefab() {
				exceptName := value.ValueType.(*TodoPrefab).Name
				if value.ValueType, exist = prefabs[exceptName]; !exist {
					return fmt.Errorf("prefab %s not existed", exceptName)
				}
				if err := prefabDepHandler(value.ValueType); err != nil {
					return err
				}
			}
		}
		return nil
	}

	for _, prefab := range prefabs {
		if err := prefabDepHandler(prefab.Type); err != nil {
			return err
		}
	}

	set.Prefabs = prefabs
	return nil
}

func parsePrefabSheet(prefabs map[string]*Prefab, sheetName string, file *excelize.File) error {
	rows, err := file.Rows(sheetName)
	if err != nil {
		return fmt.Errorf("get ToType [%s] %s rows failed: %w", sheetName, file.Path, err)
	}

	var line int
	for rows.Next() {
		line++
		if line <= 3 {
			continue
		}
		row, err := rows.Columns()
		if err != nil {
			return fmt.Errorf("get prefab sheet columns failed: [%s:%d] %s, %w", sheetName, line, file.Path, err)
		}

		if len(row) < 3 {
			return fmt.Errorf("prefab sheet has not enough columns, [%s:%d] %s", sheetName, line, file.Path)
		}

		name := row[0]
		desc := row[1]
		define := row[2]

		if name == "" {
			return fmt.Errorf("prefab name is empty, [%s:%d] %s", sheetName, line, file.Path)
		}

		prefabType, err := parserPrefab(name, define)
		if err != nil {
			return err
		}

		if _, exist := prefabs[name]; exist {
			return fmt.Errorf("prefab %s existed", name)
		}
		prefabs[name] = &Prefab{
			Type:        prefabType,
			Description: desc,
		}
	}
	return nil
}

func parserPrefab(name string, define string) (Type, error) {
	prefabParser := parser.New[Type](
		parser.SymbolLeftBrace, parser.SymbolRightBrace, parser.SymbolDot, parser.SymbolComma, parser.SymbolPipe,
		parser.SymbolColon, parser.SymbolLeftBracket, parser.SymbolRightBracket, parser.SymbolAsterisk,
	)
	tokenized := prefabParser.Tokenize(define)

	var parserHandler parser.FunctionalHandler[Type]
	parserHandler = func(tokens *parser.Tokens[Type]) (Type, error) {
		token := tokens.Peek()

		switch {
		case token.EqualSymbol(parser.SymbolLeftBrace): // 结构体
			value := &Struct{}
			token = tokens.Consume() // 消费 '{'
			for !tokens.Peek().EqualSymbol(parser.SymbolRightBrace) {
				var fieldName string
				var fieldType Type
				var optional bool
				var err error

				// 字段名
				if token = tokens.Consume(); !nameRegexp.MatchString(token.String()) {
					return nil, fmt.Errorf("struct field name invlide: %s", token)
				}
				fieldName = token.String()

				// 分隔符
				if token = tokens.Consume(); !token.EqualSymbol(parser.SymbolColon) {
					return nil, fmt.Errorf("invlide struct, except: ':' got: %s", token)
				}

				// 是否可选
				if token = tokens.Consume(); token.EqualSymbol(parser.SymbolAsterisk) {
					optional = true
				} else {
					tokens.Undo()
				}

				// 字段类型
				if fieldType, err = parserHandler(tokens); err != nil {
					return nil, err
				}

				value.Fields = append(value.Fields, &StructField{
					Owner:    value,
					Name:     fieldName,
					Optional: optional,
					Type:     fieldType,
				})

				// 消费 ','
				if token = tokens.Consume(); !token.EqualSymbol(parser.SymbolComma) {
					tokens.Undo()
					break
				}
			}
			// 消费 '}'
			if token = tokens.Consume(); !token.EqualSymbol(parser.SymbolRightBrace) {
				return nil, fmt.Errorf("invlide struct, except: '}' got: %s", token)
			}
			return value, nil
		case strings.HasPrefix(token.String(), "map"): // MAP
			// 消费 "map"
			tokens.Consume()

			// 消费 '['
			if token = tokens.Consume(); !token.EqualSymbol(parser.SymbolLeftBracket) {
				return nil, fmt.Errorf("invlide map, except: '[' got: %s", token)
			}

			// Key 类型
			keyType, err := parserHandler(tokens)
			if err != nil {
				return nil, err
			}
			if _, ok := keyType.(*Basic); !ok {
				return nil, fmt.Errorf("invlide map, key except basic type, but got: %T", keyType)
			}

			// 消费 ']'
			if token = tokens.Consume(); !token.EqualSymbol(parser.SymbolRightBracket) {
				return nil, fmt.Errorf("invlide map, except: ']' got: %s", token)
			}

			// 是否可选
			var optional bool
			if token = tokens.Consume(); token.EqualSymbol(parser.SymbolAsterisk) {
				optional = true
			} else {
				tokens.Undo()
			}

			// Val 类型
			valType, err := parserHandler(tokens)
			if err != nil {
				return nil, err
			}

			return &Map{
				KeyType:       keyType,
				ValueType:     valType,
				ValueOptional: optional,
			}, nil
		case token.EqualSymbol(parser.SymbolLeftBracket) && tokens.PeekOffset(1).EqualSymbol(parser.SymbolRightBracket): // 切片
			// 消费 '['、']'
			tokens.Consume()
			tokens.Consume()

			// 是否可选
			var optional bool
			if token = tokens.Consume(); token.EqualSymbol(parser.SymbolAsterisk) {
				optional = true
			} else {
				tokens.Undo()
			}

			valueType, err := parserHandler(tokens)
			if err != nil {
				return nil, err
			}

			return &Slice{
				Optional: optional,
				Type:     valueType,
			}, nil
		case token.EqualSymbol(parser.SymbolLeftBracket) && !tokens.PeekOffset(1).EqualSymbol(parser.SymbolRightBracket): // 数组
			// 消费 '['
			tokens.Consume()

			// 长度
			token = tokens.Consume()
			var size = -1
			if _, err := fmt.Sscanf(token.String(), "%d", &size); err != nil {
				return nil, fmt.Errorf("invalid array, expect array size, but got %s", token)
			}

			// 消费 ']'
			if token = tokens.Consume(); !token.EqualSymbol(parser.SymbolRightBracket) {
				return nil, fmt.Errorf("invalid array, except: ']' got: %s", token)
			}

			// 是否可选
			var optional bool
			if token = tokens.Consume(); token.EqualSymbol(parser.SymbolAsterisk) {
				optional = true
			} else {
				tokens.Undo()
			}

			valueType, err := parserHandler(tokens)
			if err != nil {
				return nil, err
			}

			return &Array{
				Length:   size,
				Optional: optional,
				Type:     valueType,
			}, nil
		case IsBasicType(strings.ToLower(token.String())): // 基本类型
			tokens.Consume()
			return &Basic{Name: token.String()}, nil
		default: // 预制体
			// 是否可选
			var optional bool
			if token = tokens.Consume(); token.EqualSymbol(parser.SymbolAsterisk) {
				optional = true
			} else {
				tokens.Undo()
			}

			tokens.Consume()
			return &TodoPrefab{
				Optional: optional,
				Name:     token.String(),
			}, nil
		}
	}

	prefabType, err := tokenized.Parse(parserHandler)
	if err != nil {
		return nil, err
	}
	if structType, ok := prefabType.(*Struct); ok {
		structType.Name = name
	}
	return prefabType, nil
}
