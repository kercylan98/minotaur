package xlsxsheet

import "github.com/tealeg/xlsx"

func newFieldDataScanner(sheet *xlsx.Sheet, index bool, fieldPos int) *fieldDataScanner {
	return &fieldDataScanner{
		sheet: sheet,
		pos:   fieldPos,
		index: index,
	}
}

type fieldDataScanner struct {
	sheet *xlsx.Sheet
	pos   int
	index bool
	i     int
}

func (s *fieldDataScanner) Next() string {
	if s.index {
		return s.onIndexNext()
	}
	return s.onNext()
}

func (s *fieldDataScanner) onNext() string {
	return s.sheet.Row(s.pos).Cells[4].String()
}

func (s *fieldDataScanner) onIndexNext() string {
	i := s.i
	s.i++
	row := i + 7
	return s.sheet.Row(row).Cells[s.pos].String()
}
