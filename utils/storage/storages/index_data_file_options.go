package storages

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/storage"
)

// IndexDataFileStorageOption 索引数据文件存储器选项
type IndexDataFileStorageOption[I generic.Ordered, T storage.IndexDataItem[I]] func(storage *IndexDataFileStorage[I, T])

// WithIndexDataFileStorageEncoder 设置编码器
//   - 默认为 FileStorageJSONEncoder 编码器
func WithIndexDataFileStorageEncoder[I generic.Ordered, T storage.IndexDataItem[I]](encoder FileStorageEncoder[T]) IndexDataFileStorageOption[I, T] {
	return func(storage *IndexDataFileStorage[I, T]) {
		storage.encoder = encoder
	}
}

// WithIndexDataFileStorageDecoder 设置解码器
//   - 默认为 FileStorageJSONDecoder 解码器
func WithIndexDataFileStorageDecoder[I generic.Ordered, T storage.IndexDataItem[I]](decoder FileStorageDecoder[T]) IndexDataFileStorageOption[I, T] {
	return func(storage *IndexDataFileStorage[I, T]) {
		storage.decoder = decoder
	}
}

// WithIndexDataFileStorageSuffix 设置文件后缀
//   - 默认为 IndexDataFileDefaultSuffix
func WithIndexDataFileStorageSuffix[I generic.Ordered, T storage.IndexDataItem[I]](suffix string) IndexDataFileStorageOption[I, T] {
	return func(storage *IndexDataFileStorage[I, T]) {
		storage.suffix = suffix
	}
}
