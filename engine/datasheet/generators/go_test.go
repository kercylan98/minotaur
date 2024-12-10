package generators_test

import (
	"github.com/kercylan98/minotaur/engine/datasheet"
	"github.com/kercylan98/minotaur/engine/datasheet/generators"
	"testing"
)

func TestGenerateGo(t *testing.T) {
	datasheets, err := datasheet.LoadDatasheets("../datasheet_template.xlsx")
	if err != nil {
		panic(err)
	}

	if err = generators.Go("datasheet", "./test/datasheet.go").Generate(datasheets); err != nil {
		panic(err)
	}
}
