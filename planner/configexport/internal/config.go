package internal

import (
	"errors"
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tealeg/xlsx"
	"strings"
)

func NewConfig(sheet *xlsx.Sheet) *Config {
	config := &Config{
		Sheet: sheet,
	}

	if len(config.Sheet.Rows) < 2 || len(config.Sheet.Rows[1].Cells) < 2 {
		return nil
	}

	if config.GetIndexCount() > 0 && len(config.Sheet.Rows) < 7 {
		return nil
	}

	if config.GetIndexCount() > 0 {
		var (
			skipField       = make(map[int]bool)
			describeLine    = config.Sheet.Rows[3]
			nameLine        = config.Sheet.Rows[4]
			typeLine        = config.Sheet.Rows[5]
			exportParamLine = config.Sheet.Rows[6]
		)
		// 分析数据
		{
			for i := 1; i < len(describeLine.Cells); i++ {
				describe := strings.TrimSpace(nameLine.Cells[i].String())
				if strings.HasPrefix(describe, "#") {
					skipField[i] = true
					continue
				}
				typ := strings.TrimSpace(typeLine.Cells[i].String())
				if strings.HasPrefix(typ, "#") || len(typ) == 0 {
					skipField[i] = true
					continue
				}
				exportParam := strings.TrimSpace(exportParamLine.Cells[i].String())
				if strings.HasPrefix(exportParam, "#") || len(exportParam) == 0 {
					skipField[i] = true
					continue
				}
			}
			if len(nameLine.Cells)-1-len(skipField) < config.GetIndexCount() {
				panic(errors.New("index count must greater or equal to field count"))
			}
		}

		config.skipField = skipField
		config.describeLine = describeLine
		config.nameLine = nameLine
		config.typeLine = typeLine
		config.exportParamLine = exportParamLine

		// 整理数据
		var (
			dataLine  = make([]map[any]any, len(config.Sheet.Rows))
			dataLineC = make([]map[any]any, len(config.Sheet.Rows))
		)
		for i := 1; i < len(config.describeLine.Cells); i++ {
			if skipField[i] {
				continue
			}

			var (
				//describe    = strings.TrimSpace(describeLine.Cells[i].String())
				name = strings.TrimSpace(nameLine.Cells[i].String())
				//typ         = strings.TrimSpace(typeLine.Cells[i].String())
				exportParam = strings.TrimSpace(exportParamLine.Cells[i].String())
			)

			for row := 7; row < len(config.Sheet.Rows); row++ {
				//value := slf.Sheet.Rows[row].Cells[i].String()
				var value = getValueWithType(typeLine.Cells[i].String(), config.Sheet.Rows[row].Cells[i].String())

				var needC, needS bool
				switch strings.ToLower(exportParam) {
				case "c":
					needC = true
				case "s":
					needS = true
				case "sc", "cs":
					needS = true
					needC = true
				}

				if needC {
					line := dataLineC[row]
					if line == nil {
						line = map[any]any{}
						dataLineC[row] = line
					}
					line[name] = value
				}

				if needS {
					line := dataLine[row]
					if line == nil {
						line = map[any]any{}
						dataLine[row] = line
					}
					line[name] = value
				}

			}
		}

		// 索引
		var dataSource = make(map[any]any)
		var data = dataSource
		var dataSourceC = make(map[any]any)
		var dataC = dataSourceC
		var index = config.GetIndexCount()
		var currentIndex = 0
		for row := 7; row < len(config.Sheet.Rows); row++ {
			for i, cell := range config.Sheet.Rows[row].Cells {
				if i == 0 {
					if strings.HasPrefix(typeLine.Cells[i].String(), "#") {
						break
					}
					continue
				}
				if skipField[i] {
					continue
				}
				var value = getValueWithType(typeLine.Cells[i].String(), cell.String())

				if currentIndex < index {
					currentIndex++
					m, exist := data[value]
					if !exist {
						if currentIndex == index {
							data[value] = dataLine[row]
						} else {
							m = map[any]any{}
							data[value] = m
						}
					}
					if currentIndex < index {
						data = m.(map[any]any)
					}

					m, exist = dataC[value]
					if !exist {
						if currentIndex == index {
							dataC[value] = dataLineC[row]
						} else {
							m = map[any]any{}
							dataC[value] = m
						}
					}
					if currentIndex < index {
						dataC = m.(map[any]any)
					}
				}
			}
			data = dataSource
			dataC = dataSourceC
			currentIndex = 0
		}
		config.data = dataSource
		config.dataC = dataSourceC
	} else {
		config.data = map[any]any{}
		config.dataC = map[any]any{}
		for i := 4; i < len(config.Rows); i++ {
			row := config.Sheet.Rows[i]

			desc := row.Cells[0].String()
			if strings.HasPrefix(desc, "#") {
				continue
			}
			name := row.Cells[1].String()
			typ := row.Cells[2].String()
			exportParam := row.Cells[3].String()
			value := row.Cells[4].String()

			switch strings.ToLower(exportParam) {
			case "c":
				config.dataC[name] = getValueWithType(typ, value)
			case "s":
				config.data[name] = getValueWithType(typ, value)
			case "sc", "cs":
				config.dataC[name] = getValueWithType(typ, value)
				config.data[name] = getValueWithType(typ, value)
			}
		}
	}
	return config
}

