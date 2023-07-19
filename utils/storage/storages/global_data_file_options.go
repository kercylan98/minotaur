package storages

// GlobalDataFileStorageOption 全局数据文件存储选项
type GlobalDataFileStorageOption[T any] func(storage *GlobalDataFileStorage[T])

// WithGlobalDataFileStorageEncoder 设置编码器
//   - 默认为 FileStorageJSONEncoder 编码器
func WithGlobalDataFileStorageEncoder[T any](encoder FileStorageEncoder[T]) GlobalDataFileStorageOption[T] {
	return func(storage *GlobalDataFileStorage[T]) {
		storage.encoder = encoder
	}
}

// WithGlobalDataFileStorageDecoder 设置解码器
//   - 默认为 FileStorageJSONDecoder 解码器
func WithGlobalDataFileStorageDecoder[T any](decoder FileStorageDecoder[T]) GlobalDataFileStorageOption[T] {
	return func(storage *GlobalDataFileStorage[T]) {
		storage.decoder = decoder
	}
}

// WithGlobalDataFileStorageSuffix 设置文件后缀
//   - 默认为 GlobalDataFileDefaultSuffix
func WithGlobalDataFileStorageSuffix[T any](suffix string) GlobalDataFileStorageOption[T] {
	return func(storage *GlobalDataFileStorage[T]) {
		storage.suffix = suffix
	}
}
