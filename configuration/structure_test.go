package configuration_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/configuration"
	"github.com/kercylan98/minotaur/toolkit"
	"testing"
)

func TestParseStructure(t *testing.T) {
	s, err := configuration.ParseStructure("{id:int,name:string,info:{lv:int,exp:{mux:int,count:int}}}")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(toolkit.MarshalIndentJSON(s, "", "  ")))
}
