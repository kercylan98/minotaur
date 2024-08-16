package xlsxsheet

import (
	"github.com/kercylan98/minotaur/ax/cmd/table"
	"github.com/tealeg/xlsx"
)

func newFieldScanner(table *Table, sheet *xlsx.Sheet) table.FieldScanner {
	return &fieldScanner{
		table: table,
		sheet: sheet,
	}
}

// fieldScanner 无索引字段扫描器
type fieldScanner struct {
	sheet *xlsx.Sheet
	i     int
	table *Table
}

func (s *fieldScanner) Next() table.Field {
	i := s.i + 4
	s.i++

	if i >= len(s.sheet.Rows) {
		return nil
	}

	f := &field{
		table:      s.table,
		sheet:      s.sheet,
		sheetIndex: i,
		desc:       s.sheet.Rows[i].Cells[0].String(),
		name:       s.sheet.Rows[i].Cells[1].String(),
		typ:        s.sheet.Rows[i].Cells[2].String(),
		param:      s.sheet.Rows[i].Cells[3].String(),
	}

	return f
}
