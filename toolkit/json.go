package toolkit

import jsonIter "github.com/json-iterator/go"

var json = jsonIter.ConfigCompatibleWithStandardLibrary

// MarshalJSON 将对象转换为 json
func MarshalJSON(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

// MarshalJSONE 将对象转换为 json
func MarshalJSONE(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// UnmarshalJSON 将 json 转换为对象
func UnmarshalJSON(data []byte, v interface{}) {
	_ = json.Unmarshal(data, v)
}

// UnmarshalJSONE 将 json 转换为对象
func UnmarshalJSONE(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// MarshalIndentJSON 将对象转换为 json
func MarshalIndentJSON(v interface{}, prefix, indent string) []byte {
	b, _ := json.MarshalIndent(v, prefix, indent)
	return b
}

// MarshalIndentJSONE 将对象转换为 json
func MarshalIndentJSONE(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// MarshalToTargetWithJSON 将对象转换为目标对象
func MarshalToTargetWithJSON(src, dest interface{}) {
	UnmarshalJSON(MarshalJSON(src), dest)
}
