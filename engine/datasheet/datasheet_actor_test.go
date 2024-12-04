package datasheet_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/datasheet"
	"github.com/kercylan98/minotaur/engine/datasheet/typemarshalers"
	"testing"
)

func TestDataSheetActor(t *testing.T) {
	datasheets, err := datasheet.ExportLoadDatasheets("datasheet_template.xlsx")
	if err != nil {
		panic(err)
	}

	for _, d := range datasheets {
		t := d.ToType()
		if t != nil {
			v, err := typemarshalers.NewGo(typemarshalers.FunctionalGoConfigurator(func(config *typemarshalers.GoConfiguration) {
				config.WithTypePrefix()
			})).Marshal(t)
			if err != nil {
				panic(err)
			}
			fmt.Println(v)
		}
	}
}
