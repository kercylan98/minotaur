package storages

import "encoding/json"

// FileStorageEncoder 全局数据文件存储编码器
type FileStorageEncoder[T any] func(data T) ([]byte, error)

// FileStorageDecoder 全局数据文件存储解码器
type FileStorageDecoder[T any] func(bytes []byte, data T) error

// FileStorageJSONEncoder JSON 编码器
func FileStorageJSONEncoder[T any]() FileStorageEncoder[T] {
	return func(data T) ([]byte, error) {
		return json.Marshal(data)
	}
}

// FileStorageJSONDecoder JSON 解码器
func FileStorageJSONDecoder[T any]() FileStorageDecoder[T] {
	return func(bytes []byte, data T) error {
		return json.Unmarshal(bytes, data)
	}
}
