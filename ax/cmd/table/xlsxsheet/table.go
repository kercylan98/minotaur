package xlsxsheet

import (
	"github.com/kercylan98/minotaur/ax/cmd/table"
	"github.com/tealeg/xlsx"
)

func NewTable(sheet *xlsx.Sheet) table.Table {
	return &Table{
		sheet: sheet,
	}
}

type Table struct {
	sheet        *xlsx.Sheet
	fieldScanIdx int
}

func (t *Table) GetIndex() int {
	count, err := t.sheet.Rows[1].Cells[1].Int()
	if err != nil {
		panic(err)
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
		return newFieldScannerIndex(t.sheet, count)
	} else {
		return newFieldScanner(t.sheet)
	}
}
