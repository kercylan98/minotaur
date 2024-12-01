package goluac

import (
	goluacv1 "github.com/kercylan98/minotaur/engine/vivid/goluac/v1"
	"github.com/kercylan98/minotaur/toolkit"
	lua "github.com/yuin/gopher-lua"
)

type LuaMessage = goluacv1.LuaMessage

// NewLuaMessage 创建一条可投递至 Goluac 的消息
func NewLuaMessage(name string, data any) *LuaMessage {
	return &LuaMessage{
		Name: name,
		Data: toolkit.MarshalJSON(data),
	}
}

func newFromLuaMessage(name string, data any) *LuaMessage {
	return &LuaMessage{
		Name:    name,
		Data:    toolkit.MarshalJSON(data),
		FromLua: true,
	}
}

func newFromLuaReplyMessage(name string, data []byte) *LuaMessage {
	return &LuaMessage{
		Name:         name,
		Data:         data,
		FromLua:      true,
		FromLuaReply: true,
	}
}

func toLuaMessage(state *lua.LState, m *LuaMessage) (lua.LValue, error) {
	value, err := decodeFromJson(state, m.Data)
	if err != nil {
		return nil, err
	}

	tbl := state.CreateTable(0, 2)
	tbl.RawSetH(lua.LString("name"), lua.LString(m.Name))
	tbl.RawSetH(lua.LString("data"), value)
	return tbl, nil
}
