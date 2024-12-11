package datasheet

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/parser"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// LoadDatasheets 从文件路径加载数据表，这些数据表将被聚集为一个数据表集合
func LoadDatasheets(filePaths ...string) (*Set, error) {
	set := newSet()
	files := make([]*excelize.File, 0, len(filePaths))
	prefabSheets := make(map[string][]*excelize.File)
	datasheets := make(map[string][]*excelize.File)

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
			id := fmt.Sprintf("%s:%s:%s", path, sheetName, dType)
			if err != nil {
				return nil, fmt.Errorf("get datasheet [%s] %s type failed: %w", sheetName, path, err)
			}
			switch dType {
			case DTypePrefab:
				prefabSheets[id] = append(prefabSheets[id], file)
			case DTypeIndex, DTypeStandard:
				datasheets[id] = append(datasheets[id], file)
			default:
				return nil, fmt.Errorf("unsupport datasheet type %s, [%s] %s", dType, sheetName, path)
			}

		}
	}

	// 解析预制体
	if err := parsePrefab(set, prefabSheets); err != nil {
		return nil, err
	}

	// 解析数据表
	if err := parserDatasheets(set, datasheets); err != nil {
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

func parserDatasheets(set *Set, datasheets map[string][]*excelize.File) error {
	for id, files := range datasheets {
		parts := strings.SplitN(id, ":", 3)
		sheetName := parts[1]
		datasheetType := parts[2]
		for _, file := range files {
			if err := parseDatasheet(set, sheetName, datasheetType, file); err != nil {
				return err
			}
		}
	}

	return nil
}

func parseDatasheet(set *Set, sheetName string, datasheetType DType, file *excelize.File) error {
	absFilepath, err := filepath.Abs(file.Path)
	if err != nil {
		return err
	}
	datasheet := &Datasheet{
		Description: sheetName,
		Filepath:    absFilepath,
	}

	switch datasheetType {
	case DTypeIndex:
		if err := parseIndexDatasheet(set, datasheet, file); err != nil {
			return err
		}
	case DTypeStandard:
		if err := parseStandardDatasheet(set, datasheet, file); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupport datasheet type %s, [%s] %s", datasheetType, sheetName, file.Path)
	}

	set.Datasheets[datasheetType] = append(set.Datasheets[datasheetType], datasheet)
	return nil
}

func parseStandardDatasheet(set *Set, datasheet *Datasheet, file *excelize.File) error {
	sheetName := datasheet.Description
	datasheetName, err := file.GetCellValue(sheetName, "B2")
	if err != nil {
		return fmt.Errorf("get index datasheet %s[%s] name failed: %w", file.Path, sheetName, err)
	}

	datasheet.Name = datasheetName
	rows, err := file.Rows(sheetName)
	if err != nil {
		return fmt.Errorf("get %s[%s] rows failed: %w", file.Path, sheetName, err)
	}

	var line int
	for rows.Next() {
		line++
		if line <= 4 {
			continue
		}
		row, err := rows.Columns()
		if err != nil {
			return fmt.Errorf("get  %s[%s] row line %d columns failed: %w", file.Path, sheetName, line, err)
		}

		if len(row) < 4 {
			return fmt.Errorf(" %s[%s] row line %d has not enough columns", file.Path, sheetName, line)
		}

		desc := row[0]
		fieldName := row[1]
		fieldType := strings.TrimSpace(row[2])
		groups := row[3]
		fieldValue := row[4]

		var optional = strings.HasPrefix(fieldType, "*")
		if optional {
			fieldType = fieldType[1:]
		}

		var parsed Type
		if parsed, err = parserStruct(datasheet.Name, fieldType, set.Prefabs); err != nil {
			return fmt.Errorf("parse %s[%s] row line %d field %s failed: %w", file.Path, sheetName, line, fieldName, err)
		}

		datasheet.Fields = append(datasheet.Fields, &Field{
			Owner:       datasheet,
			Name:        fieldName,
			Optional:    optional,
			Type:        parsed,
			Description: desc,
			Index:       0,
			Groups:      strings.Split(groups, ","),
			Values:      []string{fieldValue},
		})
	}

	return nil
}

func parseIndexDatasheet(set *Set, datasheet *Datasheet, file *excelize.File) error {
	sheetName := datasheet.Description
	datasheetName, err := file.GetCellValue(sheetName, "B2")
	if err != nil {
		return fmt.Errorf("get index datasheet %s[%s] name failed: %w", file.Path, sheetName, err)
	}

	datasheet.Name = datasheetName

	rows, err := file.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("get index datasheet %s[%s] rows failed: %w", file.Path, sheetName, err)
	}

	if len(rows) < 8 {
		return fmt.Errorf("ToType %s[%s] has not enough rows", file.Path, sheetName)
	}

	// 读取范围内数据
	column := 1
	for {
		if column >= len(rows[3]) {
			break
		}
		desc := rows[3][column]
		fieldName := rows[4][column]
		fieldType := strings.TrimSpace(rows[5][column])
		index := rows[6][column]
		groups := rows[7][column]
		var values []string
		var valueRow = 8
		for valueRow < len(rows) && column < len(rows[valueRow]) {
			values = append(values, rows[valueRow][column])
			valueRow++
		}

		column++

		var indexInt int
		if indexInt, err = strconv.Atoi(index); err != nil {
			return fmt.Errorf("parse index datasheet %s[%s] row line %d field %s index %s failed: %w", file.Path, sheetName, column, fieldName, index, err)
		}

		var optional = strings.HasPrefix(fieldType, "*")
		if optional {
			fieldType = fieldType[1:]
		}

		var parsed Type
		if parsed, err = parserStruct(datasheet.Name, fieldType, set.Prefabs); err != nil {
			return fmt.Errorf("parse index datasheet %s[%s] row line %d field %s failed: %w", file.Path, sheetName, column, fieldName, err)
		}

		datasheet.Fields = append(datasheet.Fields, &Field{
			Owner:       datasheet,
			Name:        fieldName,
			Optional:    optional,
			Type:        parsed,
			Description: desc,
			Index:       indexInt,
			Groups:      strings.Split(groups, ","),
			Values:      values,
		})
	}

	sort.Slice(datasheet.Fields, func(i, j int) bool {
		// 0 最后
		if datasheet.Fields[i].Index == 0 {
			return false
		}
		if datasheet.Fields[j].Index == 0 {
			return true
		}
		return datasheet.Fields[i].Index < datasheet.Fields[j].Index
	})

	// 检查索引是否由 1 开始
	if startField := datasheet.Fields[0]; startField.Index != 1 {
		return fmt.Errorf("parse index datasheet(%s:%s) failed, %w, got: %d, from: %s", datasheet.Filepath, datasheet.Name, ErrorIndexStart, startField.Index, startField.Name)
	}

	// 检查索引是否连续
	for i := 0; i < len(datasheet.Fields)-1; i++ {
		curr := datasheet.Fields[i]
		next := datasheet.Fields[i+1]
		if curr.Index+1 != next.Index && next.Index != 0 {
			return fmt.Errorf("parse index datasheet(%s:%s) failed, %w, curr:%s(%d), next: %s(%d)", datasheet.Filepath, datasheet.Name, ErrorIndexContinuous, curr.Name, curr.Index, next.Name, next.Index)
		}
	}

	return nil
}

