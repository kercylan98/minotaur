package xlsxsheet

import (
	"github.com/kercylan98/minotaur/ax/cmd/table"
	"github.com/tealeg/xlsx"
)

func newFieldScanner(sheet *xlsx.Sheet) table.FieldScanner {
	return &fieldScannerIndex{
		sheet: sheet,
	}
}

// fieldScanner 无索引字段扫描器
type fieldScanner struct {
	sheet *xlsx.Sheet
	i     int
}

func (s *fieldScanner) Next() table.Field {
	i := s.i + 4
	s.i++

	f := &field{
		desc:    s.sheet.Rows[i].Cells[0].String(),
		name:    s.sheet.Rows[i].Cells[1].String(),
		typ:     s.sheet.Rows[i].Cells[2].String(),
		param:   s.sheet.Rows[i].Cells[3].String(),
		scanner: newFieldDataScanner(s.sheet, false, i),
	}

	return f
}
