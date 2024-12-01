package goluacv1

import "github.com/kercylan98/minotaur/toolkit"

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
