package datasheet_test

import (
	"fmt"
	datasheet "github.com/kercylan98/minotaur/engine/datasheet/generators/test"
	"github.com/kercylan98/minotaur/toolkit"
	"os"
	"testing"
)

func TestDatasheet(t *testing.T) {
	err := datasheet.Load(func(sign datasheet.DatasheetSign, data any) (loaded bool, err error) {
		f, err := os.ReadFile("./" + string(sign) + ".json")
		if err != nil {
			return false, err
		}
		return true, toolkit.UnmarshalJSONE(f, data)
	})
	if err != nil {
		panic(err)
	}

	datasheet.UpAll()
	activity := datasheet.GetActivity()
	global := datasheet.GetGlobal()
	_ = activity
	_ = global
	fmt.Println(global.InitCoins.Value().String())
}
