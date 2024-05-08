package scanner

import (
	"errors"
	"fmt"
	jsonIter "github.com/json-iterator/go"
	"github.com/kercylan98/minotaur/configuration/raw"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

var json = jsonIter.ConfigCompatibleWithStandardLibrary

// NewXlsxSheetScanner 创建一个Xlsx表格扫描器
func NewXlsxSheetScanner(sheet *xlsx.Sheet) *XlsxSheetScanner {
	return &XlsxSheetScanner{
		Sheet: sheet,
	}
}

// XlsxSheetScanner Xlsx表格扫描器
type XlsxSheetScanner struct {
	*xlsx.Sheet
}

func (x *XlsxSheetScanner) StructScan() (raw.Config, error) {
	config := raw.NewConfig(x.cell(0, 1), x.Name)
	configIndexNum := x.cell(1, 1)
	if configIndexNum == "" {
		return config, errors.New("config index number is empty")
	}
	indexNum, err := strconv.Atoi(strings.TrimSpace(configIndexNum))
	if err != nil {
		return config, fmt.Errorf("config index number is not a number: %s", configIndexNum)
	}

	var index int
	var parseField = func(desc, name, typ, param string, pos int) error {
		if typ == "" || param == "" || name == "" {
			return nil
		}

		// 导出参数限制
		switch strings.ToLower(param) {
		case "s", "c", "srv", "cli", "server", "client", "sc", "cs", "c/s", "s/c", "server/client", "client/server":
		default:
			return nil
		}

		isKey := index < indexNum
		if indexNum <= 0 {
			isKey = false
		}
		if err := raw.AddField(&config, name, desc, typ, param, index, isKey, pos); err != nil {
			return err
		}

		index++
		return nil
	}

	if indexNum > 0 {
		// 宽表
		for col := 1; col < x.MaxCol; col++ {
			desc := strings.TrimSpace(x.cell(3, col))
			name := strings.TrimSpace(x.cell(4, col))
			typ := strings.TrimSpace(x.cell(5, col))
			param := strings.TrimSpace(x.cell(6, col))

			if err = parseField(desc, name, typ, param, col); err != nil {
				return config, err
			}
		}
	} else {
		// 长表
		for row := 3; row < x.MaxRow; row++ {
			desc := strings.TrimSpace(x.cell(row, 0))
			name := strings.TrimSpace(x.cell(row, 1))
			typ := strings.TrimSpace(x.cell(row, 2))
			param := strings.TrimSpace(x.cell(row, 3))

			if err = parseField(desc, name, typ, param, row); err != nil {
				return config, err
			}
		}
	}

	return config, nil
}

func (x *XlsxSheetScanner) DataScan(fields []raw.Field) (any, error) {
	configIndexNum := x.cell(1, 1)
	if configIndexNum == "" {
		return nil, errors.New("config index number is empty")
	}
	indexNum, err := strconv.Atoi(strings.TrimSpace(configIndexNum))
	if err != nil {
		return nil, fmt.Errorf("config index number is not a number: %s", configIndexNum)
	}

	var data = make(map[any]any)
	var root = data

	// 应该根据字段来查找数据，如果数据存在索引，那么应该是多个 map[any]any 嵌套，最后一层 map 是包含所有字段的 map 具体数据
	// 例如：{"1": { "john”: { "id": 1, "name": "john", "age": 18 }}}
	// 如果没有索引，那么应该是一个 map[any]any，key 是字段名，value 是字段值

	if indexNum > 0 {
		// 宽表
		for row := 7; row < x.MaxRow; row++ {
			data = root
			// 数据检查
			end := false
			for _, field := range fields {
				col := field.GetPosition()
				if field.IsKey() {
					if x.cell(row, col) == "" {
						end = true
						break
					}
				}
			}
			if end {
				break
			}
			// 数据生成
			var keyValues = make(map[any]any)
			for _, field := range fields {
				col := field.GetPosition()
				valueStr := x.cell(row, col)
				value, err := x.formatValue(field, valueStr)
				if err != nil {
					return nil, err
				}

				if field.IsKey() {
					keyValues[field.GetName()] = value
					var next map[any]any
					if m, ok := data[value]; !ok {
						next = make(map[any]any)
					} else {
						next = m.(map[any]any)
					}
					data[value] = next
					data = next
				} else {
					data[field.GetName()] = value
				}
			}
			for k, v := range keyValues {
				data[k] = v
			}
		}
	} else {
		// 长表
		for _, field := range fields {
			row := field.GetPosition()
			valueStr := x.cell(row, 4)
			value, err := x.formatValue(field, valueStr)
			if err != nil {
				return nil, err
			}
			data[field.GetName()] = value
		}
	}

	return root, nil
}

func (x *XlsxSheetScanner) cell(row, col int) string {
	if cell := x.Cell(row, col); cell != nil {
		return cell.String()
	}
	return ""
}

func (x *XlsxSheetScanner) formatValue(field raw.Field, v string) (a any, err error) {
	typ := field.GetType()
	if !raw.IsBasicType(typ) {
		switch {
		case strings.HasPrefix(v, "["): // 数组
			var data []any
			if err := json.Unmarshal([]byte(v), &data); err != nil {
				return nil, err
			}
			return data, nil
		default:
			var data = make(map[string]any)
			if err := json.Unmarshal([]byte(v), &data); err != nil {
				return nil, err
			}
			return data, nil
		}
	}

	switch typ {
	case raw.FieldTypeInt:
		if a, err = strconv.Atoi(v); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeInt8:
		if a, err = strconv.ParseInt(v, 10, 8); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeInt16:
		if a, err = strconv.ParseInt(v, 10, 16); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeInt32:
		if a, err = strconv.ParseInt(v, 10, 32); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeInt64:
		if a, err = strconv.ParseInt(v, 10, 64); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeUint:
		if a, err = strconv.ParseUint(v, 10, 0); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeUint8:
		if a, err = strconv.ParseUint(v, 10, 8); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeUint16:
		if a, err = strconv.ParseUint(v, 10, 16); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeUint32:
		if a, err = strconv.ParseUint(v, 10, 32); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeUint64:
		if a, err = strconv.ParseUint(v, 10, 64); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeFloat32:
		if a, err = strconv.ParseFloat(v, 32); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeFloat64:
		if a, err = strconv.ParseFloat(v, 64); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeBool:
		if a, err = strconv.ParseBool(v); err != nil && v == "" {
			err = nil
		}
	case raw.FieldTypeString:
		return v, nil
	default:
		return v, fmt.Errorf("unsupported type: %s", field.GetType())
	}

	return a, err
}
