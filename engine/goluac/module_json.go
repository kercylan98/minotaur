package goluac

import (
	"encoding/json"
	"errors"
	"github.com/kercylan98/minotaur/toolkit"
	lua "github.com/yuin/gopher-lua"
)

var (
	errNested      = errors.New("cannot jsonEncode recursively nested tables to JSON")
	errSparseArray = errors.New("cannot jsonEncode sparse array")
	errInvalidKeys = errors.New("cannot jsonEncode mixed or invalid key types")
	errInvalidType = errors.New("invalid type")
)

func init() {
	moduleGoFuncInject("json", "encode", jsonEncode)
	moduleGoFuncInject("json", "decode", jsonDecode)
}

func jsonEncode(ctx *actorContext) lua.LGFunction {
	return func(state *lua.LState) int {
		data, err := readToJson(state, 1)
		if err != nil {
			return pushError(state, err)
		}
		state.Push(lua.LString(data))
		return 1
	}
}
func jsonDecode(ctx *actorContext) lua.LGFunction {
	return func(state *lua.LState) int {
		str := state.CheckString(1)

		decoded, err := decodeFromJson(state, []byte(str))
		if err != nil {
			return pushError(state, err)
		}
		state.Push(decoded)
		return 1
	}
}

func decodeFromJson(state *lua.LState, data []byte) (lua.LValue, error) {
	var value any
	err := toolkit.UnmarshalJSONE(data, &value)
	if err != nil {
		return lua.LNil, err
	}
	decoded := jsonDecodeValue(state, value)
	return decoded, nil
}

func jsonDecodeValue(state *lua.LState, value any) lua.LValue {
	switch converted := value.(type) {
	case bool:
		return lua.LBool(converted)
	case float64:
		return lua.LNumber(converted)
	case string:
		return lua.LString(converted)
	case json.Number:
		return lua.LString(converted)
	case []any:
		arr := state.CreateTable(len(converted), 0)
		for _, item := range converted {
			arr.Append(jsonDecodeValue(state, item))
		}
		return arr
	case map[string]any:
		tbl := state.CreateTable(0, len(converted))
		for key, item := range converted {
			tbl.RawSetH(lua.LString(key), jsonDecodeValue(state, item))
		}
		return tbl
	case nil:
		return lua.LNil
	}

	return lua.LNil
}

func readToJson(state *lua.LState, n int) ([]byte, error) {
	value := state.CheckAny(n)
	return json.Marshal(jsonValue{
		LValue:  value,
		visited: make(map[*lua.LTable]bool),
	})
}

type jsonValue struct {
	lua.LValue
	visited map[*lua.LTable]bool
}

func (j jsonValue) MarshalJSON() (data []byte, err error) {
	switch converted := j.LValue.(type) {
	case lua.LBool:
		data, err = json.Marshal(bool(converted))
	case lua.LNumber:
		data, err = json.Marshal(float64(converted))
	case *lua.LNilType:
		data = []byte(`null`)
	case lua.LString:
		data, err = json.Marshal(string(converted))
	case *lua.LTable:
		if j.visited[converted] {
			return nil, errNested
		}
		j.visited[converted] = true

		key, value := converted.Next(lua.LNil)

		switch key.Type() {
		case lua.LTNil: // empty table
			data = []byte(`[]`)
		case lua.LTNumber:
			arr := make([]jsonValue, 0, converted.Len())
			expectedKey := lua.LNumber(1)
			for key != lua.LNil {
				if key.Type() != lua.LTNumber {
					err = errInvalidKeys
					return
				}
				if expectedKey != key {
					err = errSparseArray
					return
				}
				arr = append(arr, jsonValue{value, j.visited})
				expectedKey++
				key, value = converted.Next(key)
			}
			data, err = json.Marshal(arr)
		case lua.LTString:
			obj := make(map[string]jsonValue)
			for key != lua.LNil {
				if key.Type() != lua.LTString {
					err = errInvalidKeys
					return
				}
				obj[key.String()] = jsonValue{value, j.visited}
				key, value = converted.Next(key)
			}
			data, err = json.Marshal(obj)
		default:
			err = errInvalidKeys
		}
	default:
		err = errInvalidType
	}
	return
}
