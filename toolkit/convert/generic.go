package convert

import (
	"github.com/kercylan98/minotaur/toolkit/constraints"
	"reflect"
)

func NumberToString[T constraints.Number](v T) string {
	vof := reflect.ValueOf(v).Interface()
	switch val := vof.(type) {
	case int:
		return IntToString(val)
	case int8:
		return Int8ToString(val)
	case int16:
		return Int16ToString(val)
	case int32:
		return Int32ToString(val)
	case int64:
		return Int64ToString(val)
	case uint:
		return UintToString(val)
	case uint8:
		return Uint8ToString(val)
	case uint16:
		return Uint16ToString(val)
	case uint32:
		return Uint32ToString(val)
	case uint64:
		return Uint64ToString(val)
	case float32:
		return Float32ToString(val)
	case float64:
		return Float64ToString(val)
	default:
		panic("unsupported type")
	}
}

func NumbersToStrings[T constraints.Number](v ...T) []string {
	var result []string
	for _, val := range v {
		result = append(result, NumberToString(val))
	}
	return result
}
