package datasheet

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
	"strconv"
	"strings"
)

const (
	dataSheetTypePrefab   = "prefab"   // 预制件数据表
	dataSheetTypeStandard = "standard" // 标准数据表
	dataSheetTypeIndex    = "index"    // 索引数据表
)

type datasheet interface {
	ToType() Type
}

type prefabDatasheet struct {
	Description string             // 数据表描述
	Prefabs     []*datasheetPrefab // 预制件
}

func (p *prefabDatasheet) ToType() Type {
	return nil
}

type standardDatasheet struct {
	Name        string            // 数据表名称
	Description string            // 数据表描述
	Fields      []*datasheetField // 数据表字段
}

func (s *standardDatasheet) ToType() Type {
	value := &Struct{
		StructName: s.Name,
		Anonymity:  true,
	}

	for _, field := range s.Fields {
		value.Fields = append(value.Fields, &Field{
			Owner:     value,
			FieldName: field.Name,
			FieldType: field.Type,
		})
	}

	return value
}

type indexDatasheet struct {
	Name        string            // 数据表名称
	Description string            // 数据表描述
	Fields      []*datasheetField // 数据表字段
}

func (i *indexDatasheet) ToType() Type {
	value := &Struct{
		StructName: i.Name,
		Anonymity:  true,
	}

	for _, field := range i.Fields {
		value.Fields = append(value.Fields, &Field{
			Owner:     value,
			FieldName: field.Name,
			FieldType: field.Type,
		})
	}

	return value
}

type datasheetPrefab struct {
	Name        string  // 预制件名称
	Description string  // 预制件描述
	Prefabs     *Struct // 预制件
}

type datasheetField struct {
	Name        string   // 字段名称
	Description string   // 字段描述
	Groups      []string // 字段分组
	Index       int      // 字段索引顺序
	Type        Type     // 字段类型
}

// loadDatasheets 加载数据表
func loadDatasheets(filePaths ...string) ([]datasheet, error) {
	var prefabs = make(map[string]*Struct)
	var datasheets []datasheet

	type LoadInfo struct {
		path      string
		file      *excelize.File
		sheetName string
		sheetType string
	}

	var loadSorted []*LoadInfo

	for _, path := range filePaths {
		file, err := excelize.OpenFile(path)
		if err != nil {
			return nil, fmt.Errorf("open file %s failed: %w", path, err)
		}

		for _, sheetName := range file.GetSheetList() {
			datasheetType, err := file.GetCellValue(sheetName, "B1")
			if err != nil {
				return nil, fmt.Errorf("get ToType %s[%s] type failed: %w", path, sheetName, err)
			}
			loadSorted = append(loadSorted, &LoadInfo{
				file:      file,
				path:      path,
				sheetName: sheetName,
				sheetType: datasheetType,
			})
		}
	}

	// 预制表最前
	sort.Slice(loadSorted, func(i, j int) bool {
		if loadSorted[i].sheetType == dataSheetTypePrefab && loadSorted[j].sheetType != dataSheetTypePrefab {
			return true
		}
		if loadSorted[i].sheetType != dataSheetTypePrefab && loadSorted[j].sheetType == dataSheetTypePrefab {
			return false
		}
		return loadSorted[i].sheetName < loadSorted[j].sheetName
	})

	type Dep struct {
		value *Struct
		info  *LoadInfo
	}

	prefabDeps := make(map[string][]*Dep)
	parser := NewParser()
	for _, info := range loadSorted {
		tempPrefabDeps := make(map[string][]*Struct) // 预制体名称 => 依赖它的结构体
		switch info.sheetType {
		case dataSheetTypePrefab:
			prefab, err := loadPrefabDatasheet(info.file, info.sheetName, prefabs, tempPrefabDeps, parser)
			if err != nil {
				return nil, err
			}
			datasheets = append(datasheets, prefab)
		case dataSheetTypeStandard:

			standard, err := loadStandardDatasheet(info.file, info.sheetName, prefabs, tempPrefabDeps, parser)
			if err != nil {
				return nil, err
			}
			datasheets = append(datasheets, standard)
		case dataSheetTypeIndex:
			index, err := loadIndexDatasheet(info.file, info.sheetName, prefabs, tempPrefabDeps, parser)
			if err != nil {
				return nil, err
			}
			datasheets = append(datasheets, index)
		default:
			return nil, fmt.Errorf("unsupported ToType %s[%s] type %s", info.path, info.sheetName, info.sheetType)
		}

		for k, v := range tempPrefabDeps {
			for _, s := range v {
				prefabDeps[k] = append(prefabDeps[k], &Dep{
					value: s,
					info:  info,
				})
			}
		}
	}

	// 预制件整理
	for prefabName, deps := range prefabDeps {
		for _, dep := range deps {
			prefab, exist := prefabs[prefabName]
			if !exist {
				return nil, fmt.Errorf("ToType %s[%s] prefab %s not found", dep.info.path, dep.info.sheetName, prefabName)
			}

			dep.value.Fields = prefab.Fields
			dep.value.StructName = prefab.StructName
			dep.value.Anonymity = prefab.Anonymity
		}
	}

	return datasheets, nil
}

