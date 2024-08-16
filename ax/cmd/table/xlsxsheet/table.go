package xlsxsheet

import (
	"fmt"
	"github.com/kercylan98/minotaur/ax/cmd/table"
	"github.com/tealeg/xlsx"
	"strings"
)

const (
	ExportModeCS ExportMode = iota
	ExportModeS
	ExportModeC
)

type ExportMode = int

func NewTable(sheet *xlsx.Sheet, mode ExportMode) table.Table {
	// 将 sheet 的列展开到最大
	for i := 0; i < len(sheet.Rows); i++ {
		row := sheet.Rows[i]
		if len(row.Cells) <= sheet.MaxCol {
			for j := len(row.Cells); j < sheet.MaxCol; j++ {
				row.Cells = append(row.Cells, xlsx.NewCell(row))
			}
		}
	}

	switch mode {
	case ExportModeCS, ExportModeS, ExportModeC:
	default:
		mode = ExportModeCS
	}

	return &Table{
		exportMode: mode,
		sheet:      sheet,
	}
}

type Table struct {
	sheet        *xlsx.Sheet
	fieldScanIdx int
	exportMode   ExportMode
}

func (t *Table) IsIgnore() bool {
	return strings.HasPrefix(strings.TrimSpace(t.sheet.Name), "#")
}

func (t *Table) GetIndex() int {
	count, err := t.sheet.Rows[1].Cells[1].Int()
	if err != nil {
		panic(fmt.Errorf("%s index count is not a number: %w", t.sheet.Name, err))
	}
	return count
}

func (t *Table) GetName() string {
	return t.sheet.Row(0).Cells[1].String()
}

func (t *Table) GetDescribe() string {
	return t.sheet.Name
}

func (t *Table) GetFields() table.FieldScanner {
	count := t.GetIndex()
	if count > 0 {
		return newFieldScannerIndex(t, t.sheet, count)
	} else {
		return newFieldScanner(t, t.sheet)
	}
}
