package super

import (
	jsonIter "github.com/json-iterator/go"
	"reflect"
)

var json = jsonIter.ConfigCompatibleWithStandardLibrary

// MarshalJSON 将对象转换为 json
//   - 当转换失败时，将返回 json 格式的空对象
func MarshalJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		return jsonZero(v)
	}
	return b
}

// MarshalJSONE 将对象转换为 json
//   - 当转换失败时，将返回错误信息
func MarshalJSONE(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// UnmarshalJSON 将 json 转换为对象
func UnmarshalJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// MarshalIndentJSON 将对象转换为 json
func MarshalIndentJSON(v interface{}, prefix, indent string) []byte {
	b, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return jsonZero(v)
	}
	return b
}

// MarshalToTargetWithJSON 将对象转换为目标对象
func MarshalToTargetWithJSON(src, dest interface{}) error {
	return json.Unmarshal(MarshalJSON(src), dest)
}

// 获取 json 格式的空对象
func jsonZero(v any) []byte {
	vof := reflect.Indirect(reflect.ValueOf(v))
	switch vof.Kind() {
	case reflect.Array, reflect.Slice:
		return StringToBytes("[]")
	case reflect.Bool:
		return StringToBytes("false") // false 在 JSON 中是不合法的，但是 Golang 中能够正确解析
	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return StringToBytes("0") // 0 在 JSON 中是不合法的，但是 Golang 中能够正确解析
	case reflect.String:
		return StringToBytes("\"\"")
	default:
		return StringToBytes("{}")
	}
}
