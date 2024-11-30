package goluac

import (
	"github.com/kercylan98/minotaur/toolkit"
	lua "github.com/yuin/gopher-lua"
)

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

func (m *LuaMessage) toLuaMessage(state *lua.LState) (lua.LValue, error) {
	value, err := decodeFromJson(state, m.Data)
	if err != nil {
		return nil, err
	}

	tbl := state.CreateTable(0, 2)
	tbl.RawSetH(lua.LString("name"), lua.LString(m.Name))
	tbl.RawSetH(lua.LString("data"), value)
	return tbl, nil
}

// Unmarshal 将消息的数据解析到指定对象
func (m *LuaMessage) Unmarshal(dst any) {
	toolkit.UnmarshalJSON(m.Data, dst)
}

// UnmarshalE 将消息的数据解析到指定对象，并返回过程中发生的错误
func (m *LuaMessage) UnmarshalE(dst any) error {
	return toolkit.UnmarshalJSONE(m.Data, dst)
}

// UnmarshalP 将消息的数据解析到指定对象，当错误发生时，将会执行 panic
func (m *LuaMessage) UnmarshalP(dst any) {
	if err := toolkit.UnmarshalJSONE(m.Data, dst); err != nil {
		panic(err)
	}
}