type Config struct {
	*xlsx.Sheet
	skipField       map[int]bool
	describeLine    *xlsx.Row
	nameLine        *xlsx.Row
	typeLine        *xlsx.Row
	exportParamLine *xlsx.Row
	data            map[any]any
	dataC           map[any]any
}

// GetDisplayName 获取显示名称
func (slf *Config) GetDisplayName() string {
	return slf.Name
}

// GetName 获取配置名称
func (slf *Config) GetName() string {
	return str.FirstUpper(slf.Sheet.Rows[0].Cells[1].String())
}

// GetIndexCount 获取索引数量
func (slf *Config) GetIndexCount() int {
	index, err := slf.Sheet.Rows[1].Cells[1].Int()
	if err != nil {
		panic(err)
	}
	return index
}

// GetData 获取数据
func (slf *Config) GetData() map[any]any {
	return slf.data
}

// GetJSON 获取JSON类型数据
func (slf *Config) GetJSON() []byte {
	bytes, err := jsonIter.MarshalIndent(slf.GetData(), "", "  ")
	if err != nil {
		panic(err)
	}
	return bytes
}

// GetJSONC 获取JSON类型数据
func (slf *Config) GetJSONC() []byte {
	bytes, err := jsonIter.MarshalIndent(slf.dataC, "", "  ")
	if err != nil {
		panic(err)
	}
	return bytes
}

func (slf *Config) GetVariable() string {
	var index = slf.GetIndexCount()
	if index <= 0 {
		return fmt.Sprintf("*_%s", slf.GetName())
	}
	var mapStr = "map[%s]%s"
	for i := 1; i < len(slf.typeLine.Cells); i++ {
		if slf.skipField[i] {
			continue
		}
		typ := slf.typeLine.Cells[i].String()
		if index > 0 {
			index--
			if index == 0 {
				mapStr = fmt.Sprintf(mapStr, typ, "%s")
			} else {
				mapStr = fmt.Sprintf(mapStr, typ, "map[%s]%s")
			}
		} else {
			mapStr = fmt.Sprintf(mapStr, "*_"+str.FirstUpper(slf.GetName()))
			break
		}
	}
	return fmt.Sprintf("%s", mapStr)
}

func (slf *Config) GetStruct() string {
	var result string
	if slf.GetIndexCount() <= 0 {
		for i := 4; i < len(slf.Rows); i++ {
			row := slf.Sheet.Rows[i]

			desc := row.Cells[0].String()
			if strings.HasPrefix(desc, "#") {
				continue
			}
			name := row.Cells[1].String()
			typ := row.Cells[2].String()
			exportParam := row.Cells[3].String()
			switch strings.ToLower(exportParam) {
			case "s", "sc", "cs":
				result += fmt.Sprintf("%s %s\n", str.FirstUpper(name), slf.GetType(typ))
			case "c":
				continue
			}
		}

		return fmt.Sprintf("type _%s struct{\n%s}", slf.GetName(), result)
	}
	for i := 1; i < len(slf.typeLine.Cells); i++ {
		if slf.skipField[i] {
			continue
		}
		typ := slf.typeLine.Cells[i].String()
		name := slf.nameLine.Cells[i].String()
		exportParam := slf.exportParamLine.Cells[i].String()
		switch strings.ToLower(exportParam) {
		case "s", "sc", "cs":
			name = str.FirstUpper(name)
			result += fmt.Sprintf("%s %s\n", name, slf.GetType(typ))
		case "c":
			continue
		}
	}
	return fmt.Sprintf("type _%s struct{\n%s}", str.FirstUpper(slf.GetName()), result)
}

func (slf *Config) GetType(fieldType string) string {
	if name, exist := basicTypeName[fieldType]; exist {
		return name
	} else if strings.HasPrefix(fieldType, "[]") {
		s := strings.TrimPrefix(fieldType, "[]")
		if name, exist := basicTypeName[s]; exist {
			return fmt.Sprintf("map[int]%s", name)
		} else {
			return slf.GetType(s)
		}
	}

	var s = strings.TrimSuffix(strings.TrimPrefix(fieldType, "{"), "}")
	var fields []string
	var field string
	var leftBrackets []int
	for i, c := range s {
		switch c {
		case ',':
			if len(leftBrackets) == 0 {
				fields = append(fields, field)
				field = ""
			} else {
				field += string(c)
			}
		case '{':
			leftBrackets = append(leftBrackets, i)
			field += string(c)
		case '}':
			leftBrackets = leftBrackets[:len(leftBrackets)-1]
			field += string(c)
			if len(leftBrackets) == 0 {
				fields = append(fields, field)
				field = ""
			}
		default:
			field += string(c)
		}
	}
	if len(field) > 0 {
		fields = append(fields, field)
	}

	var result = "*struct {\n%s}"
	var fieldStr string
	for _, fieldInfo := range fields {
		fieldName, fieldType := str.KV(strings.TrimSpace(fieldInfo), ":")
		n, exist := basicTypeName[fieldType]
		if exist {
			fieldStr += fmt.Sprintf("%s %s\n", str.FirstUpper(fieldName), n)
		} else {
			fieldStr += fmt.Sprintf("%s %s\n", str.FirstUpper(fieldName), slf.GetType(fieldType))
		}
	}

	return fmt.Sprintf(result, fieldStr)
}