func loadPrefabDatasheet(file *excelize.File, sheetName string, prefabs map[string]*Struct, deps map[string][]*Struct, parser *Parser) (*prefabDatasheet, error) {
	rows, err := file.Rows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("get ToType %s[%s] rows failed: %w", file.Path, sheetName, err)
	}

	datasheet := &prefabDatasheet{
		Description: sheetName,
	}

	var line int
	for rows.Next() {
		line++
		if line <= 3 {
			continue
		}
		row, err := rows.Columns()
		if err != nil {
			return nil, fmt.Errorf("get ToType %s[%s] row line %d columns failed: %w", file.Path, sheetName, line, err)
		}

		if len(row) < 3 {
			return nil, fmt.Errorf("ToType %s[%s] row line %d has not enough columns", file.Path, sheetName, line)
		}

		name := row[0]
		desc := row[1]
		define := row[2]

		if name == "" {
			return nil, fmt.Errorf("ToType %s[%s] row line %d prefab name is empty", file.Path, sheetName, line)
		}
		if prefabs[name] != nil {
			return nil, fmt.Errorf("ToType %s[%s] row line %d prefab name %s already exists", file.Path, sheetName, line, name)
		}

		prefabType, err := parser.ParseStruct(define)
		if err != nil {
			return nil, fmt.Errorf("ToType %s[%s] row line %d prefab define %s parse failed: %w", file.Path, sheetName, line, define, err)
		}
		for k, v := range parser.GetPrefabDeps() {
			deps[k] = append(deps[k], v...)
		}
		prefab := prefabType.(*Struct)
		prefab.StructName = name
		prefab.Anonymity = false
		datasheet.Prefabs = append(datasheet.Prefabs, &datasheetPrefab{
			Name:        name,
			Description: desc,
			Prefabs:     prefab,
		})
		if _, exist := prefabs[name]; exist {
			return nil, fmt.Errorf("ToType %s[%s] row line %d prefab name %s already exists", file.Path, sheetName, line, name)
		}
		prefabs[name] = prefab
	}
	return datasheet, nil
}

