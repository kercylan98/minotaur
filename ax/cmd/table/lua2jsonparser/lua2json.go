package lua2jsonparser

import (
	"bytes"
	"encoding/json"
	"errors"
	lua "github.com/yuin/gopher-lua"
)

func encode(L *lua.LState) string {
	value := L.GetGlobal("target")

	data, err := json.Marshal(jsonConverter{
		LValue:  value,
		visited: make(map[*lua.LTable]bool),
	})
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	if err = json.Indent(&buffer, data, "", "  "); err != nil {
		panic(err)
	}

	return buffer.String()
}

type jsonConverter struct {
	lua.LValue
	visited map[*lua.LTable]bool
}

func (j jsonConverter) MarshalJSON() (data []byte, err error) {
	var result any
	switch luaValue := j.LValue.(type) {
	case lua.LBool:
		result = bool(luaValue)
	case lua.LNumber:
		result = float64(luaValue)
	case *lua.LNilType:
		result = nil
	case lua.LString:
		result = string(luaValue)
	case *lua.LTable:
		if j.visited[luaValue] {
			return nil, errors.New("cannot encode recursively nested tables to JSON")
		}
		j.visited[luaValue] = true

		key, value := luaValue.Next(lua.LNil)

		switch key.Type() {
		case lua.LTNil: // empty table
			result = []byte{}
		case lua.LTNumber:
			arr := make([]jsonConverter, 0, luaValue.Len())
			expectedKey := lua.LNumber(1)
			for key != lua.LNil {
				if key.Type() != lua.LTNumber {
					err = errors.New("cannot encode mixed or invalid key types")
					return
				}
				if expectedKey != key {
					err = errors.New("cannot encode sparse array")
					return
				}
				arr = append(arr, jsonConverter{value, j.visited})
				expectedKey++
				key, value = luaValue.Next(key)
			}
			result = arr
		case lua.LTString:
			obj := make(map[string]jsonConverter)
			for key != lua.LNil {
				if key.Type() != lua.LTString {
					err = errors.New("cannot encode mixed or invalid key types")
					return
				}
				obj[key.String()] = jsonConverter{value, j.visited}
				key, value = luaValue.Next(key)
			}
			result = obj
		default:
			err = errors.New("cannot encode mixed or invalid key types")
		}
	default:
		err = errors.New("cannot encode " + j.LValue.Type().String())
	}
	data, err = json.Marshal(result)
	return
}
