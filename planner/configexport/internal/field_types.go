package internal

import (
	"github.com/kercylan98/minotaur/utils/str"
	"github.com/tidwall/gjson"
	"math"
	"strconv"
	"strings"
)

var basicTypeName = map[string]string{
	"string":  "string",
	"int":     "int",
	"int8":    "int8",
	"int16":   "int16",
	"int32":   "int32",
	"int64":   "int64",
	"uint":    "uint",
	"uint8":   "uint8",
	"uint16":  "uint16",
	"uint32":  "uint32",
	"uint64":  "uint64",
	"float32": "float32",
	"float64": "float64",
	"float":   "float64",
	"double":  "float64",
	"number":  "float64",
	"byte":    "byte",
	"rune":    "rune",
	"bool":    "bool",
	"boolean": "bool",
}

var basicType = map[string]func(fieldValue string) any{
	"string":  withStringType,
	"int":     withIntType,
	"int8":    withInt8Type,
	"int16":   withInt16Type,
	"int32":   withInt32Type,
	"int64":   withInt64Type,
	"uint":    withUintType,
	"uint8":   withUint8Type,
	"uint16":  withUint16Type,
	"uint32":  withUint32Type,
	"uint64":  withUint64Type,
	"float32": withFloat32Type,
	"float64": withFloat64Type,
	"float":   withFloat64Type,
	"double":  withFloat64Type,
	"number":  withFloat64Type,
	"byte":    withByteType,
	"rune":    withRuneType,
	"bool":    withBoolType,
	"boolean": withBoolType,
}

func getValueWithType(fieldType string, fieldValue string) any {
	fieldType = strings.ToLower(strings.TrimSpace(fieldType))
	handle, exist := basicType[fieldType]
	if exist {
		return handle(fieldValue)
	} else {
		return withStructType(fieldType, fieldValue)
	}
}

func withStructType(fieldType string, fieldValue string) any {
	// {id:int,name:string,info:{lv:int,exp:int}}
	if strings.HasPrefix(fieldType, "[]") {
		return withSliceType(fieldType, fieldValue)
	} else if !strings.HasPrefix(fieldType, "{") || !strings.HasSuffix(fieldType, "}") {
		return nil
	}
	var s = strings.TrimSuffix(strings.TrimPrefix(fieldType, "{"), "}")
	var data = map[any]any{}
	var fields []string
	var field string
	var leftBrackets []int
	for i, c := range s {
		switch c {
		case ',':
			if len(leftBrackets) == 0 {
				fields = append(fields, field)
				field = ""
			} else {
				field += string(c)
			}
		case '{':
			leftBrackets = append(leftBrackets, i)
			field += string(c)
		case '}':
			leftBrackets = leftBrackets[:len(leftBrackets)-1]
			field += string(c)
			if len(leftBrackets) == 0 {
				fields = append(fields, field)
				field = ""
			}
		default:
			field += string(c)
		}
	}
	if len(field) > 0 {
		fields = append(fields, field)
	}

	for _, fieldInfo := range fields {
		fieldName, fieldType := str.KV(strings.TrimSpace(fieldInfo), ":")
		handle, exist := basicType[fieldType]
		if exist {
			data[fieldName] = handle(gjson.Get(fieldValue, fieldName).String())
		} else {
			data[fieldName] = withStructType(fieldType, gjson.Get(fieldValue, fieldName).String())
		}
	}

	return data
}

func withSliceType(fieldType string, fieldValue string) any {
	if !strings.HasPrefix(fieldType, "[]") {
		return nil
	}
	t := strings.TrimPrefix(fieldType, "[]")
	var data = map[any]any{}
	gjson.ForEachLine(fieldValue, func(line gjson.Result) bool {
		line.ForEach(func(key, value gjson.Result) bool {
			handle, exist := basicType[t]
			if exist {
				data[len(data)] = handle(value.String())
			} else {
				data[len(data)] = withStructType(t, value.String())
			}
			return true
		})
		return true
	})
	return data
}

func withStringType(fieldValue string) any {
	return fieldValue
}

func withIntType(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	return value
}

func withInt8Type(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	if value < 0 {
		return int8(0)
	} else if value > math.MaxInt8 {
		return int8(math.MaxInt8)
	}
	return int8(value)
}

func withInt16Type(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	if value < 0 {
		return int16(0)
	} else if value > math.MaxInt16 {
		return int16(math.MaxInt16)
	}
	return int16(value)
}

func withInt32Type(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	if value < 0 {
		return int32(0)
	} else if value > math.MaxInt32 {
		return int32(math.MaxInt32)
	}
	return int32(value)
}

func withInt64Type(fieldValue string) any {
	value, _ := strconv.ParseInt(fieldValue, 10, 64)
	return value
}

func withUintType(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	if value < 0 {
		return uint(0)
	}
	return uint(value)
}

func withUint8Type(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	if value < 0 {
		return uint8(0)
	} else if value > math.MaxUint8 {
		return uint8(math.MaxUint8)
	}
	return uint8(value)
}

func withUint16Type(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	if value < 0 {
		return uint16(0)
	} else if value > math.MaxUint16 {
		return uint16(math.MaxUint16)
	}
	return uint16(value)
}

func withUint32Type(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	if value < 0 {
		return uint32(0)
	} else if value > math.MaxUint32 {
		return uint32(math.MaxUint32)
	}
	return uint32(value)
}

func withUint64Type(fieldValue string) any {
	value, _ := strconv.ParseInt(fieldValue, 10, 64)
	return uint64(value)
}

func withFloat32Type(fieldValue string) any {
	value, _ := strconv.ParseFloat(fieldValue, 32)
	return value
}

func withFloat64Type(fieldValue string) any {
	value, _ := strconv.ParseFloat(fieldValue, 64)
	return value
}

func withByteType(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	if value < 0 {
		return byte(0)
	} else if value > math.MaxUint8 {
		return byte(math.MaxInt8)
	}
	return byte(value)
}

func withRuneType(fieldValue string) any {
	value, _ := strconv.Atoi(fieldValue)
	if value < math.MinInt32 {
		return rune(0)
	} else if value > math.MaxInt32 {
		return rune(math.MaxInt32)
	}
	return rune(value)
}

func withBoolType(fieldValue string) any {
	switch fieldValue {
	case "0", "false", "!":
		return false
	case "1", "true", "&":
		return true
	default:
		return false
	}
}