func loadIndexDatasheet(file *excelize.File, sheetName string, prefabs map[string]*Struct, deps map[string][]*Struct, parser *Parser) (*indexDatasheet, error) {
	datasheetName, err := file.GetCellValue(sheetName, "B2")
	if err != nil {
		return nil, fmt.Errorf("get ToType %s[%s] name failed: %w", file.Path, sheetName, err)
	}

	datasheet := &indexDatasheet{
		Name:        datasheetName,
		Description: sheetName,
		Fields:      nil,
	}

	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("get ToType %s[%s] rows failed: %w", file.Path, sheetName, err)
	}

	if len(rows) < 8 {
		return nil, fmt.Errorf("ToType %s[%s] has not enough rows", file.Path, sheetName)
	}

	// 读取范围内数据
	rows = rows[3:8]
	column := 1
	for {
		if column >= len(rows[0]) {
			break
		}
		desc := rows[0][column]
		fieldName := rows[1][column]
		fieldType := strings.TrimSpace(rows[2][column])
		index := rows[3][column]
		groups := rows[4][column]
		column++

		var indexInt int
		if indexInt, err = strconv.Atoi(index); err != nil {
			return nil, fmt.Errorf("ToType %s[%s] row line %d index %s parse failed: %w", file.Path, sheetName, column, index, err)
		}

		var parsed Type
		switch {
		case strings.HasPrefix(fieldType, "map["):
			parsed, err = parser.ParseMap(fieldType)
		case strings.HasPrefix(fieldType, "["):
			parsed, err = parser.ParseSliceOrArray(fieldType)
		case strings.HasPrefix(fieldType, "{"):
			parsed, err = parser.ParseStruct(fieldType)
		default:
			parsed, err = parser.ParseBasicType(fieldType)
		}
		if err != nil {
			return nil, fmt.Errorf("ToType %s[%s] row line %d field type %s parse failed: %w", file.Path, sheetName, column, fieldType, err)
		}
		for k, v := range parser.GetPrefabDeps() {
			deps[k] = append(deps[k], v...)
		}

		field := &datasheetField{
			Name:        fieldName,
			Description: desc,
			Type:        parsed,
			Index:       indexInt,
			Groups:      strings.Split(groups, ","),
		}

		datasheet.Fields = append(datasheet.Fields, field)
	}
	return datasheet, nil
}

func loadStandardDatasheet(file *excelize.File, sheetName string, prefabs map[string]*Struct, deps map[string][]*Struct, parser *Parser) (*standardDatasheet, error) {
	datasheetName, err := file.GetCellValue(sheetName, "B2")
	if err != nil {
		return nil, fmt.Errorf("get ToType %s[%s] name failed: %w", file.Path, sheetName, err)
	}

	datasheet := &standardDatasheet{
		Name:        datasheetName,
		Description: sheetName,
		Fields:      nil,
	}

	rows, err := file.Rows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("get ToType %s[%s] rows failed: %w", file.Path, sheetName, err)
	}

	var line int
	for rows.Next() {
		line++
		if line <= 4 {
			continue
		}
		row, err := rows.Columns()
		if err != nil {
			return nil, fmt.Errorf("get ToType %s[%s] row line %d columns failed: %w", file.Path, sheetName, line, err)
		}

		if len(row) < 4 {
			return nil, fmt.Errorf("ToType %s[%s] row line %d has not enough columns", file.Path, sheetName, line)
		}

		desc := row[0]
		fieldName := row[1]
		fieldType := strings.TrimSpace(row[2])
		group := row[3]

		var parsed Type
		switch {
		case strings.HasPrefix(fieldType, "map["):
			parsed, err = parser.ParseMap(fieldType)
		case strings.HasPrefix(fieldType, "["):
			parsed, err = parser.ParseSliceOrArray(fieldType)
		case strings.HasPrefix(fieldType, "{"):
			parsed, err = parser.ParseStruct(fieldType)
		default:
			parsed, err = parser.ParseBasicType(fieldType)
		}
		if err != nil {
			return nil, fmt.Errorf("ToType %s[%s] row line %d field type %s parse failed: %w", file.Path, sheetName, line, fieldType, err)
		}

		for k, v := range parser.GetPrefabDeps() {
			deps[k] = append(deps[k], v...)
		}

		datasheet.Fields = append(datasheet.Fields, &datasheetField{
			Name:        fieldName,
			Description: desc,
			Groups:      strings.Split(group, ","),
			Index:       0,
			Type:        parsed,
		})
	}

	return datasheet, nil
}
