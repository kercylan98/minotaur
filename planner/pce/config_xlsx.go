package pce

import (
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tealeg/xlsx"
	"strings"
)

func NewIndexXlsxConfig(sheet *xlsx.Sheet) *XlsxIndexConfig {
	config := &XlsxIndexConfig{
		sheet: sheet,
	}
	return config
}

// XlsxIndexConfig 内置的 Xlsx 配置
type XlsxIndexConfig struct {
	sheet *xlsx.Sheet
}

func (slf *XlsxIndexConfig) GetConfigName() string {
	return str.FirstUpper(strings.TrimSpace(slf.sheet.Rows[0].Cells[1].String()))
}

func (slf *XlsxIndexConfig) GetDisplayName() string {
	return slf.sheet.Name
}

func (slf *XlsxIndexConfig) GetDescription() string {
	return "暂无描述"
}

func (slf *XlsxIndexConfig) GetIndexCount() int {
	index, err := slf.sheet.Rows[1].Cells[1].Int()
	if err != nil {
		panic(err)
	}
	return index
}

func (slf *XlsxIndexConfig) GetFields() []dataField {
	var fields []dataField
	for x := 1; x < slf.getWidth(); x++ {
		var (
			desc       = slf.get(x, 3)
			name       = slf.get(x, 4)
			fieldType  = slf.get(x, 5)
			exportType = slf.get(x, 6)
		)
		if desc == nil || name == nil || fieldType == nil || exportType == nil {
			continue
		}
		if len(desc.String()) == 0 || len(name.String()) == 0 || len(fieldType.String()) == 0 || len(exportType.String()) == 0 {
			continue
		}
		fields = append(fields, dataField{
			Name:       name.String(),
			Type:       fieldType.String(),
			ExportType: exportType.String(),
			Desc:       desc.String(),
		})
	}
	return fields
}

func (slf *XlsxIndexConfig) GetData() [][]dataInfo {
	var data [][]dataInfo
	var fields = slf.GetFields()
	for y := 7; y < slf.getHeight(); y++ {
		var line []dataInfo
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
				line = append(line, dataInfo{
					dataField: fields[x-1],
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

	return data
}

// getWidth 获取宽度
func (slf *XlsxIndexConfig) getWidth() int {
	return slf.sheet.MaxCol
}

// getHeight 获取高度
func (slf *XlsxIndexConfig) getHeight() int {
	return slf.sheet.MaxRow
}

// get 获取单元格
func (slf *XlsxIndexConfig) get(x, y int) *xlsx.Cell {
	if x < 0 || y < 0 || y >= len(slf.sheet.Rows) {
		return nil
	}
	row := slf.sheet.Rows[y]
	if x >= len(row.Cells) {
		return nil
	}
	return row.Cells[x]
}
