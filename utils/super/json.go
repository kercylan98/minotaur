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
		switch reflect.TypeOf(v).Kind() {
		case reflect.Array, reflect.Slice:
			return StringToBytes("[]")
		default:
			return StringToBytes("{}")
		}
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
		switch reflect.TypeOf(v).Kind() {
		case reflect.Array, reflect.Slice:
			return StringToBytes("[]")
		default:
			return StringToBytes("{}")
		}
	}
	return b
}