func parsePrefab(set *Set, sheets map[string][]*excelize.File) error {
	prefabs := make(map[string]*Prefab)
	for id, files := range sheets {
		sheetName := strings.SplitN(id, ":", 3)[1]
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

		prefabType, err := parserStruct(name, define, nil)
		if err != nil {
			return err
		}

		if _, exist := prefabs[name]; exist {
			return fmt.Errorf("prefab %s existed", name)
		}
		prefabs[name] = &Prefab{
			Type:        prefabType,
			Name:        name,
			Description: desc,
		}
	}
	return nil
}

func parserStruct(name string, define string, prefabs map[string]*Prefab) (Type, error) {
	structParser := parser.New[Type](
		parser.SymbolLeftBrace, parser.SymbolRightBrace, parser.SymbolDot, parser.SymbolComma, parser.SymbolPipe,
		parser.SymbolColon, parser.SymbolLeftBracket, parser.SymbolRightBracket, parser.SymbolAsterisk,
	)
	tokenized := structParser.Tokenize(define)

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
			if prefabs != nil {
				prefab, exist := prefabs[token.String()]
				if !exist {
					return nil, fmt.Errorf("prefab %s not found", token)
				}
				return prefab, nil
			}
			return &TodoPrefab{
				Optional: optional,
				Name:     token.String(),
			}, nil
		}
	}

	parsedType, err := tokenized.Parse(parserHandler)
	if err != nil {
		return nil, err
	}
	if structType, ok := parsedType.(*Struct); ok {
		structType.Name = name
	}
	return parsedType, nil
}
