package datasheet_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/datasheet/v2"
	"testing"
)

func TestLoadDatasheets(t *testing.T) {
	datasheets, err := datasheet.LoadDatasheets("datasheet_template.xlsx")
	if err != nil {
		panic(err)
	}

	fmt.Println(datasheets)
}
