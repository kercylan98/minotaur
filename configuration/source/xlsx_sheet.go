package source

import (
	"github.com/kercylan98/minotaur/configuration"
	"github.com/tealeg/xlsx"
	"strings"
)

// NewXlsxSheet 创建一个 XLSX Sheet 配置源
func NewXlsxSheet(sheet *xlsx.Sheet) *XLSXSheet {
	s := &XLSXSheet{
		sheet: sheet,
	}
	return s
}

// XLSXSheet XLSX 配置源
type XLSXSheet struct {
	sheet *xlsx.Sheet
}

func (x *XLSXSheet) DisplayName() string {
	return x.sheet.Name
}

func (x *XLSXSheet) Name() string {
	if row := x.sheet.Row(0); row != nil {
		if len(row.Cells) > 1 {
			return row.Cells[1].String()
		}
	}
	return ""
}

func (x *XLSXSheet) Description() string {
	return x.DisplayName()
}

func (x *XLSXSheet) Fields() []*configuration.SourceField {
	fields := make([]*configuration.SourceField, 0)

	cell := 1
	for cell <= x.sheet.MaxCol {
		var field = &configuration.SourceField{
			DisplayName: x.tryRead(3, cell),
			Name:        x.tryRead(4, cell),
			Structure:   x.tryRead(5, cell),
			Index:       cell,
		}
		switch strings.ToUpper(x.tryRead(6, cell)) {
		case "S", "SRV", "SERVER":
			field.Type = configuration.SourceFieldTypeServer
		case "C", "CLI", "CLIENT":
			field.Type = configuration.SourceFieldTypeClient
		case "SC", "CS", "COMMON", "SERVER_CLIENT", "CLIENT_SERVER":
			field.Type = configuration.SourceFieldTypeCommon
		default:
			field.Type = configuration.SourceFieldTypeInvalid
		}
		fields = append(fields, field)
		cell++
	}

	return fields
}

func (x *XLSXSheet) Rows(fields []*configuration.SourceField) []configuration.SourceRow {
	var rows = make([]configuration.SourceRow, 0)

	var row = 7
	for row <= x.sheet.MaxRow {
		var cells = make(configuration.SourceRow, 0)
		for _, field := range fields {
			cells = append(cells, &configuration.SourceCell{
				Field: field,
				Value: x.tryRead(row, field.Index),
			})
		}
		rows = append(rows, cells)
		row++
	}

	return rows
}

func (x *XLSXSheet) tryRead(row, cell int) string {
	if r := x.sheet.Row(row); r != nil {
		if len(r.Cells) > cell {
			return r.Cells[cell].String()
		}
	}
	return ""
}
