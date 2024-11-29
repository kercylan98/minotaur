package libs

import lua "github.com/yuin/gopher-lua"

func init() {

}

func pushError(state *lua.LState, err error) int {
	errTable := state.NewTable()
	errTable.RawSetString("type", lua.LString("error"))
	errTable.RawSetString("message", lua.LString(err.Error()))
	state.Push(errTable)
	return 1
}
