package lua2jsonparser

import (
	"github.com/kercylan98/minotaur/ax/cmd/table"
	lua "github.com/yuin/gopher-lua"
)

func New() table.DataParser {
	return &parser{}
}

type parser struct{}

func (*parser) Parse(script string) string {
	state := lua.NewState()
	defer state.Close()

	if err := state.DoString("target = " + script); err != nil {
		panic(err)
	}

	return encode(state)
}
