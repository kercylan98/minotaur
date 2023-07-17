package cs

import (
	"github.com/kercylan98/minotaur/planner/pce"
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tealeg/xlsx"
	"regexp"
	"strings"
)

func NewXlsx(sheet *xlsx.Sheet) *Xlsx {
	config := &Xlsx{
		sheet: sheet,
	}
	return config
}

// Xlsx 内置的 Xlsx 配置
type Xlsx struct {
	sheet *xlsx.Sheet
}

func (slf *Xlsx) GetConfigName() string {
	return str.FirstUpper(strings.TrimSpace(slf.sheet.Rows[0].Cells[1].String()))
}

func (slf *Xlsx) GetDisplayName() string {
	return slf.sheet.Name
}

func (slf *Xlsx) GetDescription() string {
	return "暂无描述"
}

func (slf *Xlsx) GetIndexCount() int {
	index, err := slf.sheet.Rows[1].Cells[1].Int()
	if err != nil {
		panic(err)
	}
	return index
}

func (slf *Xlsx) GetFields() []pce.DataField {
	var handle = func(desc, name, fieldType, exportType *xlsx.Cell) (pce.DataField, bool) {
		var field pce.DataField
		if desc == nil || name == nil || fieldType == nil || exportType == nil {
			return field, false
		}
		field = pce.DataField{
			Name:       strings.ReplaceAll(strings.ReplaceAll(str.FirstUpper(name.String()), "\r", ""), "\n", ""),
			Type:       fieldType.String(),
			ExportType: exportType.String(),
			Desc:       strings.ReplaceAll(strings.ReplaceAll(desc.String(), "\r", ""), "\n", ""),
		}
		if len(field.Name) == 0 || len(field.Type) == 0 || len(field.ExportType) == 0 {
			return field, false
		}

		if slf.checkFieldInvalid(field) {
			return field, false
		}

		return field, true
	}
	var fields []pce.DataField
	if slf.GetIndexCount() > 0 {
		for x := 1; x < slf.getWidth(); x++ {
			if field, match := handle(slf.get(x, 3), slf.get(x, 4), slf.get(x, 5), slf.get(x, 6)); match {
				fields = append(fields, field)
			}
		}
	} else {
		for y := 4; y < slf.getHeight(); y++ {
			if field, match := handle(slf.get(0, y), slf.get(1, y), slf.get(2, y), slf.get(3, y)); match {
				fields = append(fields, field)
			}
		}
	}
	return fields
}

func (slf *Xlsx) GetData() [][]pce.DataInfo {
	var data [][]pce.DataInfo
	var fields = slf.GetFields()
	if slf.GetIndexCount() > 0 {
		for y := 7; y < slf.getHeight(); y++ {
			var line []pce.DataInfo
			var stop bool
			for x := 0; x < slf.getWidth(); x++ {
				if prefixCell := slf.get(x, y); prefixCell != nil {
					prefix := prefixCell.String()
					if strings.HasPrefix(prefix, "#") {
						break
					}
				}
				if x == 0 {
					continue
				}
				var isIndex = x-1 < slf.GetIndexCount()

				var value string
				if valueCell := slf.get(x, y); valueCell != nil {
					value = valueCell.String()
				} else if isIndex {
					stop = true
					break
				}
				valueCell := slf.get(x, y)
				if valueCell == nil {
					break
				}
				if len(fields) > x-1 {
					line = append(line, pce.DataInfo{
						DataField: fields[x-1],
						Value:     value,
					})
				}
			}
			if len(line) > 0 {
				data = append(data, line)
			}
			if stop {
				break
			}
		}
	} else {
		var line []pce.DataInfo
		for i, field := range slf.GetFields() {
			var value string
			if valueCell := slf.get(4, 4+i); valueCell != nil {
				value = valueCell.String()
			}
			line = append(line, pce.DataInfo{
				DataField: field,
				Value:     value,
			})
		}
		data = append(data, line)
	}
	return data
}

// getWidth 获取宽度
func (slf *Xlsx) getWidth() int {
	return slf.sheet.MaxCol
}

// getHeight 获取高度
func (slf *Xlsx) getHeight() int {
	return slf.sheet.MaxRow
}

// get 获取单元格
func (slf *Xlsx) get(x, y int) *xlsx.Cell {
	if x < 0 || y < 0 || y >= len(slf.sheet.Rows) {
		return nil
	}
	row := slf.sheet.Rows[y]
	if x >= len(row.Cells) {
		return nil
	}
	return row.Cells[x]
}

func (slf *Xlsx) checkFieldInvalid(field pce.DataField) bool {
	switch strings.ToLower(field.ExportType) {
	case "s", "c", "sc", "cs":
	default:
		return true
	}

	pattern := "^[a-zA-Z][a-zA-Z0-9]*$"
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(field.Name) {
		return true
	}

	if strings.HasPrefix(field.Name, "#") || strings.HasPrefix(field.Type, "#") {
		return true
	}

	return false
}
