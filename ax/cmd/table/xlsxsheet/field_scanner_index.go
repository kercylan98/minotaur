package xlsxsheet

import (
	"github.com/kercylan98/minotaur/ax/cmd/table"
	"github.com/tealeg/xlsx"
)

func newFieldScannerIndex(table *Table, sheet *xlsx.Sheet, indexCount int) table.FieldScanner {
	return &fieldScannerIndex{
		table: table,
		sheet: sheet,
		count: indexCount,
	}
}

// fieldScannerIndex 有索引字段扫描器
type fieldScannerIndex struct {
	table *Table
	sheet *xlsx.Sheet
	i     int
	count int
}

func (s *fieldScannerIndex) Next() table.Field {
	i := s.i
	s.i++
	var index = i + 1
	if index > s.count {
		index = 0
	}
	i++
	if i >= s.sheet.MaxCol {
		return nil
	}
	f := &field{
		table:      s.table,
		sheet:      s.sheet,
		sheetIndex: i,
		index:      index,
		desc:       s.sheet.Rows[3].Cells[i].String(),
		name:       s.sheet.Rows[4].Cells[i].String(),
		typ:        s.sheet.Rows[5].Cells[i].String(),
		param:      s.sheet.Rows[6].Cells[i].String(),
		idxTable:   true,
	}

	return f
}
