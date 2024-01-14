package astgo_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/generator/astgo"
	"github.com/kercylan98/minotaur/utils/super"
	"testing"
)

func TestNewPackage(t *testing.T) {
	p, err := astgo.NewPackage(`/Users/kercylan/Coding.localized/Go/minotaur/server`)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(super.MarshalIndentJSON(p, "", "  ")))
}
