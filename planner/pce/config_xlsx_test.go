package pce_test

import (
	"github.com/kercylan98/minotaur/planner/pce"
	"github.com/tealeg/xlsx"
	"testing"
)

func TestNewXlsxIndexConfig(t *testing.T) {
	f, err := xlsx.OpenFile(`D:\sources\minotaur\planner\ce\template.xlsx`)
	if err != nil {
		panic(err)
	}
	xlsxConfig := pce.NewIndexXlsxConfig(f.Sheets[1])

	loader := pce.NewLoader()
	loader.BindField(
		new(pce.Int),
		new(pce.String),
	)

	loader.LoadStruct(xlsxConfig)
	loader.LoadData(xlsxConfig)

}
