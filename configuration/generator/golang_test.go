package generator_test

import (
	"github.com/kercylan98/minotaur/configuration"
	"github.com/kercylan98/minotaur/configuration/generator"
	"github.com/kercylan98/minotaur/configuration/source"
	"github.com/tealeg/xlsx"
	"testing"
)

func TestGolang_Generate(t *testing.T) {
	xlsxFile, err := xlsx.OpenFile(`D:\sources\minotaur\configuration\xlsx_template.xlsx`)
	if err != nil {
		panic(err)
	}

	s := source.NewXlsxSheet(xlsxFile.Sheets[1])

	sourceFields := s.Fields()

	var fields = make([]*configuration.Field, 0)

	for _, sourceField := range sourceFields {
		field, err := configuration.ParseStructure(sourceField.Structure)
		if err != nil {
			panic(err)
		}
		fields = append(fields, &configuration.Field{
			Source:    sourceField,
			Structure: field,
		})
	}

	if err = new(generator.Golang).Generate(fields); err != nil {
		panic(err)
	}
}
